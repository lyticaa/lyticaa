package dashboard

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/helpers"
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
