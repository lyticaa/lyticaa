package app

import "net/http"

func (a *App) expenses(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"expenses", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}
