package home

import (
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/pkg/accounts"
)

func (h *Home) Welcome(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(h.sessionStore, h.logger, w, r)
	user := helpers.GetSessionUser(session)

	accountPreferences := accounts.FetchAccountPreferences(r.Context(), user.ID, h.db)
	if accountPreferences.SetupCompleted {
		http.Redirect(w, r, helpers.DashboardRoute(), http.StatusFound)
	}

	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.HomeWelcome), session.Values)
}
