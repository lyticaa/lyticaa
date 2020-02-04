package app

import "net/http"

func (a *App) cohortAnalysis(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)

	t := []string{"partials/nav/_main", "cohort_analysis", "partials/_filters"}
	a.renderTemplate(w, t, session.Values)
}
