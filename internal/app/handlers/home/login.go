package home

import (
	"net/http"

	"github.com/lyticaa/lyticaa/internal/app/helpers"
)

func (h *Home) Login(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(h.sessionStore, h.logger, w, r)
	if _, ok := session.Values["profile"]; ok {
		http.Redirect(w, r, helpers.DashboardRoute(), http.StatusFound)
	}

	if err := helpers.SetSessionHandler(helpers.HomeLogin, session, w, r); err != nil {
		h.logger.Error().Err(err).Msg("unable to set session handler")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	helpers.RenderTemplate(w, helpers.BareLayout, helpers.TemplateList(helpers.HomeLogin), session.Values)
}
