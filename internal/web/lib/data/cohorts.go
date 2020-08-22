package data

import (
	"gitlab.com/lyticaa/lyticaa-app/internal/models"
	"gitlab.com/lyticaa/lyticaa-app/internal/web/helpers"
	"gitlab.com/lyticaa/lyticaa-app/internal/web/types"
)

func (d *Data) Cohorts(userId, dateRange, view string, cohort *types.Cohort, filter *models.Filter) {
	current := models.LoadCohortsSummary(userId, dateRange, view, d.db)

	var previous []models.Cohort
	if !helpers.IsDateRangeAllTime(dateRange) {
		previous = *models.LoadCohortsSummary(userId, helpers.PreviousDateRangeLabel(dateRange), view, d.db)
	}

	d.cohortsTotalSales(current, &previous, cohort)
	d.cohortsAmazonCosts(current, &previous, cohort)
	d.cohortsProductCosts(current, &previous, cohort)
	d.cohortsAdvertisingSpend(current, &previous, cohort)
	d.cohortsNetMargin(current, &previous, cohort)

	table := models.LoadCohorts(userId, dateRange, view, filter, d.db)
	for _, item := range *table {
		cohort.Data = append(cohort.Data, types.CohortTable{
			SKU:                item.SKU,
			Description:        item.Description,
			Marketplace:        item.Marketplace,
			TotalSales:         item.TotalSales,
			Quantity:           item.Quantity,
			AmazonCosts:        item.AmazonCosts,
			ProductCosts:       item.ProductCosts,
			AdvertisingSpend:   item.AdvertisingSpend,
			Refunds:            item.Refunds,
			ShippingCredits:    item.ShippingCredits,
			PromotionalRebates: item.PromotionalRebates,
			TotalCosts:         item.TotalCosts,
			NetMargin:          item.NetMargin,
		})
	}

	total := models.TotalCohorts(userId, dateRange, view, d.db)
	cohort.RecordsTotal = total
	cohort.RecordsFiltered = total

	if len(cohort.Data) == 0 {
		cohort.Data = []types.CohortTable{}
	}
}

func (d *Data) cohortsTotalSales(current, previous *[]models.Cohort, cohort *types.Cohort) {
	cohort.TotalSales.Total = d.cohortCard(current, previous, models.TotalSales)
	cohort.TotalSales.Chart = d.chart.Sparkline(d.cohortSummary(current, models.TotalSales))
}

func (d *Data) cohortsAmazonCosts(current, previous *[]models.Cohort, cohort *types.Cohort) {
	cohort.AmazonCosts.Total = d.cohortCard(current, previous, models.AmazonCosts)
	cohort.AmazonCosts.Chart = d.chart.Sparkline(d.cohortSummary(current, models.AmazonCosts))
}

func (d *Data) cohortsProductCosts(current, previous *[]models.Cohort, cohort *types.Cohort) {
	cohort.ProductCosts.Total = d.cohortCard(current, previous, models.ProductCosts)
	cohort.ProductCosts.Chart = d.chart.Sparkline(d.cohortSummary(current, models.ProductCosts))
}

func (d *Data) cohortsAdvertisingSpend(current, previous *[]models.Cohort, cohort *types.Cohort) {
	cohort.AdvertisingSpend.Total = d.cohortCard(current, previous, models.AdvertisingSpend)
	cohort.AdvertisingSpend.Chart = d.chart.Sparkline(d.cohortSummary(current, models.AdvertisingSpend))
}

func (d *Data) cohortsNetMargin(current, previous *[]models.Cohort, cohort *types.Cohort) {
	cohort.NetMargin.Total = d.cohortCard(current, previous, models.NetMargin)
	cohort.NetMargin.Chart = d.chart.Sparkline(d.cohortSummary(current, models.NetMargin))
}

func (d *Data) cohortCard(current, previous *[]models.Cohort, card string) types.Total {
	var currentTotal float64
	for _, item := range *current {
		currentTotal += d.cohortItem(card, item)
	}

	var previousTotal float64
	for _, item := range *previous {
		previousTotal += d.cohortItem(card, item)
	}

	return types.Total{Value: currentTotal, Diff: helpers.PercentDiff(int64(currentTotal), int64(previousTotal))}
}

func (d *Data) cohortSummary(current *[]models.Cohort, card string) *[]types.Summary {
	var summary []types.Summary
	for _, item := range *current {
		summary = append(summary, types.Summary{
			Date:        item.DateTime,
			Marketplace: item.Marketplace,
			Total:       d.cohortItem(card, item),
		})
	}

	return &summary
}

func (d *Data) cohortItem(card string, item models.Cohort) float64 {
	switch card {
	case models.TotalSales:
		return item.TotalSales
	case models.AmazonCosts:
		return item.AmazonCosts
	case models.ProductCosts:
		return item.ProductCosts
	case models.AdvertisingSpend:
		return item.AdvertisingSpend
	case models.NetMargin:
		return item.NetMargin
	default:
		return 0.0
	}
}
