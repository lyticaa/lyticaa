package app

import "net/http"

func (a *App) details(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"onboard/details"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) team(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"onboard/team"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) importData(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"onboard/import_data"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) subscribe(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"onboard/subscribe"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) complete(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"onboard/complete"}

	a.renderTemplate(w, t, session.Values)
}
