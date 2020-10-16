package dashboard

import (
	"encoding/json"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/mux"
)

func (d *Dashboard) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(d.sessionStore, d.logger, w, r)
	helpers.RenderTemplate(w, helpers.TemplateList(helpers.DashboardOverview), session.Values)
}

func (d *Dashboard) MetricsByDate(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(d.sessionStore, d.logger, w, r))

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
