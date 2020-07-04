package dashboard

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/types"

	"github.com/gorilla/mux"
)

func (d *Dashboard) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(d.sessionStore, d.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"partials/filters/_filters",
		"partials/filters/_date",
		"partials/filters/_import",
		"dashboard/overview",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (d *Dashboard) MetricsByDate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(d.sessionStore, d.logger, w, r)
	user := session.Values["User"].(models.User)

	params := mux.Vars(r)
	dateRange := params["dateRange"]

	ok, _ := helpers.ValidateInput(helpers.ValidateDateRange{DateRange: dateRange}, &d.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var dashboard types.Dashboard
	d.data.Dashboard(user.UserId, dateRange, &dashboard)

	js, err := json.Marshal(dashboard)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
