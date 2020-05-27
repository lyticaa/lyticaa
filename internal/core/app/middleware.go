package app

import (
	"net/http"
	"os"

	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (a *App) forceSsl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("ENV") != "development" {
			if r.Header.Get("x-forwarded-proto") != "https" {
				sslUrl := "https://" + r.Host + r.RequestURI
				http.Redirect(w, r, sslUrl, http.StatusTemporaryRedirect)
				return
			}
		}

		a.setConfig(w, r)
		next.ServeHTTP(w, r)
	})
}

func (a *App) isAuthenticated(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session := a.getSession(w, r)
	if _, ok := session.Values["profile"]; !ok {
		http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
	} else {
		next(w, r)
	}
}

func (a *App) setupComplete(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session := a.getSession(w, r)
	if session.Values["User"] == nil {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}

	user := session.Values["User"].(models.User)
	if !user.SetupCompleted {
		a.Logger.Info().Msgf("%v", user)
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	} else {
		next(w, r)
	}
}

func (a *App) setConfig(w http.ResponseWriter, r *http.Request) {
	session := a.getSession(w, r)
	session.Values["Config"] = types.NewConfig()
}
