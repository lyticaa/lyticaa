package app

import "net/http"

func (a *App) sales(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "dashboard/sales", session.Values)
}

func (a *App) unitsSold(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "dashboard/units_sold", session.Values)
}

func (a *App) amazonCosts(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "dashboard/amazon/costs", session.Values)
}

func (a *App) advertising(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "dashboard/advertising", session.Values)
}

func (a *App) refunds(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "dashboard/refunds", session.Values)
}

func (a *App) shippingCredits(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "dashboard/shipping_credits", session.Values)
}

func (a *App) promotionalRebates(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "dashboard/promotional_rebates", session.Values)
}

func (a *App) totalCosts(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "dashboard/total_costs", session.Values)
}

func (a *App) netMargin(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "dashboard/net_margin", session.Values)
}
