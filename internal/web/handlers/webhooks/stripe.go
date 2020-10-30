package webhooks

import (
	"database/sql"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/payments"

	"github.com/stripe/stripe-go/v72"
)

func (wh *Webhooks) Stripe(w http.ResponseWriter, r *http.Request) {
	body, err := wh.readBody(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e, err := wh.stripeEvent(body, w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = wh.parseStripeEvent(e, w); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (wh *Webhooks) stripeEvent(body []byte, w http.ResponseWriter, r *http.Request) (stripe.Event, error) {
	e, err := wh.stripe.ConstructEvent(body, r.Header.Get("Stripe-Signature"))
	if err != nil {
		wh.logger.Error().Err(err).Msg("bad signature")
		w.WriteHeader(http.StatusBadRequest)
		return stripe.Event{}, err
	}

	return e, nil
}

func (wh *Webhooks) parseStripeEvent(event stripe.Event, w http.ResponseWriter) error {
	switch event.Type {
	case payments.CheckoutSessionCompleted:
		if err := wh.checkoutSessionCompleted(event, w); err != nil {
			return err
		}
	}

	return nil
}

func (wh *Webhooks) checkoutSessionCompleted(event stripe.Event, w http.ResponseWriter) error {
	session, err := wh.stripe.EventSession(event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	user := models.LoadUser(wh.stripe.CustomerRefID(&session), wh.db)

	var customer sql.NullString
	if err := customer.Scan(*wh.stripe.CustomerID(&session)); err != nil {
		return err
	}
	user.StripeUserID = customer

	var subscription sql.NullString
	if err := subscription.Scan(*wh.stripe.SubscriptionID(&session)); err != nil {
		return err
	}
	user.StripeSubscriptionID = subscription

	var plan sql.NullString
	if err := plan.Scan(*wh.stripe.PlanID(&session)); err != nil {
		return err
	}
	user.StripePlanID = plan

	if err := user.Save(wh.db); err != nil {
		return err
	}

	return nil
}
