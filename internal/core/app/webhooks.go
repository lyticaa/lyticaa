package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/webhooks"
	"gitlab.com/getlytica/lytica-app/internal/core/payments"

	"github.com/urfave/negroni"
)

func (a *App) webhookHandlers() {
	wh := webhooks.NewWebhooks(a.Db, a.Logger, payments.NewStripePayments())

	a.Router.Handle("/webhooks/stripe", negroni.New(
		negroni.Wrap(http.HandlerFunc(wh.Stripe)),
	))
}
