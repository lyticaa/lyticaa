package types

import (
	"time"
)

type Dashboard struct {
	TotalSales         Card `json:"totalSales,omitempty"`
	UnitsSold          Card `json:"unitsSold,omitempty"`
	AmazonCosts        Card `json:"amazonCosts,omitempty"`
	AdvertisingSpend   Card `json:"advertisingSpend,omitempty"`
	Refunds            Card `json:"refunds,omitempty"`
	ShippingCredits    Card `json:"shippingCredits,omitempty"`
	PromotionalRebates Card `json:"promotionalRebates,omitempty"`
	TotalCosts         Card `json:"totatCosts,omitempty"`
	NetMargin          Card `json:"netMargin,omitempty"`
}

type TotalSalesTable struct {
	Date        time.Time `json:"date"`
	SKU         string    `json:"sku"`
	ASIN        string    `json:"asin"`
	ProductName string    `json:"productName"`
	Sales       float64   `json:"sales"`
}

type TotalSales struct {
	Chart        Flot              `json:"chart"`
	Data         []TotalSalesTable `json:"data"`
	Draw         int64             `json:"draw"`
	RecordsTotal int64             `json:"recordsTotal"`
}

type UnitsSoldTable struct {
	Date              time.Time `json:"date"`
	SKU               string    `json:"sku"`
	ASIN              string    `json:"asin"`
	ProductName       string    `json:"productName"`
	GrossQuantitySold int64     `json:"grossQuantitySold"`
	NetQuantitySold   int64     `json:"netQuantitySold"`
}

type UnitsSold struct {
	Chart        Flot             `json:"chart"`
	Data         []UnitsSoldTable `json:"data"`
	Draw         int64            `json:"draw"`
	RecordsTotal int64            `json:"recordsTotal"`
}

type AmazonCostsTable struct {
	Date        time.Time `json:"date"`
	SKU         string    `json:"sku"`
	ASIN        string    `json:"asin"`
	ProductName string    `json:"productName"`
	Type        string    `json:"type"`
	AmazonCosts float64   `json:"amazonCosts"`
}

type AmazonCosts struct {
	Chart        Flot               `json:"chart"`
	Data         []AmazonCostsTable `json:"data"`
	Draw         int64              `json:"draw"`
	RecordsTotal int64              `json:"recordsTotal"`
}

type AdvertisingSpendTable struct {
	Date              time.Time `json:"date"`
	SKU               string    `json:"sku"`
	ASIN              string    `json:"asin"`
	ProductName       string    `json:"productName"`
	AdvertisingSpend  float64   `json:"advertisingSpend"`
	PercentageOfSales float64   `json:"percentageOfSales"`
}

type AdvertisingSpend struct {
	Chart        Flot                    `json:"chart"`
	Data         []AdvertisingSpendTable `json:"data"`
	Draw         int64                   `json:"draw"`
	RecordsTotal int64                   `json:"recordsTotal"`
}

type RefundsTable struct {
	Date              time.Time `json:"date"`
	SKU               string    `json:"sku"`
	ASIN              string    `json:"asin"`
	ProductName       string    `json:"productName"`
	Refunds           float64   `json:"refunds"`
	PercentageOfSales float64   `json:"percentageOfSales"`
}

type Refunds struct {
	Chart        Flot           `json:"chart"`
	Data         []RefundsTable `json:"data"`
	Draw         int64          `json:"draw"`
	RecordsTotal int64          `json:"recordsTotal"`
}

type ShippingCreditsTable struct {
	Date            time.Time `json:"date"`
	SKU             string    `json:"sku"`
	ASIN            string    `json:"asin"`
	ProductName     string    `json:"productName"`
	ShippingCredits float64   `json:"shippingCredits"`
}

type ShippingCredits struct {
	Chart        Flot                   `json:"chart"`
	Data         []ShippingCreditsTable `json:"data"`
	Draw         int64                  `json:"draw"`
	RecordsTotal int64                  `json:"recordsTotal"`
}

type PromotionalRebatesTable struct {
	Date               time.Time `json:"date"`
	SKU                string    `json:"sku"`
	ASIN               string    `json:"asin"`
	ProductName        string    `json:"productName"`
	CostOfCoupons      float64   `json:"costOfCoupons"`
	Quantity           int64     `json:"quantity"`
	PromotionalRebates float64   `json:"promotionalRebates"`
}

type PromotionalRebates struct {
	Chart        Flot                      `json:"chart"`
	Data         []PromotionalRebatesTable `json:"data"`
	Draw         int64                     `json:"draw"`
	RecordsTotal int64                     `json:"recordsTotal"`
}

type TotalCostsTable struct {
	Date               time.Time `json:"date"`
	SKU                string    `json:"sku"`
	ASIN               string    `json:"asin"`
	ProductName        string    `json:"productName"`
	AmazonCosts        float64   `json:"amazonCosts"`
	ProductCosts       float64   `json:"productCosts"`
	ProductCostPerUnit float64   `json:"productCostPerUnit"`
	TotalCosts         float64   `json:"totalCosts"`
	Percentage         float64   `json:"percentage"`
	PercentageOfSales  float64   `json:"percentageOfSales"`
}

type TotalCosts struct {
	Chart        Flot              `json:"chart"`
	Data         []TotalCostsTable `json:"data"`
	Draw         int64             `json:"draw"`
	RecordsTotal int64             `json:"recordsTotal"`
}

type NetMarginTable struct {
	Date             time.Time `json:"date"`
	SKU              string    `json:"sku"`
	ASIN             string    `json:"asin"`
	ProductName      string    `json:"productName"`
	NetMargin        float64   `json:"netMargin"`
	Percentage       float64   `json:"percentage"`
	NetMarginPerUnit float64   `json:"netMarginPerUnit"`
	ROI              float64   `json:"roi"`
}

type NetMargin struct {
	Chart        Flot             `json:"chart"`
	Data         []NetMarginTable `json:"data"`
	Draw         int64            `json:"draw"`
	RecordsTotal int64            `json:"recordsTotal"`
}
