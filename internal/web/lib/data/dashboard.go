package data

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/types"
)

func (d *Data) Dashboard(userId, dateRange string, dashboard *types.Dashboard) {
	var summary []types.Summary

	totalSales := models.LoadDashboardTotalSales(userId, dateRange, d.db)
	for _, sales := range *totalSales {
		summary = append(summary, types.Summary{
			Date:        sales.DateTime,
			Marketplace: sales.Marketplace,
			Total:       sales.TotalSales,
		})
	}

	dashboard.TotalSales = d.chart.Line(&summary, dateRange)
	current := models.LoadDashboard(userId, dateRange, d.db)

	var previous []models.Dashboard
	if !helpers.IsDateRangeAllTime(dateRange) {
		previous = *models.LoadDashboard(userId, helpers.PreviousDateRangeLabel(dateRange), d.db)
	}

	d.dashboardUnitsSold(current, &previous, dashboard)
	d.dashboardAmazonCosts(current, &previous, dashboard)
	d.dashboardProductCosts(current, &previous, dashboard)
	d.dashboardAdvertisingSpend(current, &previous, dashboard)
	d.dashboardRefunds(current, &previous, dashboard)
	d.dashboardShippingCredits(current, &previous, dashboard)
	d.dashboardPromotionalRebates(current, &previous, dashboard)
	d.dashboardTotalCosts(current, &previous, dashboard)
	d.dashboardGrossMargin(current, &previous, dashboard)
	d.dashboardNetMargin(current, &previous, dashboard)
}

func (d *Data) dashboardUnitsSold(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.UnitsSold.Total = d.dashboardCard(current, previous, unitsSold)
	dashboard.UnitsSold.Chart = d.chart.Sparkline(d.dashboardSummary(current, unitsSold))
}

func (d *Data) dashboardAmazonCosts(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.AmazonCosts.Total = d.dashboardCard(current, previous, amazonCosts)
	dashboard.AmazonCosts.Chart = d.chart.Sparkline(d.dashboardSummary(current, amazonCosts))
}

func (d *Data) dashboardProductCosts(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.ProductCosts.Total = d.dashboardCard(current, previous, productCosts)
	dashboard.ProductCosts.Chart = d.chart.Sparkline(d.dashboardSummary(current, productCosts))
}

func (d *Data) dashboardAdvertisingSpend(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.AdvertisingSpend.Total = d.dashboardCard(current, previous, advertisingSpend)
	dashboard.AdvertisingSpend.Chart = d.chart.Sparkline(d.dashboardSummary(current, advertisingSpend))
}

func (d *Data) dashboardRefunds(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.Refunds.Total = d.dashboardCard(current, previous, refunds)
	dashboard.Refunds.Chart = d.chart.Sparkline(d.dashboardSummary(current, refunds))
}

func (d *Data) dashboardShippingCredits(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.ShippingCredits.Total = d.dashboardCard(current, previous, shippingCredits)
	dashboard.ShippingCredits.Chart = d.chart.Sparkline(d.dashboardSummary(current, shippingCredits))
}

func (d *Data) dashboardPromotionalRebates(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.PromotionalRebates.Total = d.dashboardCard(current, previous, promotionalRebates)
	dashboard.PromotionalRebates.Chart = d.chart.Sparkline(d.dashboardSummary(current, promotionalRebates))
}

func (d *Data) dashboardTotalCosts(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.TotalCosts.Total = d.dashboardCard(current, previous, totalCosts)
	dashboard.TotalCosts.Chart = d.chart.Sparkline(d.dashboardSummary(current, totalCosts))
}

func (d *Data) dashboardGrossMargin(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.GrossMargin.Total = d.dashboardCard(current, previous, grossMargin)
	dashboard.GrossMargin.Chart = d.chart.Sparkline(d.dashboardSummary(current, grossMargin))
}

func (d *Data) dashboardNetMargin(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.NetMargin.Total = d.dashboardCard(current, previous, netMargin)
	dashboard.NetMargin.Chart = d.chart.Sparkline(d.dashboardSummary(current, netMargin))
}

func (d *Data) dashboardCard(current, previous *[]models.Dashboard, view string) types.Total {
	var currentTotal float64
	for _, item := range *current {
		currentTotal += d.dashboardItem(view, item)
	}

	var previousTotal float64
	for _, item := range *previous {
		previousTotal += d.dashboardItem(view, item)
	}

	return types.Total{Value: currentTotal, Diff: helpers.PercentDiff(int64(currentTotal), int64(previousTotal))}
}

func (d *Data) dashboardSummary(current *[]models.Dashboard, view string) *[]types.Summary {
	var summary []types.Summary
	for _, item := range *current {
		summary = append(summary, types.Summary{
			Date:        item.DateTime,
			Marketplace: item.Marketplace,
			Total:       d.dashboardItem(view, item),
		})
	}

	return &summary
}

func (d *Data) dashboardItem(view string, item models.Dashboard) float64 {
	switch view {
	case unitsSold:
		return float64(item.UnitsSold)
	default:
		return 0.0
	}
}
