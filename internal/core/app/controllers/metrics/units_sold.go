package metrics

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"

	"github.com/gorilla/mux"
)

func (m *Metrics) UnitsSold(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"partials/_filters",
		"metrics/units_sold",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (m *Metrics) UnitsSoldByDate(w http.ResponseWriter, r *http.Request) {
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

	var byDate types.UnitsSold

	summary := m.summaryData(dateRange, helpers.UnitsSoldView, current)
	m.chartData(dateRange, summary, &byDate.Metrics)
	m.paintUnitsSoldTable(summary, &byDate)

	transaction := models.Transaction{User: user}
	byDate.RecordsTotal = transaction.Count(dateRange, m.db)
	byDate.RecordsFiltered = byDate.RecordsTotal
	byDate.Draw = helpers.DtDraw(r)

	js, err := json.Marshal(byDate)
	if err != nil {
		m.logger.Error().Err(err).Msg("unable to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (m *Metrics) paintUnitsSoldTable(summary *[]types.Summary, byDate *types.UnitsSold) {
	if len(*summary) == 0 {
		byDate.Data = []types.UnitsSoldTable{}
		byDate.RecordsTotal = 0
		byDate.RecordsFiltered = 0
		return
	}

	for _, txn := range *summary {
		byDate.Data = append(byDate.Data, types.UnitsSoldTable{
			SKU:            txn.SKU,
			Description:    txn.Description,
			Marketplace:    txn.Marketplace,
			TotalUnitsSold: int64(txn.Total),
		})
	}
}
