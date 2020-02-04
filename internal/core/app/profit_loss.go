package app

import "net/http"

func (a *App) profitLoss(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"partials/nav/_main", "profit_loss", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}
