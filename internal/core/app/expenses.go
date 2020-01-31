package app

import "net/http"

func (a *App) expenses(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "expenses", session.Values)
}
