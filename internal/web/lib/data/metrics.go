package data

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/types"
)

func (d *Data) MetricsTotalSales(userId, dateRange string, metric *types.TotalSales, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, totalSales, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(metrics, totalSales), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.TotalSalesTable{
			SKU:         item.SKU,
			Description: item.Description,
			Marketplace: item.Marketplace,
			TotalSales:  item.TotalSales,
		})
	}

	total := d.totalMetrics(userId, dateRange, totalSales)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.TotalSalesTable{}
	}
}

func (d *Data) MetricsUnitsSold(userId, dateRange string, metric *types.UnitsSold, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, unitsSold, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(metrics, unitsSold), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.UnitsSoldTable{
			SKU:         item.SKU,
			Description: item.Description,
			Marketplace: item.Marketplace,
			Quantity:    item.Quantity,
		})
	}

	total := d.totalMetrics(userId, dateRange, unitsSold)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.UnitsSoldTable{}
	}
}

func (d *Data) MetricsAmazonCosts(userId, dateRange string, metric *types.AmazonCosts, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, amazonCosts, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(metrics, amazonCosts), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.AmazonCostsTable{
			SKU:         item.SKU,
			Description: item.Description,
			Marketplace: item.Marketplace,
			AmazonCosts: item.AmazonCosts,
		})
	}

	total := d.totalMetrics(userId, dateRange, amazonCosts)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.AmazonCostsTable{}
	}
}

func (d *Data) MetricsProductCosts(userId, dateRange string, metric *types.ProductCosts, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, productCosts, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(metrics, productCosts), dateRange)

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

	total := d.totalMetrics(userId, dateRange, productCosts)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.ProductCostsTable{}
	}
}

func (d *Data) MetricsAdvertisingSpend(userId, dateRange string, metric *types.AdvertisingSpend, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, advertisingSpend, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(metrics, advertisingSpend), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.AdvertisingSpendTable{
			SKU:                        item.SKU,
			Description:                item.Description,
			Marketplace:                item.Marketplace,
			AdvertisingSpend:           item.AdvertisingSpend,
			AdvertisingSpendPercentage: item.AdvertisingSpendPercentage,
		})
	}

	total := d.totalMetrics(userId, dateRange, advertisingSpend)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.AdvertisingSpendTable{}
	}
}

func (d *Data) MetricsRefunds(userId, dateRange string, metric *types.Refunds, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, refunds, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(metrics, refunds), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.RefundsTable{
			SKU:               item.SKU,
			Description:       item.Description,
			Marketplace:       item.Marketplace,
			Refunds:           item.Refunds,
			RefundsPercentage: item.RefundsPercentage,
		})
	}

	total := d.totalMetrics(userId, dateRange, refunds)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.RefundsTable{}
	}
}

func (d *Data) MetricsShippingCredits(userId, dateRange string, metric *types.ShippingCredits, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, shippingCredits, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(metrics, shippingCredits), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.ShippingCreditsTable{
			SKU:                item.SKU,
			Description:        item.Description,
			Marketplace:        item.Marketplace,
			ShippingCredits:    item.ShippingCredits,
			ShippingCreditsTax: item.ShippingCreditsTax,
		})
	}

	total := d.totalMetrics(userId, dateRange, shippingCredits)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.ShippingCreditsTable{}
	}
}

func (d *Data) MetricsPromotionalRebates(userId, dateRange string, metric *types.PromotionalRebates, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, promotionalRebates, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(metrics, promotionalRebates), dateRange)

	for _, item := range *metrics {
		metric.Data = append(metric.Data, types.PromotionalRebatesTable{
			SKU:                   item.SKU,
			Description:           item.Description,
			Marketplace:           item.Marketplace,
			PromotionalRebates:    item.PromotionalRebates,
			PromotionalRebatesTax: item.PromotionalRebatesTax,
		})
	}

	total := d.totalMetrics(userId, dateRange, promotionalRebates)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.PromotionalRebatesTable{}
	}
}

func (d *Data) MetricsTotalCosts(userId, dateRange string, metric *types.TotalCosts, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, totalCosts, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(metrics, totalCosts), dateRange)

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

	total := d.totalMetrics(userId, dateRange, totalCosts)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.TotalCostsTable{}
	}
}

func (d *Data) MetricsGrossMargin(userId, dateRange string, metric *types.GrossMargin, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, grossMargin, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(metrics, grossMargin), dateRange)

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

	total := d.totalMetrics(userId, dateRange, grossMargin)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.GrossMarginTable{}
	}
}

func (d *Data) MetricsNetMargin(userId, dateRange string, metric *types.NetMargin, filter *models.Filter) {
	metrics := d.loadMetrics(userId, dateRange, netMargin, filter)
	metric.Chart = d.chart.Line(d.metricsSummary(metrics, netMargin), dateRange)

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

	total := d.totalMetrics(userId, dateRange, netMargin)
	metric.RecordsTotal = total
	metric.RecordsFiltered = total

	if len(metric.Data) == 0 {
		metric.Data = []types.NetMarginTable{}
	}
}

func (d *Data) loadMetrics(userId, dateRange, view string, filter *models.Filter) *[]models.Metric {
	return models.LoadMetrics(userId, dateRange, view, filter, d.db)
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
	case totalSales:
		return metric.TotalSales
	case unitsSold:
		return float64(metric.Quantity)
	case amazonCosts:
		return metric.AmazonCosts
	case productCosts:
		return metric.ProductCosts
	case advertisingSpend:
		return metric.AdvertisingSpend
	case refunds:
		return metric.Refunds
	case shippingCredits:
		return metric.ShippingCredits
	case promotionalRebates:
		return metric.PromotionalRebates
	case totalCosts:
		return metric.TotalCosts
	case grossMargin:
		return metric.GrossMargin
	case netMargin:
		return metric.NetMargin
	default:
		return 0.0
	}
}
