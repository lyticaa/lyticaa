package types

import (
	"time"
)

type Total struct {
	Value int64   `json:"value"`
	Diff  float64 `json:"diff"`
}

type Card struct {
	Total `json:"total"`
	Chart Chart `json:"chart"`
}

type Dashboard struct {
	TotalSales         Chart `json:"totalSales,omitempty"`
	UnitsSold          Card  `json:"unitsSold,omitempty"`
	AmazonCosts        Card  `json:"amazonCosts,omitempty"`
	AdvertisingSpend   Card  `json:"advertisingSpend,omitempty"`
	Refunds            Card  `json:"refunds,omitempty"`
	ShippingCredits    Card  `json:"shippingCredits,omitempty"`
	PromotionalRebates Card  `json:"promotionalRebates,omitempty"`
	TotalCosts         Card  `json:"totalCosts,omitempty"`
	NetMargin          Card  `json:"netMargin,omitempty"`
}

type TotalSalesTable struct {
	Date        time.Time `json:"date"`
	SKU         string    `json:"sku"`
	ASIN        string    `json:"asin"`
	ProductName string    `json:"productName"`
	Sales       float64   `json:"sales"`
}

type TotalSales struct {
	Chart           Chart             `json:"chart"`
	Card            Card              `json:"card,omitempty"`
	Data            []TotalSalesTable `json:"data"`
	Draw            int64             `json:"draw"`
	RecordsTotal    int64             `json:"recordsTotal"`
	RecordsFiltered int64             `json:"recordsFiltered"`
}

type UnitsSoldTable struct {
	Date         time.Time `json:"date"`
	SKU          string    `json:"sku"`
	QuantitySold int64     `json:"quantitySold"`
}

type UnitsSold struct {
	Chart           Chart            `json:"chart"`
	Card            Card             `json:"card,omitempty"`
	Data            []UnitsSoldTable `json:"data"`
	Draw            int64            `json:"draw"`
	RecordsTotal    int64            `json:"recordsTotal"`
	RecordsFiltered int64            `json:"recordsFiltered"`
}

type AmazonCostsTable struct {
	Date        time.Time `json:"date"`
	SKU         string    `json:"sku"`
	Type        string    `json:"type"`
	AmazonCosts float64   `json:"amazonCosts"`
}

type AmazonCosts struct {
	Chart           Chart              `json:"chart"`
	Card            Card               `json:"card,omitempty"`
	Data            []AmazonCostsTable `json:"data"`
	Draw            int64              `json:"draw"`
	RecordsTotal    int64              `json:"recordsTotal"`
	RecordsFiltered int64              `json:"recordsFiltered"`
}

type AdvertisingSpendTable struct {
	Date              time.Time `json:"date"`
	SKU               string    `json:"sku"`
	AdvertisingSpend  float64   `json:"advertisingSpend"`
	PercentageOfSales float64   `json:"percentageOfSales"`
}

type AdvertisingSpend struct {
	Chart           Chart                   `json:"chart"`
	Card            Card                    `json:"card,omitempty"`
	Data            []AdvertisingSpendTable `json:"data"`
	Draw            int64                   `json:"draw"`
	RecordsTotal    int64                   `json:"recordsTotal"`
	RecordsFiltered int64                   `json:"recordsFiltered"`
}

type RefundsTable struct {
	Date              time.Time `json:"date"`
	SKU               string    `json:"sku"`
	Refunds           float64   `json:"refunds"`
	PercentageOfSales float64   `json:"percentageOfSales"`
}

type Refunds struct {
	Chart           Chart          `json:"chart"`
	Card            Card           `json:"card,omitempty"`
	Data            []RefundsTable `json:"data"`
	Draw            int64          `json:"draw"`
	RecordsTotal    int64          `json:"recordsTotal"`
	RecordsFiltered int64          `json:"recordsFiltered"`
}

type ShippingCreditsTable struct {
	Date            time.Time `json:"date"`
	SKU             string    `json:"sku"`
	ShippingCredits float64   `json:"shippingCredits"`
}

type ShippingCredits struct {
	Chart           Chart                  `json:"chart"`
	Card            Card                   `json:"card,omitempty"`
	Data            []ShippingCreditsTable `json:"data"`
	Draw            int64                  `json:"draw"`
	RecordsTotal    int64                  `json:"recordsTotal"`
	RecordsFiltered int64                  `json:"recordsFiltered"`
}

type PromotionalRebatesTable struct {
	Date               time.Time `json:"date"`
	SKU                string    `json:"sku"`
	CostOfCoupons      float64   `json:"costOfCoupons"`
	Quantity           int64     `json:"quantity"`
	PromotionalRebates float64   `json:"promotionalRebates"`
}

type PromotionalRebates struct {
	Chart           Chart                     `json:"chart"`
	Card            Card                      `json:"card,omitempty"`
	Data            []PromotionalRebatesTable `json:"data"`
	Draw            int64                     `json:"draw"`
	RecordsTotal    int64                     `json:"recordsTotal"`
	RecordsFiltered int64                     `json:"recordsFiltered"`
}

type TotalCostsTable struct {
	Date               time.Time `json:"date"`
	SKU                string    `json:"sku"`
	AmazonCosts        float64   `json:"amazonCosts"`
	ProductCosts       float64   `json:"productCosts"`
	ProductCostPerUnit float64   `json:"productCostPerUnit"`
	TotalCosts         float64   `json:"totalCosts"`
	Percentage         float64   `json:"percentage"`
	PercentageOfSales  float64   `json:"percentageOfSales"`
}

type TotalCosts struct {
	Chart           Chart             `json:"chart"`
	Card            Card              `json:"card,omitempty"`
	Data            []TotalCostsTable `json:"data"`
	Draw            int64             `json:"draw"`
	RecordsTotal    int64             `json:"recordsTotal"`
	RecordsFiltered int64             `json:"recordsFiltered"`
}

type NetMarginTable struct {
	Date             time.Time `json:"date"`
	SKU              string    `json:"sku"`
	NetMargin        float64   `json:"netMargin"`
	Percentage       float64   `json:"percentage"`
	NetMarginPerUnit float64   `json:"netMarginPerUnit"`
	ROI              float64   `json:"roi"`
}

type NetMargin struct {
	Chart           Chart            `json:"chart"`
	Card            Card             `json:"card,omitempty"`
	Data            []NetMarginTable `json:"data"`
	Draw            int64            `json:"draw"`
	RecordsTotal    int64            `json:"recordsTotal"`
	RecordsFiltered int64            `json:"recordsFiltered"`
}
