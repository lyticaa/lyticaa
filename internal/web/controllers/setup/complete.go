package setup

import (
	"net/http"

	"gitlab.com/lyticaa/lyticaa-app/internal/models"
	"gitlab.com/lyticaa/lyticaa-app/internal/web/helpers"
)

func (s *Setup) Complete(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsSubscribed(s.sessionStore, s.logger, w, r) {
		http.Redirect(w, r, "/setup/subscribe/cancel", http.StatusSeeOther)
	}

	session := helpers.GetSession(s.sessionStore, s.logger, w, r)

	user := session.Values["User"].(models.User)
	user.SetupCompleted = true

	err := user.Save(s.db)
	if err != nil {
		s.logger.Error().Err(err).Msg("unable to save the user")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["User"] = user
	err = session.Save(r, w)
	if err != nil {
		s.logger.Error().Err(err).Msg("unable to save the session")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := []string{
		"partials/_nav",
		"partials/nav/_setup",
		"partials/nav/account/_account",
		"partials/nav/account/_setup",
		"setup/complete",
	}
	helpers.RenderTemplate(w, t, session.Values)
}
