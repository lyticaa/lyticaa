package account

import (
	"encoding/json"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/types"
)

func (a *Account) Notifications(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	helpers.RenderTemplate(w, helpers.TemplateList(helpers.AccountNotifications), session.Values)
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
