package app

import (
	"gitlab.com/getlytica/lytica-app/internal/core/app/auth"
)

func (a *App) authHandlers() {
	au := auth.NewAuth(a.Db, a.SessionStore, a.Logger)

	a.Router.HandleFunc("/auth/login", au.Login)
	a.Router.HandleFunc("/auth/logout", au.Logout)
	a.Router.HandleFunc("/auth/callback", au.Callback)
}
