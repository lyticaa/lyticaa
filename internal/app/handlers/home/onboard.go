package home

import (
	"net/http"
	"strings"

	"github.com/lyticaa/lyticaa/internal/app/helpers"
	"github.com/lyticaa/lyticaa/internal/app/pkg/accounts"
)

func (h *Home) Onboard(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(h.sessionStore, h.logger, w, r)
	user := helpers.GetSessionUser(session)

	accountPreferences := accounts.Preferences(r.Context(), user.ID, h.db)
	if accountPreferences.SetupCompleted {
		http.Redirect(w, r, helpers.DashboardRoute(), http.StatusFound)
	}

	session.Values["Class"] = strings.Replace(helpers.HomeOnboard, "/", "-", -1)
	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.HomeOnboard), session.Values)
}
