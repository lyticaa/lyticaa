package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/cohort_analysis"

	"github.com/urfave/negroni"
)

func (a *App) cohortAnalysisHandlers() {
	c := cohort_analysis.NewCohortAnalysis(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/cohort_analysis", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.Overview)),
	))
}
