package forecast

import (
	"encoding/json"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/mux"
)

func (f *Forecast) UnitsSold(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(f.sessionStore, f.logger, w, r)
	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.ForecastTotalSales), session.Values)
}

func (f *Forecast) UnitsSoldByDate(w http.ResponseWriter, r *http.Request) {
	_ = helpers.GetSessionUser(helpers.GetSession(f.sessionStore, f.logger, w, r))

	params := mux.Vars(r)
	dateRange := params["dateRange"]

	ok, _ := helpers.ValidateInput(helpers.ValidateDateRange{DateRange: dateRange}, &f.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var byDate types.Forecast
	js, err := json.Marshal(byDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
