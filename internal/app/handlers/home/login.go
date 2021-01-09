package home

import (
	"net/http"
	"strings"

	"github.com/lyticaa/lyticaa/internal/app/helpers"
)

func (h *Home) Login(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(h.sessionStore, h.logger, w, r)
	if _, ok := session.Values["profile"]; ok {
		http.Redirect(w, r, helpers.DashboardRoute(), http.StatusFound)
	}

	session.Values["Class"] = strings.Replace(helpers.HomeLogin, "/", "-", -1)
	helpers.RenderTemplate(w, helpers.BareLayout, helpers.TemplateList(helpers.HomeLogin), session.Values)
}
