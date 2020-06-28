package metrics

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"

	"github.com/gorilla/mux"
)

func (m *Metrics) GrossMargin(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(m.sessionStore, m.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"partials/_filters",
		"metrics/gross_margin",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (m *Metrics) GrossMarginByDate(w http.ResponseWriter, r *http.Request) {
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

	var byDate types.GrossMargin

	summary := m.summaryData(dateRange, helpers.GrossMarginView, current, &[]models.SponsoredProduct{})
	m.chartData(dateRange, summary, &byDate.Metrics)
	m.paintGrossMarginTable(summary, &byDate)

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

func (m *Metrics) paintGrossMarginTable(summary *[]types.Summary, byDate *types.GrossMargin) {
	if len(*summary) == 0 {
		byDate.Data = []types.GrossMarginTable{}
		byDate.RecordsTotal = 0
		byDate.RecordsFiltered = 0
		return
	}

	for _, txn := range *summary {
		byDate.Data = append(byDate.Data, types.GrossMarginTable{
			SKU:                  txn.SKU,
			Description:          txn.Description,
			Marketplace:          txn.Marketplace,
			ProductCosts:         txn.ProductCosts,
			QuantitySold:         txn.QuantitySold,
			TotalRevenue:         txn.TotalRevenue,
			AmazonCosts:          txn.AmazonCosts,
			ShippingCredits:      txn.ShippingCredits,
			PromotionalRebates:   txn.PromotionalRebates,
			GrossMargin:          txn.GrossMargin,
			SalesTaxCollected:    txn.SalesTaxCollected,
			TotalAmountCollected: txn.Total,
		})
	}
}
