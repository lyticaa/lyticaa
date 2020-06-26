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

func (m *Metrics) summaryData(dateRange, view string, current *[]models.Transaction) *[]types.Summary {
	var summary []types.Summary

	switch view {
	case helpers.TotalSalesView:
		summary = m.amazon.TotalSales(current)
	case helpers.UnitsSoldView:
		summary = m.amazon.UnitsSold(current)
	case helpers.AmazonCostsView:
		summary = m.amazon.AmazonCosts(current)
	case helpers.ProductCostsView:
		summary = m.amazon.ProductCosts(current)
	case helpers.AdvertisingSpendView:
		summary = m.amazon.AdvertisingSpend(current)
	case helpers.RefundsView:
		summary = m.amazon.Refunds(current)
	case helpers.ShippingCreditsView:
		summary = m.amazon.ShippingCredits(current)
	case helpers.PromotionalRebatesView:
		summary = m.amazon.PromotionalRebates(current)
	case helpers.TotalCostsView:
		summary = m.amazon.TotalCosts(current)
	case helpers.GrossMarginView:
		summary = m.amazon.GrossMargin(current)
	case helpers.NetMarginView:
		summary = m.amazon.NetMargin(current)
	}

	return &summary
}

func (m *Metrics) chartData(dateRange string, summary *[]types.Summary, byDate *types.Metrics) {
	byDate.Chart = m.chart.Line(summary, dateRange)
}
