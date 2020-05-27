package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/account"

	"github.com/urfave/negroni"
)

func (a *App) accountHandlers() {
	acct := account.NewAccount(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/account/notifications", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(acct.Notifications)),
	))
	a.Router.Handle("/account/invitations", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(acct.Invitations)),
	))
	a.Router.Handle("/account/subscription", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(acct.Subscription)),
	))
	a.Router.Handle("/account/change_password", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(acct.ChangePassword)),
	))
}
