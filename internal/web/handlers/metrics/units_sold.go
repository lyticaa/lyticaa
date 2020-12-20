package metrics

import (
	"encoding/json"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/web/helpers"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/gorilla/mux"
)

func (m *Metrics) UnitsSold(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)
	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.MetricsUnitsSold), session.Values)
}

func (m *Metrics) UnitsSoldByDate(w http.ResponseWriter, r *http.Request) {
	_ = helpers.GetSessionUser(helpers.GetSession(m.sessionStore, m.logger, w, r))

	params := mux.Vars(r)
	dateRange := params["dateRange"]

	ok, _ := helpers.ValidateInput(helpers.ValidateDateRange{DateRange: dateRange}, &m.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var byDate types.UnitsSold
	byDate.Draw = helpers.DtDraw(r)

	js, err := json.Marshal(byDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
