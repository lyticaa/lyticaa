package data

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/types"
)

func (d *Data) MetricsTotalSales(userId, dateRange string, metric *types.TotalSales, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, models.TotalSales, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userId, dateRange, models.TotalSales), models.TotalSales), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.TotalSalesTable{
			SKU:         item.SKU,
			Description: item.Description,
			Marketplace: item.Marketplace,
			TotalSales:  item.TotalSales,
		})
	}

	total := d.totalMetrics(userId, dateRange, models.TotalSales)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.TotalSalesTable{}
	}
}

func (d *Data) MetricsUnitsSold(userId, dateRange string, metric *types.UnitsSold, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, models.UnitsSold, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userId, dateRange, models.UnitsSold), models.UnitsSold), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.UnitsSoldTable{
			SKU:         item.SKU,
			Description: item.Description,
			Marketplace: item.Marketplace,
			Quantity:    item.Quantity,
		})
	}

	total := d.totalMetrics(userId, dateRange, models.UnitsSold)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.UnitsSoldTable{}
	}
}

func (d *Data) MetricsAmazonCosts(userId, dateRange string, metric *types.AmazonCosts, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, models.AmazonCosts, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userId, dateRange, models.AmazonCosts), models.AmazonCosts), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.AmazonCostsTable{
			SKU:         item.SKU,
			Description: item.Description,
			Marketplace: item.Marketplace,
			AmazonCosts: item.AmazonCosts,
		})
	}

	total := d.totalMetrics(userId, dateRange, models.AmazonCosts)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.AmazonCostsTable{}
	}
}

func (d *Data) MetricsProductCosts(userId, dateRange string, metric *types.ProductCosts, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, models.ProductCosts, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userId, dateRange, models.ProductCosts), models.ProductCosts), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.ProductCostsTable{
			SKU:              item.SKU,
			Description:      item.Description,
			Marketplace:      item.Marketplace,
			Quantity:         item.Quantity,
			ProductCosts:     item.ProductCosts,
			AdvertisingSpend: item.AdvertisingSpend,
			Refunds:          item.Refunds,
			TotalCosts:       item.TotalCosts,
		})
	}

	total := d.totalMetrics(userId, dateRange, models.ProductCosts)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.ProductCostsTable{}
	}
}

func (d *Data) MetricsAdvertisingSpend(userId, dateRange string, metric *types.AdvertisingSpend, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, models.AdvertisingSpend, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userId, dateRange, models.AdvertisingSpend), models.AdvertisingSpend), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.AdvertisingSpendTable{
			SKU:                        item.SKU,
			Description:                item.Description,
			Marketplace:                item.Marketplace,
			AdvertisingSpend:           item.AdvertisingSpend,
			AdvertisingSpendPercentage: item.AdvertisingSpendPercentage,
		})
	}

	total := d.totalMetrics(userId, dateRange, models.AdvertisingSpend)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.AdvertisingSpendTable{}
	}
}

func (d *Data) MetricsRefunds(userId, dateRange string, metric *types.Refunds, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, models.Refunds, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userId, dateRange, models.Refunds), models.Refunds), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.RefundsTable{
			SKU:               item.SKU,
			Description:       item.Description,
			Marketplace:       item.Marketplace,
			Refunds:           item.Refunds,
			RefundsPercentage: item.RefundsPercentage,
		})
	}

	total := d.totalMetrics(userId, dateRange, models.Refunds)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.RefundsTable{}
	}
}

func (d *Data) MetricsShippingCredits(userId, dateRange string, metric *types.ShippingCredits, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, models.ShippingCredits, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userId, dateRange, models.ShippingCredits), models.ShippingCredits), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.ShippingCreditsTable{
			SKU:                item.SKU,
			Description:        item.Description,
			Marketplace:        item.Marketplace,
			ShippingCredits:    item.ShippingCredits,
			ShippingCreditsTax: item.ShippingCreditsTax,
		})
	}

	total := d.totalMetrics(userId, dateRange, models.ShippingCredits)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.ShippingCreditsTable{}
	}
}

func (d *Data) MetricsPromotionalRebates(userId, dateRange string, metric *types.PromotionalRebates, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, models.PromotionalRebates, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userId, dateRange, models.PromotionalRebates), models.PromotionalRebates), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.PromotionalRebatesTable{
			SKU:                   item.SKU,
			Description:           item.Description,
			Marketplace:           item.Marketplace,
			PromotionalRebates:    item.PromotionalRebates,
			PromotionalRebatesTax: item.PromotionalRebatesTax,
		})
	}

	total := d.totalMetrics(userId, dateRange, models.PromotionalRebates)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.PromotionalRebatesTable{}
	}
}

func (d *Data) MetricsTotalCosts(userId, dateRange string, metric *types.TotalCosts, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, models.TotalCosts, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userId, dateRange, models.TotalCosts), models.TotalCosts), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.TotalCostsTable{
			SKU:                  item.SKU,
			Description:          item.Description,
			Marketplace:          item.Marketplace,
			AmazonCosts:          item.AmazonCosts,
			ProductCosts:         item.ProductCosts,
			ProductCostPerUnit:   item.ProductCostsUnit,
			TotalCosts:           item.TotalCosts,
			TotalCostsPercentage: item.TotalCostsPercentage,
		})
	}

	total := d.totalMetrics(userId, dateRange, models.TotalCosts)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.TotalCostsTable{}
	}
}

func (d *Data) MetricsGrossMargin(userId, dateRange string, metric *types.GrossMargin, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, models.GrossMargin, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userId, dateRange, models.GrossMargin), models.GrossMargin), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.GrossMarginTable{
			SKU:                item.SKU,
			Description:        item.Description,
			Marketplace:        item.Marketplace,
			ProductCosts:       item.ProductCosts,
			Quantity:           item.Quantity,
			TotalSales:         item.TotalSales,
			AmazonCosts:        item.AmazonCosts,
			ShippingCredits:    item.ShippingCredits,
			PromotionalRebates: item.PromotionalRebates,
			GrossMargin:        item.GrossMargin,
			SalesTaxCollected:  item.SalesTaxCollected,
			TotalCollected:     item.TotalCollected,
		})
	}

	total := d.totalMetrics(userId, dateRange, models.GrossMargin)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.GrossMarginTable{}
	}
}

func (d *Data) MetricsNetMargin(userId, dateRange string, metric *types.NetMargin, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, models.NetMargin, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(d.loadMetricsSummary(userId, dateRange, models.NetMargin), models.NetMargin), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.NetMarginTable{
			SKU:           item.SKU,
			Description:   item.Description,
			Marketplace:   item.Marketplace,
			Quantity:      item.Quantity,
			GrossMargin:   item.GrossMargin,
			TotalCosts:    item.TotalCosts,
			NetMargin:     item.NetMargin,
			NetMarginUnit: item.NetMarginUnit,
			ROI:           item.ROI,
		})
	}

	total := d.totalMetrics(userId, dateRange, models.NetMargin)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.NetMarginTable{}
	}
}

func (d *Data) loadMetrics(userId, dateRange, view string, filter *models.Filter) *[]models.Metric {
	return models.LoadMetrics(userId, dateRange, view, filter, d.db)
}

func (d *Data) loadMetricsSummary(userId, dateRange, view string) *[]models.Metric {
	switch view {
	case models.TotalSales:
		return models.LoadMetricsTotalSalesSummary(userId, dateRange, d.db)
	case models.UnitsSold:
		return models.LoadMetricsUnitsSoldSummary(userId, dateRange, d.db)
	case models.AmazonCosts:
		return models.LoadMetricsAmazonCostsSummary(userId, dateRange, d.db)
	case models.ProductCosts:
		return models.LoadMetricsProductCostsSummary(userId, dateRange, d.db)
	case models.AdvertisingSpend:
		return models.LoadMetricsAdvertisingSpendSummary(userId, dateRange, d.db)
	case models.Refunds:
		return models.LoadMetricsRefundsSummary(userId, dateRange, d.db)
	case models.ShippingCredits:
		return models.LoadMetricsShippingCreditsSummary(userId, dateRange, d.db)
	case models.PromotionalRebates:
		return models.LoadMetricsPromotionalRebatesSummary(userId, dateRange, d.db)
	case models.TotalCosts:
		return models.LoadMetricsTotalCostsSummary(userId, dateRange, d.db)
	case models.GrossMargin:
		return models.LoadMetricsGrossMarginSummary(userId, dateRange, d.db)
	case models.NetMargin:
		return models.LoadMetricsNetMarginSummary(userId, dateRange, d.db)
	default:
		return &[]models.Metric{}
	}
}

func (d *Data) totalMetrics(userId, dateRange, view string) int64 {
	return models.TotalMetrics(userId, dateRange, view, d.db)
}

func (d *Data) metricsSummary(metrics *[]models.Metric, view string) *[]types.Summary {
	var summary []types.Summary
	for _, metric := range *metrics {
		summary = append(summary, types.Summary{
			Date:        metric.DateTime,
			Marketplace: metric.Marketplace,
			Total:       d.metricsItem(metric, view),
		})
	}

	return &summary
}

func (d *Data) metricsItem(metric models.Metric, view string) float64 {
	switch view {
	case models.TotalSales:
		return metric.TotalSales
	case models.UnitsSold:
		return float64(metric.Quantity)
	case models.AmazonCosts:
		return metric.AmazonCosts
	case models.ProductCosts:
		return metric.ProductCosts
	case models.AdvertisingSpend:
		return metric.AdvertisingSpend
	case models.Refunds:
		return metric.Refunds
	case models.ShippingCredits:
		return metric.ShippingCredits
	case models.PromotionalRebates:
		return metric.PromotionalRebates
	case models.TotalCosts:
		return metric.TotalCosts
	case models.GrossMargin:
		return metric.GrossMargin
	case models.NetMargin:
		return metric.NetMargin
	default:
		return 0.0
	}
}
