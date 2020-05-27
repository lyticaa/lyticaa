package dashboard

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (d *Dashboard) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(d.sessionStore, d.logger, w, r)

	t := []string{
		"partials/nav/_main",
		"dashboard/overview",
		"partials/_filters",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (d *Dashboard) Metrics(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(d.sessionStore, d.logger, w, r)
	_ = session.Values["User"].(models.User)

	w.WriteHeader(http.StatusOK)
}
