package dashboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gitlab.com/getlytica/lytica-app/internal/core/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"

	"github.com/gorilla/mux"
)

type ValidateDateRange struct {
	DateRange string `validate:"required,oneof=today last_thirty_days this_month last_month last_three_months last_six_months this_year all_time"`
}

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

	ok, _ := helpers.ValidateInput(ValidateDateRange{DateRange: dateRange}, &d.logger)
	if !ok {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var byDate types.Dashboard
	d.chartData(user.Id, dateRange, &byDate)

	byDate.UnitsSold = d.cardData(user.Id, "units_sold", dateRange)
	byDate.AmazonCosts = d.cardData(user.Id, "amazon_costs", dateRange)
	byDate.AdvertisingSpend = d.cardData(user.Id, "advertising_spend", dateRange)
	byDate.Refunds = d.cardData(user.Id, "refunds", dateRange)
	byDate.ShippingCredits = d.cardData(user.Id, "shipping_credits", dateRange)
	byDate.PromotionalRebates = d.cardData(user.Id, "promotional_rebates", dateRange)
	byDate.TotalCosts = d.cardData(user.Id, "total_costs", dateRange)
	byDate.NetMargin = d.cardData(user.Id, "net_margin", dateRange)

	js, err := json.Marshal(byDate)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to marshal data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func (d *Dashboard) chartData(userId int64, dateRange string, byDate *types.Dashboard) {
	var categories []string
	series := make(map[string][]string)

	totalSales := models.LoadSummary(userId, "total_sales", dateRange, d.db)
	for _, sales := range *totalSales {
		categories = append(categories, fmt.Sprintf("%v", helpers.DateFormat(dateRange, sales.OrderDate)))
		series[sales.Marketplace] = append(series[sales.Marketplace], fmt.Sprintf("%v", sales.Total))
	}

	byDate.TotalSales.Line.Categories = append(
		byDate.TotalSales.Line.Categories,
		types.Category{Category: strings.Join(categories, "|")},
	)

	for marketplace, data := range series {
		dataSet := types.DataSet{
			SeriesName: marketplace,
			Data:       strings.Join(data, "|"),
		}

		byDate.TotalSales.Line.DataSets = append(byDate.TotalSales.Line.DataSets, dataSet)
	}

	if len(byDate.TotalSales.Line.DataSets) == 0 {
		byDate.TotalSales.Line.DataSets = append(byDate.TotalSales.Line.DataSets, types.DataSet{})
	}
}

func (d *Dashboard) cardData(userId int64, view, dateRange string) types.Card {
	card := types.Card{}
	current := models.LoadSummary(userId, view, dateRange, d.db)

	var currentTotal int64
	for _, item := range *current {
		currentTotal += int64(item.Total)
		card.Chart.Sparkline.Data = append(
			card.Chart.Sparkline.Data, types.SparklineData{Value: item.Total},
		)
	}

	card.Value = currentTotal

	if len(card.Chart.Sparkline.Data) == 0 {
		card.Chart.Sparkline.Data = []types.SparklineData{}
	}

	var previous *[]models.Summary
	if !helpers.IsDateRangeAllTime(dateRange) {
		previous = models.LoadSummary(userId, view, helpers.PreviousDateRangeLabel(dateRange), d.db)

		var previousTotal int64
		for _, item := range *previous {
			previousTotal += int64(item.Total)
		}

		card.Diff = helpers.PercentDiff(currentTotal, previousTotal)
	}

	return card
}
