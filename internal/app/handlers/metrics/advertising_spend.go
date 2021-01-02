package metrics

import (
	"encoding/json"
	"net/http"

	"github.com/lyticaa/lyticaa-app/internal/app/helpers"
	"github.com/lyticaa/lyticaa-app/internal/app/types"

	"github.com/gorilla/mux"
)

func (m *Metrics) AdvertisingSpend(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)
	helpers.RenderTemplate(w, helpers.AppLayout, helpers.TemplateList(helpers.MetricsAdvertisingSpend), session.Values)
}

func (m *Metrics) AdvertisingSpendByDate(w http.ResponseWriter, r *http.Request) {
	_ = helpers.GetSessionUser(helpers.GetSession(m.sessionStore, m.logger, w, r))

	params := mux.Vars(r)
	dateRange := params["dateRange"]

	ok, _ := helpers.ValidateInput(helpers.ValidateDateRange{DateRange: dateRange}, &m.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var byDate types.AdvertisingSpend
	byDate.Draw = helpers.DtDraw(r)

	js, err := json.Marshal(byDate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}