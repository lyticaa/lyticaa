package webhooks

import (
	"fmt"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/app/pkg/accounts"
	"github.com/lyticaa/lyticaa-app/internal/app/pkg/accounts/payments"

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
	event, err := wh.stripe.ConstructEvent(body, r.Header.Get("Stripe-Signature"))
	if err != nil {
		return stripe.Event{}, err
	}

	return event, nil
}

func (wh *Webhooks) parseStripeEvent(event stripe.Event, w http.ResponseWriter, r *http.Request) error {
	switch event.Type {
	case payments.CheckoutSessionCompleted:
		if err := accounts.Checkout(r.Context(), event, wh.db); err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}
