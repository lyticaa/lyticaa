package account

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/core/user"
)

func (a *Account) ChangePassword(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)

	u := user.NewUser(session.Values["userId"].(string), session.Values["email"].(string), a.logger)
	_ = u.ResetPassword()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
