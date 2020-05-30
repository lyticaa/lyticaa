package app

import (
	"database/sql"
	"fmt"
	"golang.org/x/text/currency"
	"io/ioutil"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/core/payments"
	"gitlab.com/getlytica/lytica-app/internal/models"

	"github.com/stripe/stripe-go/v71"
)

const (
	checkoutSessionCompleted    = "checkout.session.completed"
	customerSubscriptionCreated = "customer.subscription.created"
	customerSubscriptionDeleted = "customer.subscription.deleted"
	invoiceCreated              = "invoice.created"
	invoicePaymentFailed        = "invoice.payment_failed"
	invoicePaymentSucceeded     = "invoice.payment_succeeded"
)

func (a *App) stripeWebhooks(w http.ResponseWriter, r *http.Request) {
	body, err := a.readWebhookBody(w, r)
	if err != nil {
		return
	}

	e, err := a.getStripeWebhookEvent(body, w, r)
	if err != nil {
		return
	}

	if err = a.parseStripeWebhookEvent(e, w); err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *App) readWebhookBody(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.Logger.Error().Err(err).Msg("unable to read the body of the stripe webhook call")
		w.WriteHeader(http.StatusServiceUnavailable)
		return nil, err
	}

	return body, nil
}

func (a *App) getStripeWebhookEvent(body []byte, w http.ResponseWriter, r *http.Request) (stripe.Event, error) {
	e, err := payments.ConstructEvent(body, r.Header.Get("Stripe-Signature"))
	if err != nil {
		a.Logger.Error().Err(err).Msg("bad signature")
		w.WriteHeader(http.StatusBadRequest)
		return stripe.Event{}, err
	}

	return e, nil
}

func (a *App) parseStripeWebhookEvent(event stripe.Event, w http.ResponseWriter) error {
	switch event.Type {
	case checkoutSessionCompleted:
		session, err := payments.EventSession(event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return err
		}

		var customer sql.NullString
		if err := customer.Scan(payments.CustomerId(&session)); err != nil {
			a.Logger.Error().Err(err).Msg("unable to assign stripe customer id")
		}

		customerRefId := payments.CustomerRefId(&session)
		user := models.FindUser(customerRefId, a.Db)
		user.StripeUserId = customer

		var subscription sql.NullString
		if err := subscription.Scan(payments.SubscriptionId(&session)); err != nil {
			a.Logger.Error().Err(err).Msg("unable to assign stripe subscription id")
		}
		user.StripeSubscriptionId = subscription

		var plan sql.NullString
		if err := plan.Scan(payments.PlanId(&session)); err != nil {
			a.Logger.Error().Err(err).Msg("unable to assign stripe plan id")
		}
		user.StripePlanId = plan

		if err := user.Save(a.Db); err != nil {
			a.Logger.Error().Err(err).Msg("unable to save user")
		}
	case customerSubscriptionCreated:
		subscription, err := payments.EventSubscription(event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return err
		}

		user := models.FindUserByEmail(subscription.Customer.Email, a.Db)
		if user.Id == 0 {
			a.Logger.Error().Msg("unable to find user")
			w.WriteHeader(http.StatusBadRequest)
			return nil
		}

		message := fmt.Sprintf(types.NotificationMessages[customerSubscriptionCreated], subscription.Plan.Nickname)
		_ = models.CreateNotification(user.Id, message, a.Db)
	case customerSubscriptionDeleted:
		subscription, err := payments.EventSubscription(event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return err
		}

		user := models.FindUserByEmail(subscription.Customer.Email, a.Db)
		if user.Id == 0 {
			a.Logger.Error().Msg("unable to find user")
			w.WriteHeader(http.StatusBadRequest)
			return nil
		}

		message := fmt.Sprintf(types.NotificationMessages[customerSubscriptionDeleted], subscription.Plan.Nickname)
		_ = models.CreateNotification(user.Id, message, a.Db)
	case invoiceCreated:
		inv, err := payments.EventInvoice(event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return err
		}

		user := models.FindUserByEmail(inv.Customer.Email, a.Db)
		if user.Id == 0 {
			a.Logger.Error().Msg("unable to find user")
			w.WriteHeader(http.StatusBadRequest)
			return nil
		}

		unit, _ := currency.ParseISO(string(inv.Lines.Data[0].Currency))
		message := fmt.Sprintf(
			types.NotificationMessages[invoiceCreated],
			inv.Lines.Data[0].Description,
			fmt.Sprintf("%v %v", currency.Symbol(unit), payments.FormatAmount(inv.Lines.Data[0].Amount)),
		)
		_ = models.CreateNotification(user.Id, message, a.Db)
	case invoicePaymentFailed:
		inv, err := payments.EventInvoice(event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return err
		}

		user := models.FindUserByEmail(inv.Customer.Email, a.Db)
		if user.Id == 0 {
			a.Logger.Error().Msg("unable to find user")
			w.WriteHeader(http.StatusBadRequest)
			return nil
		}

		unit, _ := currency.ParseISO(string(inv.Lines.Data[0].Currency))
		message := fmt.Sprintf(
			types.NotificationMessages[invoicePaymentFailed],
			inv.Lines.Data[0].Description,
			fmt.Sprintf("%v %v", currency.Symbol(unit), payments.FormatAmount(inv.Lines.Data[0].Amount)),
		)
		_ = models.CreateNotification(user.Id, message, a.Db)
	case invoicePaymentSucceeded:
		inv, err := payments.EventInvoice(event)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return err
		}

		user := models.FindUserByEmail(inv.Customer.Email, a.Db)
		if user.Id == 0 {
			a.Logger.Error().Msg("unable to find user")
			w.WriteHeader(http.StatusBadRequest)
			return nil
		}

		unit, _ := currency.ParseISO(string(inv.Lines.Data[0].Currency))
		message := fmt.Sprintf(
			types.NotificationMessages[invoicePaymentSucceeded],
			inv.Lines.Data[0].Description,
			fmt.Sprintf("%v %v", currency.Symbol(unit), payments.FormatAmount(inv.Lines.Data[0].Amount)),
		)
		_ = models.CreateNotification(user.Id, message, a.Db)
	}

	return nil
}
