package account

import (
	"encoding/json"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/app/helpers"
	"github.com/lyticaa/lyticaa-app/internal/app/pkg/accounts"
	"github.com/lyticaa/lyticaa-app/internal/app/types"
)

func (a *Account) Notifications(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.AccountNotifications), session.Values)
}

func (a *Account) NotificationsByDate(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(a.sessionStore, a.logger, w, r))

	var notifications types.Notifications
	accounts.Notifications(r.Context(), &notifications, helpers.BuildFilter(r), user.ID, a.db)

	notifications.Draw = helpers.DtDraw(r)

	js, err := json.Marshal(notifications)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
