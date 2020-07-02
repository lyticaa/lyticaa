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
	dashboard.UnitsSold.Total = d.dashboardCard(current, previous, models.UnitsSold)
	dashboard.UnitsSold.Chart = d.chart.Sparkline(d.dashboardSummary(current, models.UnitsSold))
}

func (d *Data) dashboardAmazonCosts(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.AmazonCosts.Total = d.dashboardCard(current, previous, models.AmazonCosts)
	dashboard.AmazonCosts.Chart = d.chart.Sparkline(d.dashboardSummary(current, models.AmazonCosts))
}

func (d *Data) dashboardProductCosts(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.ProductCosts.Total = d.dashboardCard(current, previous, models.ProductCosts)
	dashboard.ProductCosts.Chart = d.chart.Sparkline(d.dashboardSummary(current, models.ProductCosts))
}

func (d *Data) dashboardAdvertisingSpend(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.AdvertisingSpend.Total = d.dashboardCard(current, previous, models.AdvertisingSpend)
	dashboard.AdvertisingSpend.Chart = d.chart.Sparkline(d.dashboardSummary(current, models.AdvertisingSpend))
}

func (d *Data) dashboardRefunds(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.Refunds.Total = d.dashboardCard(current, previous, models.Refunds)
	dashboard.Refunds.Chart = d.chart.Sparkline(d.dashboardSummary(current, models.Refunds))
}

func (d *Data) dashboardShippingCredits(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.ShippingCredits.Total = d.dashboardCard(current, previous, models.ShippingCredits)
	dashboard.ShippingCredits.Chart = d.chart.Sparkline(d.dashboardSummary(current, models.ShippingCredits))
}

func (d *Data) dashboardPromotionalRebates(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.PromotionalRebates.Total = d.dashboardCard(current, previous, models.PromotionalRebates)
	dashboard.PromotionalRebates.Chart = d.chart.Sparkline(d.dashboardSummary(current, models.PromotionalRebates))
}

func (d *Data) dashboardTotalCosts(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.TotalCosts.Total = d.dashboardCard(current, previous, models.TotalCosts)
	dashboard.TotalCosts.Chart = d.chart.Sparkline(d.dashboardSummary(current, models.TotalCosts))
}

func (d *Data) dashboardGrossMargin(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.GrossMargin.Total = d.dashboardCard(current, previous, models.GrossMargin)
	dashboard.GrossMargin.Chart = d.chart.Sparkline(d.dashboardSummary(current, models.GrossMargin))
}

func (d *Data) dashboardNetMargin(current, previous *[]models.Dashboard, dashboard *types.Dashboard) {
	dashboard.NetMargin.Total = d.dashboardCard(current, previous, models.NetMargin)
	dashboard.NetMargin.Chart = d.chart.Sparkline(d.dashboardSummary(current, models.NetMargin))
}

func (d *Data) dashboardCard(current, previous *[]models.Dashboard, card string) types.Total {
	var currentTotal float64
	for _, item := range *current {
		currentTotal += d.dashboardItem(card, item)
	}

	var previousTotal float64
	for _, item := range *previous {
		previousTotal += d.dashboardItem(card, item)
	}

	return types.Total{Value: currentTotal, Diff: helpers.PercentDiff(int64(currentTotal), int64(previousTotal))}
}

func (d *Data) dashboardSummary(current *[]models.Dashboard, card string) *[]types.Summary {
	var summary []types.Summary
	for _, item := range *current {
		summary = append(summary, types.Summary{
			Date:        item.DateTime,
			Marketplace: item.Marketplace,
			Total:       d.dashboardItem(card, item),
		})
	}

	return &summary
}

func (d *Data) dashboardItem(card string, item models.Dashboard) float64 {
	switch card {
	case models.UnitsSold:
		return float64(item.UnitsSold)
	case models.AmazonCosts:
		return item.AmazonCosts
	case models.ProductCosts:
		return item.ProductCosts
	case models.AdvertisingSpend:
		return item.AdvertisingSpend
	case models.Refunds:
		return item.Refunds
	case models.ShippingCredits:
		return item.ShippingCredits
	case models.PromotionalRebates:
		return item.PromotionalRebates
	case models.TotalCosts:
		return item.TotalCosts
	case models.GrossMargin:
		return item.GrossMargin
	case models.NetMargin:
		return item.NetMargin
	default:
		return 0.0
	}
}
