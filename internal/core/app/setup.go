package app

import "net/http"

func (a *App) details(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"setup/details"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) invite(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"setup/invite"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) importData(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"setup/import_data"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) subscribe(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"setup/subscribe"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) complete(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{"setup/complete"}

	a.renderTemplate(w, t, session.Values)
}
