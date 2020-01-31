package app

import "net/http"

func (a *App) profitLoss(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "profit_loss", session.Values)
}
