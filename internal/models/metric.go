package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Metric struct {
	ID                         int64     `db:"id"`
	UserID                     string    `db:"user_id"`
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
	TotalCostsPercentage       float64   `db:"total_costs_percentage"`
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

var (
	metricsTotalSalesSortMap = map[int64]string{
		0: "marketplace",
		1: "sku",
		2: "description",
	}
	metricsUnitsSoldSortMap = map[int64]string{
		0: "marketplace",
		1: "sku",
		2: "description",
		3: "quantity",
	}
	metricsAmazonCostsSortMap = map[int64]string{
		0: "marketplace",
		1: "sku",
		2: "description",
		3: "amazon_costs",
	}
	metricsProductCostsSortMap = map[int64]string{
		0: "marketplace",
		1: "sku",
		2: "description",
		3: "quantity",
		4: "product_costs",
		5: "advertising_spend",
		6: "refunds",
		7: "total_costs",
	}
	metricsAdvertisingSpendSortMap = map[int64]string{
		0: "marketplace",
		1: "sku",
		2: "description",
		3: "advertising_spend",
		4: "advertising_spend_percentage",
	}
	metricsRefundsSortMap = map[int64]string{
		0: "marketplace",
		1: "sku",
		2: "description",
		3: "refunds",
		4: "refunds_percentage",
	}
	metricsShippingCreditsSortMap = map[int64]string{
		0: "marketplace",
		1: "sku",
		2: "description",
		3: "shipping_credits",
		4: "shipping_credits_tax",
	}
	metricsPromotionalRebatesSortMap = map[int64]string{
		0: "marketplace",
		1: "sku",
		2: "description",
		3: "promotional_rebates",
		4: "promotional_rebates_tax",
	}
	metricsTotalCostsSortMap = map[int64]string{
		0: "marketplace",
		1: "sku",
		2: "description",
		3: "amazon_costs",
		4: "product_costs",
		5: "product_costs_unit",
		6: "total_costs",
		7: "total_costs_percentage",
	}
	metricsGrossMarginSortMap = map[int64]string{
		0:  "marketplace",
		1:  "sku",
		2:  "description",
		3:  "product_costs",
		4:  "quantity",
		5:  "total_sales",
		6:  "amazon_costs",
		7:  "shipping_credits",
		8:  "promotional_rebates",
		9:  "gross_margin",
		10: "sales_tax_collected",
		11: "total_collected",
	}
	metricsNetMarginSortMap = map[int64]string{
		0: "marketplace",
		1: "sku",
		2: "description",
		3: "quantity",
		4: "gross_margin",
		5: "total_costs",
		6: "net_margin",
		7: "net_margin_unit",
		8: "roi",
	}
)

func LoadMetrics(userID, dateRange, view string, filter *Filter, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT * FROM metrics_%v WHERE user_id = $1 AND date_range = $2 ORDER BY $3 LIMIT $4 OFFSET $5`
	_ = db.Select(&metrics,
		fmt.Sprintf(query, view),
		userID,
		dateRange,
		fmt.Sprintf("%v %v", sortColumn(metricsViewFilter(view), filter.Sort), filter.Dir),
		filter.Length,
		filter.Start,
	)

	return &metrics
}

func LoadMetricsTotalSalesSummary(userID, dateRange string, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT date_time,
       marketplace,
       SUM(total_sales) AS total_sales FROM metrics_total_sales WHERE user_id = $1
                                                                  AND date_range = $2 GROUP BY date_time, marketplace`
	_ = db.Select(&metrics,
		query,
		userID,
		dateRange,
	)

	return &metrics
}

func LoadMetricsUnitsSoldSummary(userID, dateRange string, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT date_time,
       marketplace,
       SUM(quantity) AS quantity FROM metrics_units_sold WHERE user_id = $1
                                                           AND date_range = $2 GROUP BY date_time, marketplace`
	_ = db.Select(&metrics,
		query,
		userID,
		dateRange,
	)

	return &metrics
}

func LoadMetricsAmazonCostsSummary(userID, dateRange string, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT date_time,
       marketplace,
       SUM(amazon_costs) AS amazon_costs FROM metrics_amazon_costs WHERE user_id = $1
                                                                     AND date_range = $2 GROUP BY date_time, marketplace`
	_ = db.Select(&metrics,
		query,
		userID,
		dateRange,
	)

	return &metrics
}

func LoadMetricsProductCostsSummary(userID, dateRange string, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT date_time,
       marketplace,
       SUM(product_costs) AS product_costs FROM metrics_product_costs WHERE user_id = $1
                                                                        AND date_range = $2 GROUP BY date_time, marketplace`
	_ = db.Select(&metrics,
		query,
		userID,
		dateRange,
	)

	return &metrics
}

func LoadMetricsAdvertisingSpendSummary(userID, dateRange string, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT date_time,
       marketplace,
       SUM(advertising_spend) AS advertising_spend FROM metrics_advertising_spend WHERE user_id = $1
                                                                                    AND date_range = $2 GROUP BY date_time, marketplace`
	_ = db.Select(&metrics,
		query,
		userID,
		dateRange,
	)

	return &metrics
}

func LoadMetricsRefundsSummary(userID, dateRange string, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT date_time,
       marketplace,
       SUM(refunds) AS refunds FROM metrics_refunds WHERE user_id = $1
                                                      AND date_range = $2 GROUP BY date_time, marketplace`
	_ = db.Select(&metrics,
		query,
		userID,
		dateRange,
	)

	return &metrics
}

func LoadMetricsShippingCreditsSummary(userID, dateRange string, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT date_time,
       marketplace,
       SUM(shipping_credits) AS shipping_credits FROM metrics_shipping_credits WHERE user_id = $1
                                                                                 AND date_range = $2 GROUP BY date_time, marketplace`
	_ = db.Select(&metrics,
		query,
		userID,
		dateRange,
	)

	return &metrics
}

func LoadMetricsPromotionalRebatesSummary(userID, dateRange string, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT date_time,
       marketplace,
       SUM(promotional_rebates) AS promotional_rebates FROM metrics_promotional_rebates WHERE user_id = $1
                                                                                          AND date_range = $2 GROUP BY date_time, marketplace`
	_ = db.Select(&metrics,
		query,
		userID,
		dateRange,
	)

	return &metrics
}

func LoadMetricsTotalCostsSummary(userID, dateRange string, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT date_time,
       marketplace,
       SUM(total_costs) AS total_costs FROM metrics_total_costs WHERE user_id = $1
                                                                  AND date_range = $2 GROUP BY date_time, marketplace`
	_ = db.Select(&metrics,
		query,
		userID,
		dateRange,
	)

	return &metrics
}

func LoadMetricsGrossMarginSummary(userID, dateRange string, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT date_time,
       marketplace,
       SUM(gross_margin) AS gross_margin FROM metrics_gross_margin WHERE user_id = $1
                                                                     AND date_range = $2 GROUP BY date_time, marketplace`
	_ = db.Select(&metrics,
		query,
		userID,
		dateRange,
	)

	return &metrics
}

func LoadMetricsNetMarginSummary(userID, dateRange string, db *sqlx.DB) *[]Metric {
	var metrics []Metric

	query := `SELECT date_time,
       marketplace,
       SUM(net_margin) AS net_margin FROM metrics_net_margin WHERE user_id = $1
                                                               AND date_range = $2 GROUP BY date_time, marketplace`
	_ = db.Select(&metrics,
		query,
		userID,
		dateRange,
	)

	return &metrics
}

func TotalMetrics(userID, dateRange, view string, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(id) FROM metrics_%v WHERE user_id = $1 AND date_range = $2`
	_ = db.QueryRow(fmt.Sprintf(query, view), userID, dateRange).Scan(&count)

	return count
}

func metricsViewFilter(view string) map[int64]string {
	switch view {
	case TotalSales:
		return metricsTotalSalesSortMap
	case UnitsSold:
		return metricsUnitsSoldSortMap
	case AmazonCosts:
		return metricsAmazonCostsSortMap
	case ProductCosts:
		return metricsProductCostsSortMap
	case AdvertisingSpend:
		return metricsAdvertisingSpendSortMap
	case Refunds:
		return metricsRefundsSortMap
	case ShippingCredits:
		return metricsShippingCreditsSortMap
	case PromotionalRebates:
		return metricsPromotionalRebatesSortMap
	case TotalCosts:
		return metricsTotalCostsSortMap
	case GrossMargin:
		return metricsGrossMarginSortMap
	case NetMargin:
		return metricsNetMarginSortMap
	default:
		return map[int64]string{}
	}
}
