package account

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (a *Account) Notifications(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)

	t := []string{
		helpers.NavForSession(helpers.IsSubscribed(a.sessionStore, a.logger, w, r)),
		"account/notifications",
		"partials/_filters",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (a *Account) NotificationsByDate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(a.sessionStore, a.logger, w, r)
	user := session.Values["User"].(models.User)

	notifications := models.FindNotificationsByUser(user.Id, helpers.BuildFilter(r), a.db)
	var byDate types.Notifications

	for _, notification := range *notifications {
		t := types.NotificationTable{
			Notification: notification.Notification,
			Date:         notification.CreatedAt,
		}

		byDate.Data = append(byDate.Data, t)
	}

	byDate.RecordsTotal = models.TotalNotificationsByUser(user.Id, a.db)

	js, err := json.Marshal(byDate)
	if err != nil {
		a.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
