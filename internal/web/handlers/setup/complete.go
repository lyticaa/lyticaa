package setup

import (
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
)

func (s *Setup) Complete(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(s.sessionStore, s.logger, w, r)
	helpers.ReloadSessionUser(session, w, r, s.db)

	if !helpers.IsSubscribed(session) {
		http.Redirect(w, r, "/setup/subscribe/cancel", http.StatusSeeOther)
	}

	user := helpers.GetSessionUser(session)
	user.SetupCompleted = true

	err := user.Save(s.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	helpers.ReloadSessionUser(session, w, r, s.db)
	helpers.RenderTemplate(w, helpers.TemplateList(helpers.SetupComplete), session.Values)
}
