package app

import "net/http"

func (a *App) totalSales(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "metrics/total_sales", session.Values)
}

func (a *App) unitsSold(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "metrics/units_sold", session.Values)
}

func (a *App) amazonCosts(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "metrics/amazon_costs", session.Values)
}

func (a *App) advertising(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "metrics/advertising", session.Values)
}

func (a *App) refunds(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "metrics/refunds", session.Values)
}

func (a *App) shippingCredits(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "metrics/shipping_credits", session.Values)
}

func (a *App) promotionalRebates(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "metrics/promotional_rebates", session.Values)
}

func (a *App) totalCosts(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "metrics/total_costs", session.Values)
}

func (a *App) netMargin(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.renderTemplate(w, "metrics/net_margin", session.Values)
}
