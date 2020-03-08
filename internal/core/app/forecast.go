package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/forecast"

	"github.com/urfave/negroni"
)

func (a *App) forecastHandlers() {
	f := forecast.NewForecast(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/forecast", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(f.Overview)),
	))
}
