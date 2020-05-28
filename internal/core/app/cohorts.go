package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/cohorts"

	"github.com/urfave/negroni"
)

func (a *App) cohortsHandlers() {
	c := cohorts.NewCohorts(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/cohorts/high_margin", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.HighMargin)),
	))
	a.Router.Handle("/cohorts/high_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.HighMarginByDate)),
	))
	a.Router.Handle("/cohorts/low_margin", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.LowMargin)),
	))
	a.Router.Handle("/cohorts/low_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.LowMarginByDate)),
	))
	a.Router.Handle("/cohorts/negative_margin", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.NegativeMargin)),
	))
	a.Router.Handle("/cohorts/negative_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.NegativeMarginByDate)),
	))
}
