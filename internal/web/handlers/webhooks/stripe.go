package webhooks

import (
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/web/pkg/accounts"
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/accounts/payments"

	"github.com/stripe/stripe-go/v72"
)

func (wh *Webhooks) Stripe(w http.ResponseWriter, r *http.Request) {
	body, err := wh.readBody(w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event, err := wh.stripeEvent(body, w, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = wh.parseStripeEvent(event, w, r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (wh *Webhooks) stripeEvent(body []byte, w http.ResponseWriter, r *http.Request) (stripe.Event, error) {
	e, err := wh.stripe.ConstructEvent(body, r.Header.Get("Stripe-Signature"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return stripe.Event{}, err
	}

	return e, nil
}

func (wh *Webhooks) parseStripeEvent(event stripe.Event, w http.ResponseWriter, r *http.Request) error {
	switch event.Type {
	case payments.CheckoutSessionCompleted:
		if err := accounts.Checkout(r.Context(), event, wh.db); err != nil {
			return err
		}
	}

	return nil
}
