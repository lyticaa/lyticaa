package app

import "net/http"

func (a *App) forecast(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"forecast", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}
