package metrics

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/helpers"
	"gitlab.com/getlytica/lytica-app/internal/web/app/lib/amazon"
	"gitlab.com/getlytica/lytica-app/internal/web/app/lib/chart"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"gopkg.in/boj/redistore.v1"
)

type Metrics struct {
	amazon       *amazon.Amazon
	chart        *chart.Chart
	db           *sqlx.DB
	sessionStore *redistore.RediStore
	logger       zerolog.Logger
}

func NewMetrics(db *sqlx.DB, sessionStore *redistore.RediStore, log zerolog.Logger) *Metrics {
	return &Metrics{
		amazon:       amazon.NewAmazon(db),
		chart:        chart.NewChart(),
		sessionStore: sessionStore,
		logger:       log,
		db:           db,
	}
}

func (m *Metrics) summaryData(dateRange, view string, txns *[]models.Transaction, sps *[]models.SponsoredProduct) *[]types.Summary {
	var summary []types.Summary

	switch view {
	case helpers.TotalSalesView:
		summary = m.amazon.TotalSales(txns)
	case helpers.UnitsSoldView:
		summary = m.amazon.UnitsSold(txns)
	case helpers.AmazonCostsView:
		summary = m.amazon.AmazonCosts(txns)
	case helpers.ProductCostsView:
		summary = m.amazon.ProductCosts(txns)
	case helpers.AdvertisingSpendView:
		summary = m.amazon.AdvertisingSpend(txns, sps)
	case helpers.RefundsView:
		summary = m.amazon.Refunds(txns)
	case helpers.ShippingCreditsView:
		summary = m.amazon.ShippingCredits(txns)
	case helpers.PromotionalRebatesView:
		summary = m.amazon.PromotionalRebates(txns)
	case helpers.TotalCostsView:
		summary = m.amazon.TotalCosts(txns)
	case helpers.GrossMarginView:
		summary = m.amazon.GrossMargin(txns)
	case helpers.NetMarginView:
		summary = m.amazon.NetMargin(txns)
	}

	return &summary
}

func (m *Metrics) chartData(dateRange string, summary *[]types.Summary, byDate *types.Metrics) {
	byDate.Chart = m.chart.Line(summary, dateRange)
}
