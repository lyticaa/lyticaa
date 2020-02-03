package app

import "net/http"

func (a *App) cohortAnalysis(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"cohort_analysis", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}
