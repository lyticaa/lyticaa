package metrics

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"

	"github.com/gorilla/mux"
)

func (m *Metrics) ProductCosts(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"partials/_filters",
		"metrics/product_costs",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (m *Metrics) ProductCostsByDate(w http.ResponseWriter, r *http.Request) {
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

	var byDate types.ProductCosts

	summary := m.summaryData(dateRange, helpers.ProductCostsView, current, &[]models.SponsoredProduct{})
	m.chartData(dateRange, summary, &byDate.Metrics)
	m.paintProductCostsTable(summary, &byDate)

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

func (m *Metrics) paintProductCostsTable(summary *[]types.Summary, byDate *types.ProductCosts) {
	if len(*summary) == 0 {
		byDate.Data = []types.ProductCostsTable{}
		byDate.RecordsTotal = 0
		byDate.RecordsFiltered = 0
		return
	}

	for _, txn := range *summary {
		byDate.Data = append(byDate.Data, types.ProductCostsTable{
			SKU:               txn.SKU,
			Description:       txn.Description,
			Marketplace:       txn.Marketplace,
			QuantitySold:      txn.QuantitySold,
			ProductCosts:      txn.ProductCosts,
			AdvertisingSpend:  txn.AdvertisingSpend,
			Refunds:           txn.Refunds,
			TotalProductCosts: txn.Total,
		})
	}
}
