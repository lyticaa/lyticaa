package account

import (
	"net/http"
	"strconv"

	"github.com/lyticaa/lyticaa-app/internal/app/helpers"
	"github.com/lyticaa/lyticaa-app/internal/app/pkg/accounts"
)

func (a *Account) SetupCompleted(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(a.sessionStore, a.logger, w, r))

	setupCompleted, _ := strconv.ParseBool(r.FormValue("setup_completed"))
	preferences := make(map[string]interface{})
	preferences["setup_completed"] = setupCompleted

	if err := accounts.UpdatePreferences(r.Context(), user.ID, preferences, a.db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *Account) MailingList(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(a.sessionStore, a.logger, w, r))

	mailingList, _ := strconv.ParseBool(r.FormValue("mailing_list"))
	preferences := make(map[string]interface{})
	preferences["mailing_list"] = mailingList

	if err := accounts.UpdatePreferences(r.Context(), user.ID, preferences, a.db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
