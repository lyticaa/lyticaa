package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/expenses"

	"github.com/urfave/negroni"
)

func (a *App) expensesHandlers() {
	e := expenses.NewExpenses(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/expenses", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(e.Overview)),
	))
}
