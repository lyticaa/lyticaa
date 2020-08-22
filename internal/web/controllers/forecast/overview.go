package forecast

import (
	"encoding/json"
	"net/http"

	"gitlab.com/lyticaa/lyticaa-app/internal/web/helpers"
	"gitlab.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/mux"
)

func (f *Forecast) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(f.sessionStore, f.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_account",
		"partials/nav/account/_main",
		"partials/admin/_impersonate",
		"partials/filters/_filters",
		"partials/filters/_date",
		"partials/filters/_import",
		"forecast/overview",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (f *Forecast) ForecastByDate(w http.ResponseWriter, r *http.Request) {
	user := helpers.GetSessionUser(helpers.GetSession(f.sessionStore, f.logger, w, r))

	params := mux.Vars(r)
	dateRange := params["dateRange"]

	ok, _ := helpers.ValidateInput(helpers.ValidateDateRange{DateRange: dateRange}, &f.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	view := params["view"]
	ok, _ = helpers.ValidateInput(helpers.ValidateView{View: view}, &f.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var byDate types.Forecast

	f.data.Forecast(user.UserId, dateRange, view, &byDate)
	js, err := json.Marshal(byDate)
	if err != nil {
		f.logger.Error().Err(err).Msg("failed to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
