package app

import (
	"fmt"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/accounts"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/sessions"
)

func (a *App) ForceSSL(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if helpers.Production() && r.Header.Get("x-forwarded-proto") != "https" {
			http.Redirect(w, r, fmt.Sprintf("https://%s%s", r.Host, r.RequestURI), http.StatusTemporaryRedirect)
			return
		}

		a.setConfig(w, r)
		next.ServeHTTP(w, r)
	})
}

func (a *App) Authenticated(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session := a.getSession(w, r)
	if _, ok := session.Values["profile"]; !ok {
		http.Redirect(w, r, helpers.RootRoute(), http.StatusSeeOther)
	} else {
		next(w, r)
	}
}

func (a *App) Admin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session := a.getSession(w, r)
	if session.Values["User"] == nil {
		http.Redirect(w, r, helpers.RootRoute(), http.StatusFound)
	}

	user := session.Values["User"].(models.UserModel)
	if !user.Admin {
		http.Redirect(w, r, helpers.DashboardRoute(), http.StatusFound)
	} else {
		next(w, r)
	}
}

func (a *App) SetupComplete(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session := a.getSession(w, r)

	user := helpers.GetSessionUser(session)
	accountPreferences := accounts.Preferences(r.Context(), user.ID, a.Data.Db)

	if !accountPreferences.SetupCompleted {
		http.Redirect(w, r, helpers.OnboardRoute(), http.StatusSeeOther)
		return
	} else {
		next(w, r)
	}
}

func (a *App) setConfig(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	session.Values["Config"] = types.NewConfig()
}

func (a *App) getSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := a.Data.SessionStore.Get(r, "auth-session")
	if err != nil {
		a.Monitoring.Logger.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	return session
}
