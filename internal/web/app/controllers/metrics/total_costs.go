package metrics

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"

	"github.com/gorilla/mux"
)

func (m *Metrics) TotalCosts(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"partials/_filters",
		"metrics/total_costs",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (m *Metrics) TotalCostsByDate(w http.ResponseWriter, r *http.Request) {
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

	var byDate types.TotalCosts

	summary := m.summaryData(dateRange, helpers.TotalCostsView, current, &[]models.SponsoredProduct{})
	m.chartData(dateRange, summary, &byDate.Metrics)
	m.paintTotalCostsTable(summary, &byDate)

	byDate.RecordsTotal = models.TotalTransactions(user.Id, dateRange, m.db)
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

func (m *Metrics) paintTotalCostsTable(summary *[]types.Summary, byDate *types.TotalCosts) {
	if len(*summary) == 0 {
		byDate.Data = []types.TotalCostsTable{}
		byDate.RecordsTotal = 0
		byDate.RecordsFiltered = 0
		return
	}

	for _, txn := range *summary {
		byDate.Data = append(byDate.Data, types.TotalCostsTable{
			SKU:                  txn.SKU,
			Description:          txn.Description,
			Marketplace:          txn.Marketplace,
			AmazonCosts:          txn.AmazonCosts,
			ProductCosts:         txn.ProductCosts,
			ProductCostPerUnit:   txn.ProductCostsUnit,
			TotalCosts:           txn.Total,
			TotalCostsPercentage: txn.TotalCostsPercentage,
		})
	}
}
