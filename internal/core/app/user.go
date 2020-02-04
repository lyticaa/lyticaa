package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica/internal/core/user"
)

func (a *App) notifications(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{helpers.NavForSession(a.isSubscribed(w, r)), "user/notifications", "partials/_filters"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) invitations(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{helpers.NavForSession(a.isSubscribed(w, r)), "user/invitations"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) subscription(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	t := []string{helpers.NavForSession(a.isSubscribed(w, r)), "user/subscription"}

	a.renderTemplate(w, t, session.Values)
}

func (a *App) changePassword(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	u := user.NewUser(session.Values["userId"].(string), session.Values["email"].(string), a.Logger)
	_ = u.ResetPassword()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
