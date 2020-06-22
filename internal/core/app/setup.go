package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/setup"
	"gitlab.com/getlytica/lytica-app/internal/core/payments"

	"github.com/urfave/negroni"
)

func (a *App) setupHandlers() {
	s := setup.NewSetup(a.Db, a.SessionStore, a.Logger, payments.NewStripePayments())

	a.Router.Handle("/setup", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.Subscribe)),
	))
	a.Router.Handle("/setup/subscribe", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.Subscribe)),
	))
	a.Router.Handle("/setup/subscribe/success", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.SubscribeSuccess)),
	))
	a.Router.Handle("/setup/subscribe/cancel", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.SubscribeCancel)),
	))
	a.Router.Handle("/setup/complete", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.Complete)),
	))
}
