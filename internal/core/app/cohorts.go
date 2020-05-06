package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/cohorts"

	"github.com/urfave/negroni"
)

func (a *App) cohortsHandlers() {
	c := cohorts.NewCohorts(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/cohorts/high_margin", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.HighMargin)),
	))
	a.Router.Handle("/cohorts/low_margin", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.LowMargin)),
	))
	a.Router.Handle("/cohorts/negative_margin", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.NegativeMargin)),
	))
}
