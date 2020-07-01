package data

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/types"
)

func (d *Data) Cohorts(userId, dateRange, view string, cohort *types.Cohort, filter *models.Filter) {
	current := models.LoadCohorts(userId, dateRange, view, filter, d.db)

	var previous []models.Cohort
	if !helpers.IsDateRangeAllTime(dateRange) {
		previous = *models.LoadCohorts(userId, helpers.PreviousDateRangeLabel(dateRange), view, filter, d.db)
	}

	d.cohortsTotalSales(current, &previous, cohort)
	d.cohortsAmazonCosts(current, &previous, cohort)
	d.cohortsProductCosts(current, &previous, cohort)
	d.cohortsAdvertisingSpend(current, &previous, cohort)
	d.cohortsNetMargin(current, &previous, cohort)

	for _, item := range *current {
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
	cohort.TotalSales.Total = d.cohortCard(current, previous, totalSales)
	cohort.TotalSales.Chart = d.chart.Sparkline(d.cohortSummary(current, totalSales))
}

func (d *Data) cohortsAmazonCosts(current, previous *[]models.Cohort, cohort *types.Cohort) {
	cohort.AmazonCosts.Total = d.cohortCard(current, previous, amazonCosts)
	cohort.AmazonCosts.Chart = d.chart.Sparkline(d.cohortSummary(current, amazonCosts))
}

func (d *Data) cohortsProductCosts(current, previous *[]models.Cohort, cohort *types.Cohort) {
	cohort.ProductCosts.Total = d.cohortCard(current, previous, productCosts)
	cohort.ProductCosts.Chart = d.chart.Sparkline(d.cohortSummary(current, productCosts))
}

func (d *Data) cohortsAdvertisingSpend(current, previous *[]models.Cohort, cohort *types.Cohort) {
	cohort.AdvertisingSpend.Total = d.cohortCard(current, previous, advertisingSpend)
	cohort.AdvertisingSpend.Chart = d.chart.Sparkline(d.cohortSummary(current, advertisingSpend))
}

func (d *Data) cohortsNetMargin(current, previous *[]models.Cohort, cohort *types.Cohort) {
	cohort.NetMargin.Total = d.cohortCard(current, previous, netMargin)
	cohort.NetMargin.Chart = d.chart.Sparkline(d.cohortSummary(current, netMargin))
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
	case totalSales:
		return item.TotalSales
	case amazonCosts:
		return item.AmazonCosts
	case productCosts:
		return item.ProductCosts
	case advertisingSpend:
		return item.AdvertisingSpend
	case netMargin:
		return item.NetMargin
	default:
		return 0.0
	}
}
