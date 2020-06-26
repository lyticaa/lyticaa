package dashboard

import (
	"encoding/json"
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"

	"github.com/gorilla/mux"
)

func (d *Dashboard) Overview(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(d.sessionStore, d.logger, w, r)

	t := []string{
		"partials/_nav",
		"partials/nav/_main",
		"partials/nav/account/_main",
		"dashboard/overview",
		"partials/_filters",
	}
	helpers.RenderTemplate(w, t, session.Values)
}

func (d *Dashboard) MetricsByDate(w http.ResponseWriter, r *http.Request) {
	session := helpers.GetSession(d.sessionStore, d.logger, w, r)
	user := session.Values["User"].(models.User)

	params := mux.Vars(r)
	dateRange := params["dateRange"]

	ok, _ := helpers.ValidateInput(helpers.ValidateDateRange{DateRange: dateRange}, &d.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	current := d.amazon.LoadTransactions(user.Id, dateRange)

	var byDate types.Dashboard
	d.chartData(user.Id, dateRange, current, &byDate)

	var previous []models.Transaction
	if !helpers.IsDateRangeAllTime(dateRange) {
		previous = *d.amazon.LoadTransactions(user.Id, helpers.PreviousDateRangeLabel(dateRange))
	} else {
		previous = *current
	}

	d.cards(user.Id, dateRange, current, &previous, &byDate)

	js, err := json.Marshal(byDate)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (d *Dashboard) chartData(userId int64, dateRange string, current *[]models.Transaction, byDate *types.Dashboard) {
	summary := d.amazon.TotalSales(current)
	byDate.TotalSales = d.chart.Line(&summary, dateRange)
}

func (d *Dashboard) cards(userId int64, dateRange string, current, previous *[]models.Transaction, byDate *types.Dashboard) {
	byDate.UnitsSold = d.cardData(userId, helpers.UnitsSoldView, dateRange, current, previous)
	byDate.AmazonCosts = d.cardData(userId, helpers.AmazonCostsView, dateRange, current, previous)
	byDate.ProductCosts = d.cardData(userId, helpers.ProductCostsView, dateRange, current, previous)
	byDate.AdvertisingSpend = d.cardData(userId, helpers.AdvertisingSpendView, dateRange, current, previous)
	byDate.Refunds = d.cardData(userId, helpers.RefundsView, dateRange, current, previous)
	byDate.ShippingCredits = d.cardData(userId, helpers.ShippingCreditsView, dateRange, current, previous)
	byDate.PromotionalRebates = d.cardData(userId, helpers.PromotionalRebatesView, dateRange, current, previous)
	byDate.TotalCosts = d.cardData(userId, helpers.TotalCostsView, dateRange, current, previous)
	byDate.GrossMargin = d.cardData(userId, helpers.GrossMarginView, dateRange, current, previous)
	byDate.NetMargin = d.cardData(userId, helpers.NetMarginView, dateRange, current, previous)
}

func (d *Dashboard) cardData(userId int64, view, dateRange string, current, previous *[]models.Transaction) types.Card {
	var (
		currentValues  = make([]types.Summary, 1)
		previousValues = make([]types.Summary, 1)
	)

	switch view {
	case helpers.UnitsSoldView:
		currentValues = d.amazon.UnitsSold(current)
		previousValues = d.amazon.UnitsSold(previous)
	case helpers.AmazonCostsView:
		currentValues = d.amazon.AmazonCosts(current)
		previousValues = d.amazon.AmazonCosts(previous)
	case helpers.ProductCostsView:
		currentValues = d.amazon.ProductCosts(current)
		previousValues = d.amazon.ProductCosts(previous)
	case helpers.AdvertisingSpendView:
		currentValues = d.amazon.AdvertisingSpend(current)
		previousValues = d.amazon.AdvertisingSpend(previous)
	case helpers.RefundsView:
		currentValues = d.amazon.Refunds(current)
		previousValues = d.amazon.Refunds(previous)
	case helpers.ShippingCreditsView:
		currentValues = d.amazon.ShippingCredits(current)
		previousValues = d.amazon.ShippingCredits(previous)
	case helpers.PromotionalRebatesView:
		currentValues = d.amazon.PromotionalRebates(current)
		previousValues = d.amazon.PromotionalRebates(previous)
	case helpers.TotalCostsView:
		currentValues = d.amazon.TotalCosts(current)
		previousValues = d.amazon.TotalCosts(previous)
	case helpers.GrossMarginView:
		currentValues = d.amazon.GrossMargin(current)
		previousValues = d.amazon.GrossMargin(previous)
	case helpers.NetMarginView:
		currentValues = d.amazon.NetMargin(current)
		previousValues = d.amazon.NetMargin(previous)
	}

	card := helpers.PaintCard(&currentValues, &previousValues)
	card.Chart = d.chart.Sparkline(&currentValues)

	return card
}
