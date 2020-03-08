package account

import (
	"net/http"

	"gitlab.com/getlytica/lytica/internal/core/app/helpers"
)

func (a *Account) Invitations(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)

	t := []string{helpers.NavForSession(helpers.IsSubscribed(a.sessionStore, a.logger, w, r)), "account/invitations"}
	helpers.RenderTemplate(w, t, session.Values)
}
