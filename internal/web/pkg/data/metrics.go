package data

import (
	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/web/types"
)

func (d *Data) MetricsTotalSales(userID, dateRange string, metric *types.TotalSales, filter *models.Filter) {
}

func (d *Data) MetricsUnitsSold(userID, dateRange string, metric *types.UnitsSold, filter *models.Filter) {
}

func (d *Data) MetricsAmazonCosts(userID, dateRange string, metric *types.AmazonCosts, filter *models.Filter) {
}

func (d *Data) MetricsProductCosts(userID, dateRange string, metric *types.ProductCosts, filter *models.Filter) {
}

func (d *Data) MetricsAdvertisingSpend(userID, dateRange string, metric *types.AdvertisingSpend, filter *models.Filter) {
}

func (d *Data) MetricsRefunds(userID, dateRange string, metric *types.Refunds, filter *models.Filter) {
}

func (d *Data) MetricsShippingCredits(userID, dateRange string, metric *types.ShippingCredits, filter *models.Filter) {
}

func (d *Data) MetricsPromotionalRebates(userID, dateRange string, metric *types.PromotionalRebates, filter *models.Filter) {
}

func (d *Data) MetricsTotalCosts(userID, dateRange string, metric *types.TotalCosts, filter *models.Filter) {
}

func (d *Data) MetricsGrossMargin(userID, dateRange string, metric *types.GrossMargin, filter *models.Filter) {
}

func (d *Data) MetricsNetMargin(userID, dateRange string, metric *types.NetMargin, filter *models.Filter) {
}
