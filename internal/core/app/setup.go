package app

import (
	"net/http"
	"strings"

	"gitlab.com/getlytica/lytica/internal/core/app/types"
	"gitlab.com/getlytica/lytica/internal/core/stripe"
	"gitlab.com/getlytica/lytica/internal/models"

	"github.com/gorilla/sessions"
)

func (a *App) details(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)

	t := []string{"partials/nav/_setup", "setup/details"}
	a.renderTemplate(w, t, session.Values)
}

func (a *App) invite(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)

	t := []string{"partials/nav/_setup", "setup/invite"}
	a.renderTemplate(w, t, session.Values)
}

func (a *App) subscribe(w http.ResponseWriter, r *http.Request) {
	if a.isSubscribed(w, r) {
		http.Redirect(w, r, "/setup/subscribe/success", http.StatusSeeOther)
	}

	session := a.getSession(w, r)
	a.stripeSessions(w, session)

	session.Values["showPlans"] = true

	t := []string{"partials/nav/_setup", "setup/subscribe"}
	a.renderTemplate(w, t, session.Values)
}

func (a *App) subscribeSuccess(w http.ResponseWriter, r *http.Request) {
	if !a.isSubscribed(w, r) {
		from := r.Header.Get("Referer")
		if !strings.Contains(from, "stripe.com") {
			http.Redirect(w, r, "/setup/subscribe/cancel", http.StatusSeeOther)
		}
	}

	session := a.getSession(w, r)
	session.Values["Flash"] = types.Flash{
		Success: types.FlashMessages["setup"]["subscribe"]["success"],
	}
	session.Values["allowNext"] = true
	session.Values["showPlans"] = false
	session.Values["isSubscribed"] = true

	err := session.Save(r, w)
	if err != nil {
		a.Logger.Error().Err(err).Msg("unable to save session")
	}

	t := []string{"partials/nav/_setup", "setup/subscribe", "partials/_flash"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) subscribeCancel(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	a.stripeSessions(w, session)

	session.Values["Flash"] = types.Flash{
		Error: types.FlashMessages["setup"]["subscribe"]["error"],
	}
	session.Values["showPlans"] = true

	t := []string{"partials/nav/_setup", "setup/subscribe", "partials/_flash"}
	a.renderTemplate(w, t, session.Values)
}

func (a *App) importData(w http.ResponseWriter, r *http.Request) {
	if !a.isSubscribed(w, r) {
		http.Redirect(w, r, "/setup/subscribe/cancel", http.StatusSeeOther)
	}

	session := a.getSession(w, r)

	t := []string{"partials/nav/_setup", "setup/import_data"}
	a.renderTemplate(w, t, session.Values)
}

func (a *App) complete(w http.ResponseWriter, r *http.Request) {
	if !a.isSubscribed(w, r) {
		http.Redirect(w, r, "/setup/subscribe/cancel", http.StatusSeeOther)
	}

	session := a.getSession(w, r)

	user := session.Values["User"].(models.User)
	user.SetupCompleted = true
	user.Save(a.Db)

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := []string{"partials/nav/_setup", "setup/complete"}
	a.renderTemplate(w, t, session.Values)
}

func (a *App) isSubscribed(w http.ResponseWriter, r *http.Request) bool {
	session := a.getSession(w, r)
	if session.Values["isSubscribed"] == nil {
		return false
	}

	subscribed := session.Values["isSubscribed"].(bool)
	if !subscribed {
		return false
	}

	return true
}

func (a *App) stripeSessions(w http.ResponseWriter, session *sessions.Session) {
	monthly, err := stripe.CheckoutSession("monthly")
	if err != nil {
		a.Logger.Error().Err(err).Msg("unable to generate a new stripe session")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	annual, err := stripe.CheckoutSession("annual")
	if err != nil {
		a.Logger.Error().Err(err).Msg("unable to generate a new stripe session")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	session.Values["stripeMonthlyId"] = monthly.ID
	session.Values["stripeAnnualId"] = annual.ID
}
