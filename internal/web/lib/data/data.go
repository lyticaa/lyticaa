package data

import (
	"gitlab.com/getlytica/lytica-app/internal/web/lib/chart"

	"github.com/jmoiron/sqlx"
)

const (
	unitsSold          = "units_sold"
	amazonCosts        = "amazon_costs"
	productCosts       = "product_costs"
	advertisingSpend   = "advertising_spend"
	refunds            = "refunds"
	shippingCredits    = "shipping_credits"
	promotionalRebates = "promotional_rebates"
	totalCosts         = "total_costs"
	grossMargin        = "gross_margin"
	netMargin          = "net_margin"
)

type Data struct {
	chart *chart.Chart
	db    *sqlx.DB
}

func NewData(db *sqlx.DB) *Data {
	return &Data{
		chart: chart.NewChart(),
		db:    db,
	}
}
