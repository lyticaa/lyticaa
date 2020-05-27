package setup

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
)

func (s *Setup) Invite(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(s.sessionStore, s.logger, w, r)

	t := []string{"partials/nav/_setup", "setup/invite"}
	helpers.RenderTemplate(w, t, session.Values)
}
