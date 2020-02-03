package app

import "net/http"

func (a *App) totalSales(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"metrics/total_sales", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) unitsSold(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"metrics/units_sold", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) amazonCosts(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"metrics/amazon_costs", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) advertising(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"metrics/advertising", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) refunds(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"metrics/refunds", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) shippingCredits(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"metrics/shipping_credits", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) promotionalRebates(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"metrics/promotional_rebates", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) totalCosts(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"metrics/total_costs", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) netMargin(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"metrics/net_margin", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}
