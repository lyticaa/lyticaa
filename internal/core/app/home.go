package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/helpers"
)

func (a *App) home(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.SessionStore, a.Logger, w, r)

	t := []string{"partials/nav/_main", "home", "partials/_filters"}
	helpers.RenderTemplate(w, t, session.Values)
}
