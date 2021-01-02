package home

import (
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/app/helpers"
)

func (h *Home) Login(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(h.sessionStore, h.logger, w, r)
	if _, ok := session.Values["profile"]; ok {
		http.Redirect(w, r, helpers.DashboardRoute(), http.StatusFound)
	}

	helpers.SetSessionHandler(helpers.HomeLogin, session, w, r)
	helpers.RenderTemplate(w, helpers.BareLayout, helpers.TemplateList(helpers.HomeLogin), session.Values)
}
