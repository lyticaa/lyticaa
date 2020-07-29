package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Dashboard struct {
	Id                 int64     `db:"id"`
	DateRange          string    `db:"date_range"`
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

	query := `SELECT * FROM dashboard WHERE user_id = $1 AND date_range = $2`
	_ = db.Select(&dashboard, query, userId, dateRange)

	return &dashboard
}

func LoadDashboardByMarketplace(userId, dateRange, marketplace string, dateTime time.Time, db *sqlx.DB) *Dashboard {
	var dashboard Dashboard

	query := `SELECT * FROM dashboard WHERE user_id = $1 AND date_range = $2 AND marketplace = $3 AND date_time = $4`
	_ = db.QueryRow(query, userId, dateRange, marketplace, dateTime).Scan(
		&dashboard.Id,
		&dashboard.DateRange,
		&dashboard.UserId,
		&dashboard.DateTime,
		&dashboard.Marketplace,
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

func LoadDashboardTotalSales(userId, dateRange string, db *sqlx.DB) *[]Dashboard {
	var dashboard []Dashboard

	query := `SELECT
       date_time,
       marketplace, 
       SUM(total_sales) FROM dashboard WHERE user_id = $1
                                         AND date_range = $2 GROUP BY date_time, marketplace`
	_ = db.Select(&dashboard, query, userId, dateRange)

	return &dashboard
}

func LoadDashboardTotals(userId, dateRange string, db *sqlx.DB) *Dashboard {
	var dashboard Dashboard

	query := `SELECT
       SUM(units_sold),
       SUM(amazon_costs),
       SUM(product_costs),
       SUM(advertising_spend),
       SUM(refunds),
       SUM(shipping_credits),
       SUM(promotional_rebates),
       SUM(total_costs),
       SUM(gross_margin),
       SUM(net_margin) FROM dashboard WHERE user_id = $1 
                                        AND date_range = $2`
	_ = db.QueryRow(query, userId, dateRange).Scan(
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

func (d *Dashboard) Save(db *sqlx.DB) error {
	query := `INSERT INTO dashboard (
                       date_range,
                       user_id,
                       date_time,
                       marketplace,
                       units_sold,
                       amazon_costs,
                       product_costs,
                       advertising_spend,
                       refunds,
                       shipping_credits,
                       promotional_rebates,
                       total_costs,
                       gross_margin,
                       net_margin)
                       VALUES (
                               :date_range,
                               :user_id,
                               :date_time,
                               :marketplace,
                               :units_sold,
                               :amazon_costs,
                               :product_costs,
                               :advertising_spend,
                               :refunds,
                               :shipping_credits,
                               :promotional_rebates,
                               :total_costs,
                               :gross_margin,
                               :net_margin)
                               ON CONFLICT (date_range, user_id, date_time, marketplace)
                                   DO UPDATE SET units_sold = :units_sold,
                                                 amazon_costs = :amazon_costs,
                                                 product_costs = :product_costs,
                                                 advertising_spend = :advertising_spend,
                                                 refunds = :refunds,
                                                 shipping_credits = :shipping_credits,
                                                 promotional_rebates = :promotional_rebates,
                                                 total_costs = :total_costs,
                                                 gross_margin = :gross_margin,
                                                 net_margin = :net_margin,
                                                 updated_at = NOW()`
	_, err := db.NamedExec(query, map[string]interface{}{
		"date_range":          d.DateRange,
		"user_id":             d.UserId,
		"date_time":           d.DateTime,
		"marketplace":         d.Marketplace,
		"units_sold":          d.UnitsSold,
		"amazon_costs":        d.AmazonCosts,
		"product_costs":       d.ProductCosts,
		"advertising_spend":   d.AdvertisingSpend,
		"refunds":             d.Refunds,
		"shipping_credits":    d.ShippingCredits,
		"promotional_rebates": d.PromotionalRebates,
		"total_costs":         d.TotalCosts,
		"gross_margin":        d.GrossMargin,
		"net_margin":          d.NetMargin,
	})

	if err != nil {
		return err
	}

	return nil
}
