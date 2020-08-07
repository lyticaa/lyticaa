package account

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/types"
)

func (a *Account) Notifications(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_account",
		"partials/nav/account/_main",
		"partials/admin/_impersonate",
		"partials/filters/_filters",
		"partials/filters/_date",
		"partials/filters/_import",
		"account/notifications",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (a *Account) NotificationsByDate(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(a.sessionStore, a.logger, w, r))

	js, err := json.Marshal(a.loadNotifications(user.Id, r))
	if err != nil {
		a.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (a *Account) loadNotifications(userId int64, r *http.Request) types.Notifications {
	notifications := models.LoadNotificationsByUser(userId, helpers.BuildFilter(r), a.db)
	var byDate types.Notifications

	for _, notification := range *notifications {
		t := types.NotificationTable{
			Notification: notification.Notification,
			Date:         notification.CreatedAt,
		}

		byDate.Data = append(byDate.Data, t)
	}

	if len(byDate.Data) == 0 {
		byDate.Data = []types.NotificationTable{}
	}

	byDate.Draw = helpers.DtDraw(r)
	byDate.RecordsTotal = models.TotalNotificationsByUser(userId, a.db)

	return byDate
}
