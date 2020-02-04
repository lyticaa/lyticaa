package app

import "net/http"

func (a *App) forecast(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)

	t := []string{"partials/nav/_main", "forecast", "partials/_filters"}
	a.renderTemplate(w, t, session.Values)
}
