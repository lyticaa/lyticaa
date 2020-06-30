package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Cohort struct {
	ID                 int64     `db:"id"`
	UserId             string    `db:"user_id"`
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

func LoadCohorts(userId, dateRange, view string, db *sqlx.DB) *[]Cohort {
	var cohorts []Cohort

	query := `SELECT * FROM cohorts_%v_%v WHERE user_id = $1`
	_ = db.Select(&cohorts, fmt.Sprintf(query, view, dateRange), userId)

	return &cohorts
}
