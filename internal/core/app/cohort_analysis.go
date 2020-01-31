package app

import "net/http"

func (a *App) cohortAnalysis(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "cohort_analysis", session.Values)
}
