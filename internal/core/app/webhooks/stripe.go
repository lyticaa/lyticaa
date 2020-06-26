package webhooks

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"

	"github.com/stripe/stripe-go/v71"
)

const (
	checkoutCompletedEvent = "checkout.session.completed"
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
	case checkoutCompletedEvent:
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			wh.logger.Error().Err(err).Msg("failed to unmarshal stripe data")
			w.WriteHeader(http.StatusBadRequest)
			return err
		}

		var customer sql.NullString
		if err := customer.Scan(wh.stripe.CustomerId(&session)); err != nil {
			wh.logger.Error().Err(err).Msg("failed to assign customer reference")
		}

		user := models.User{UserId: wh.stripe.CustomerRefId(&session)}
		user.Load(wh.db)

		user.StripeUserId = customer
		if err := user.Save(wh.db); err != nil {
			wh.logger.Error().Err(err).Msg("failed to save user")
		}
	}

	return nil
}
