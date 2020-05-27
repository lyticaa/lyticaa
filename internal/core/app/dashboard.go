package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/dashboard"

	"github.com/urfave/negroni"
)

func (a *App) dashboardHandlers() {
	dashboard := dashboard.NewDashboard(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(dashboard.Overview)),
	))
	a.Router.Handle("/dashboard/metrics/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(dashboard.Metrics)),
	))
}
