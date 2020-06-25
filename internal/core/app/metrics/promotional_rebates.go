package metrics

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"

	"github.com/gorilla/mux"
)

func (m *Metrics) PromotionalRebates(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"partials/_filters",
		"metrics/promotional_rebates",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (m *Metrics) PromotionalRebatesByDate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)
	user := session.Values["User"].(models.User)

	params := mux.Vars(r)
	dateRange := params["dateRange"]

	ok, _ := helpers.ValidateInput(helpers.ValidateDateRange{DateRange: dateRange}, &m.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	current := m.amazon.LoadTransactions(user.Id, dateRange)

	var byDate types.PromotionalRebates
	m.chartData(user.Id, dateRange, helpers.PromotionalRebatesView, current, &byDate.Metrics)
	byDate.Data = []types.PromotionalRebatesTable{}

	js, err := json.Marshal(byDate)
	if err != nil {
		m.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}
