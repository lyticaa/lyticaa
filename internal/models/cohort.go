package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Cohort struct {
	ID                 int64     `db:"id"`
	UserID             string    `db:"user_id"`
	DateTime           time.Time `db:"date_time"`
	Marketplace        string    `db:"marketplace"`
	SKU                string    `db:"sku"`
	Description        string    `db:"description"`
	Quantity           int64     `db:"quantity"`
	TotalSales         float64   `db:"total_sales"`
	AmazonCosts        float64   `db:"amazon_costs"`
	ProductCosts       float64   `db:"product_costs"`
	AdvertisingSpend   float64   `db:"advertising_spend"`
	Refunds            float64   `db:"refunds"`
	ShippingCredits    float64   `db:"shipping_credits"`
	PromotionalRebates float64   `db:"promotional_rebates"`
	TotalCosts         float64   `db:"total_costs"`
	NetMargin          float64   `db:"net_margin"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

func LoadCohorts(userID, dateRange, view string, filter *Filter, db *sqlx.DB) *[]Cohort {
	var cohorts []Cohort

	query := `SELECT * FROM cohorts WHERE user_id = $1 AND date_range = $2 AND view = $3 LIMIT $4 OFFSET $5`
	_ = db.Select(&cohorts,
		query,
		userID,
		dateRange,
		view,
		filter.Length,
		filter.Start,
	)

	return &cohorts
}

func LoadCohortsSummary(userID, dateRange, view string, db *sqlx.DB) *[]Cohort {
	var cohorts []Cohort

	query := `SELECT date_time, 
       marketplace, 
       SUM(quantity) AS quantity,
       SUM(total_sales) AS total_sales,
       SUM(amazon_costs) AS amazon_costs,
       SUM(product_costs) AS product_costs,
       SUM(advertising_spend) AS advertising_spend,
       SUM(refunds) AS refunds,
       SUM(shipping_credits) AS shipping_credits,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(total_costs) AS total_costs,
       SUM(net_margin) FROM cohorts WHERE user_id = $1 
                                      AND date_range = $2 AND view = $3 GROUP BY date_time, marketplace`
	_ = db.Select(&cohorts,
		query,
		userID,
		dateRange,
		view,
	)

	return &cohorts
}

func TotalCohorts(userID, dateRange, view string, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(id) FROM cohorts WHERE user_id = $1 AND date_range = $2 AND view = $3`
	_ = db.QueryRow(query, userID, dateRange, view).Scan(&count)

	return count
}
