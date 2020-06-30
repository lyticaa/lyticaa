package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Dashboard struct {
	Id                 int64     `db:"id"`
	UserId             string    `db:"user_id"`
	DateTime           time.Time `db:"date_time"`
	Marketplace        string    `db:"marketplace"`
	TotalSales         float64   `db:"total_sales"`
	UnitsSold          int64     `db:"units_sold"`
	AmazonCosts        float64   `db:"amazon_costs"`
	ProductCosts       float64   `db:"product_costs"`
	AdvertisingSpend   float64   `db:"advertising_spend"`
	Refunds            float64   `db:"refunds"`
	ShippingCredits    float64   `db:"shipping_credits"`
	PromotionalRebates float64   `db:"promotional_rebates"`
	TotalCosts         float64   `db:"total_costs"`
	GrossMargin        float64   `db:"gross_margin"`
	NetMargin          float64   `db:"net_margin"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

func LoadDashboard(userId, dateRange string, db *sqlx.DB) *[]Dashboard {
	var dashboard []Dashboard

	query := `SELECT * FROM dashboard_%v WHERE user_id = $1`
	_ = db.Select(&dashboard, fmt.Sprintf(query, dateRange), userId)

	return &dashboard
}

func LoadDashboardTotalSales(userId, dateRange string, db *sqlx.DB) *[]Dashboard {
	var dashboard []Dashboard

	query := `SELECT date_time, marketplace, SUM(total_sales) FROM dashboard_%v WHERE user_id = $1 GROUP BY date_time, marketplace_id`
	_ = db.Select(&dashboard, fmt.Sprintf(query, dateRange), userId)

	return &dashboard
}

func LoadDashboardTotals(userId, dateRange string, db *sqlx.DB) *Dashboard {
	var dashboard Dashboard

	query := `SELECT,
       SUM(units_sold),
       SUM(amazon_costs),
       SUM(product_costs),
       SUM(advertising_spend),
       SUM(refunds),
       SUM(shipping_credits),
       SUM(promotional_rebates),
       SUM(total_costs),
       SUM(gross_margin),
       SUM(net_margin)
FROM dashboard_%v
WHERE user_id = $1`
	_ = db.QueryRow(fmt.Sprintf(query, dateRange), userId).Scan(
		&dashboard.UnitsSold,
		&dashboard.AmazonCosts,
		&dashboard.ProductCosts,
		&dashboard.AdvertisingSpend,
		&dashboard.Refunds,
		&dashboard.ShippingCredits,
		&dashboard.PromotionalRebates,
		&dashboard.TotalCosts,
		&dashboard.GrossMargin,
		&dashboard.NetMargin,
	)

	return &dashboard
}
