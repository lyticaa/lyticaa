package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/dashboard"

	"github.com/urfave/negroni"
)

func (a *App) dashboardHandlers() {
	dashboard := dashboard.NewDashboard(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(dashboard.Overview)),
	))
}
