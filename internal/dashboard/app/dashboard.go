package app

import "net/http"

func (a *App) sales(w http.ResponseWriter, r *http.Request) {
	a.renderTemplate(w, "dashboard/sales", nil)
}

func (a *App) unitsSold(w http.ResponseWriter, r *http.Request) {
	a.renderTemplate(w, "dashboard/units_sold", nil)
}

func (a *App) amazonCosts(w http.ResponseWriter, r *http.Request) {
	a.renderTemplate(w, "dashboard/amazon/costs", nil)
}

func (a *App) advertising(w http.ResponseWriter, r *http.Request) {
	a.renderTemplate(w, "dashboard/advertising", nil)
}

func (a *App) refunds(w http.ResponseWriter, r *http.Request) {
	a.renderTemplate(w, "dashboard/refunds", nil)
}

func (a *App) shippingCredits(w http.ResponseWriter, r *http.Request) {
	a.renderTemplate(w, "dashboard/shipping_credits", nil)
}

func (a *App) promotionalRebates(w http.ResponseWriter, r *http.Request) {
	a.renderTemplate(w, "dashboard/promotional_rebates", nil)
}

func (a *App) totalCosts(w http.ResponseWriter, r *http.Request) {
	a.renderTemplate(w, "dashboard/total_costs", nil)
}

func (a *App) netMargin(w http.ResponseWriter, r *http.Request) {
	a.renderTemplate(w, "dashboard/net_margin", nil)
}
