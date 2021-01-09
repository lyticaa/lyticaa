package home

import (
	"net/http"

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

	if err := helpers.SetSessionHandler("home-onboard", session, w, r); err != nil {
		h.logger.Error().Err(err).Msg("unable to set session handler")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.HomeOnboard), session.Values)
}
