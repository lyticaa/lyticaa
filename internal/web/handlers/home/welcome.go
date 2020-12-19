package home

import (
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
)

func (h *Home) Welcome(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(h.sessionStore, h.logger, w, r)
	user := helpers.GetSessionUser(session)

	if user.SetupCompleted {
		http.Redirect(w, r, helpers.DashboardRoute(), http.StatusFound)
	}

	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.HomeWelcome), session.Values)
}
