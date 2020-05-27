package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/setup"

	"github.com/urfave/negroni"
)

func (a *App) setupHandlers() {
	s := setup.NewSetup(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/setup", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.Details)),
	))
	a.Router.Handle("/setup/details", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.Details)),
	))
	a.Router.Handle("/setup/invite", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.Invite)),
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
	a.Router.Handle("/setup/import_data", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.ImportData)),
	))
	a.Router.Handle("/setup/complete", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.Complete)),
	))
}
