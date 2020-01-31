package app

import "net/http"

func (a *App) forecast(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "forecast", session.Values)
}
