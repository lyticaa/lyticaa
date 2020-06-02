package setup

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
)

func (s *Setup) ImportData(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsSubscribed(s.sessionStore, s.logger, w, r) {
		http.Redirect(w, r, "/setup/subscribe/cancel", http.StatusSeeOther)
	}

	session := helpers.GetSession(s.sessionStore, s.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_setup",
		"partials/nav/account/_setup",
		"setup/import_data",
	}
	helpers.RenderTemplate(w, t, session.Values)
}
