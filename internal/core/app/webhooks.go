package app

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/payments"
	"gitlab.com/getlytica/lytica-app/internal/models"

	"github.com/stripe/stripe-go"
)

const (
	checkoutCompletedEvent = "checkout.session.completed"
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
	case checkoutCompletedEvent:
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			a.Logger.Error().Err(err).Msg("unable to unmarshal stripe webhook JSON")
			w.WriteHeader(http.StatusBadRequest)
			return err
		}

		var customer sql.NullString
		if err := customer.Scan(payments.CustomerId(&session)); err != nil {
			a.Logger.Error().Err(err).Msg("unable to assign customer reference")
		}

		customerRefId := payments.CustomerRefId(&session)
		user := models.FindUser(customerRefId, a.Db)
		user.StripeUserId = customer

		if err := user.Save(a.Db); err != nil {
			a.Logger.Error().Err(err).Msg("unable to save user")
		}
	}

	return nil
}
