package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/account"
	"gitlab.com/getlytica/lytica-app/internal/core/payments"

	"github.com/urfave/negroni"
)

func (a *App) accountHandlers() {
	acct := account.NewAccount(a.Db, a.SessionStore, a.Logger, payments.NewStripePayments())

	a.Router.Handle("/account/notifications", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.Notifications)),
	))
	a.Router.Handle("/account/notifications/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.NotificationsByDate)),
	))
	a.Router.Handle("/account/subscription", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.Subscription)),
	))
	a.Router.Handle("/account/subscription/subscribe/{planId}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.Subscribe)),
	))
	a.Router.Handle("/account/subscription/cancel", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.CancelSubscription)),
	))
	a.Router.Handle("/account/subscription/change/{planId}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.ChangePlan)),
	))
	a.Router.Handle("/account/subscription/invoices", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.InvoicesByUser)),
	))
	a.Router.Handle("/account/change_password", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.ChangePassword)),
	))
}
