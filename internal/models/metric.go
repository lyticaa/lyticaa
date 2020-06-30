package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Metric struct {
	ID                         int64     `db:"id"`
	UserId                     string    `db:"user_id"`
	DateTime                   time.Time `db:"date_time"`
	Marketplace                string    `db:"marketplace"`
	SKU                        string    `db:"sku"`
	Description                string    `db:"description"`
	TotalSales                 float64   `db:"total_sales"`
	Quantity                   int64     `db:"quantity"`
	AmazonCosts                float64   `db:"amazon_costs"`
	CostOfGoods                float64   `db:"cost_of_goods"`
	ProductCosts               float64   `db:"product_costs"`
	ProductCostsUnit           float64   `db:"product_costs_unit"`
	AdvertisingSpend           float64   `db:"advertising_spend"`
	AdvertisingSpendPercentage float64   `db:"advertising_spend_percentage"`
	Refunds                    float64   `db:"refunds"`
	RefundsPercentage          float64   `db:"refunds_percentage"`
	TotalCosts                 float64   `db:"total_costs"`
	ShippingCredits            float64   `db:"shipping_credits"`
	ShippingCreditsTax         float64   `db:"shipping_credits_tax"`
	PromotionalRebates         float64   `db:"promotional_rebates"`
	PromotionalRebatesTax      float64   `db:"promotional_rebates_tax"`
	SalesTaxCollected          float64   `db:"sales_tax_collected"`
	TotalCollected             float64   `db:"total_collected"`
	GrossMargin                float64   `db:"gross_margin"`
	NetMargin                  float64   `db:"net_margin"`
	NetMarginUnit              float64   `db:"net_margin_unit"`
	ROI                        float64   `db:"roi"`
	CreatedAt                  time.Time `db:"created_at"`
	UpdatedAt                  time.Time `db:"updated_at"`
}

func LoadMetrics(userId, dateRange, view string, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT * FROM metrics_%v_%v WHERE user_id = $1`
	_ = db.Select(&metrics, fmt.Sprintf(query, view, dateRange), userId)

	return &metrics
}
