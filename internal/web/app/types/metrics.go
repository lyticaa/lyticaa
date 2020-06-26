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

type Metrics struct {
	Chart           Chart `json:"chart"`
	Card            Card  `json:"card,omitempty"`
	Draw            int64 `json:"draw"`
	RecordsTotal    int64 `json:"recordsTotal"`
	RecordsFiltered int64 `json:"recordsFiltered"`
}

type Dashboard struct {
	TotalSales         Chart `json:"totalSales,omitempty"`
	UnitsSold          Card  `json:"unitsSold,omitempty"`
	AmazonCosts        Card  `json:"amazonCosts,omitempty"`
	ProductCosts       Card  `json:"productCosts,omitempty"`
	AdvertisingSpend   Card  `json:"advertisingSpend,omitempty"`
	Refunds            Card  `json:"refunds,omitempty"`
	ShippingCredits    Card  `json:"shippingCredits,omitempty"`
	PromotionalRebates Card  `json:"promotionalRebates,omitempty"`
	TotalCosts         Card  `json:"totalCosts,omitempty"`
	GrossMargin        Card  `json:"grossMargin,omitempty"`
	NetMargin          Card  `json:"netMargin,omitempty"`
}

type TotalSalesTable struct {
	SKU         string  `json:"sku"`
	Description string  `json:"description"`
	Marketplace string  `json:"marketplace"`
	TotalSales  float64 `json:"totalSales"`
}

type TotalSales struct {
	Metrics
	Data []TotalSalesTable `json:"data"`
}

type UnitsSoldTable struct {
	SKU            string `json:"sku"`
	Description    string `json:"description"`
	Marketplace    string `json:"marketplace"`
	TotalUnitsSold int64  `json:"totalUnitsSold"`
}

type UnitsSold struct {
	Metrics
	Data []UnitsSoldTable `json:"data"`
}

type AmazonCostsTable struct {
	SKU              string  `json:"sku"`
	Description      string  `json:"type"`
	Marketplace      string  `json:"marketplace"`
	TotalAmazonCosts float64 `json:"totalAmazonCosts"`
}

type AmazonCosts struct {
	Metrics
	Data []AmazonCostsTable `json:"data"`
}

type ProductCostsTable struct {
	SKU               string  `json:"sku"`
	Description       string  `json:"description"`
	Marketplace       string  `json:"marketplace"`
	QuantitySold      int64   `json:"quantitySold"`
	ProductCosts      float64 `json:"productCosts"`
	AdvertisingCosts  float64 `json:"advertisingCosts"`
	Refunds           float64 `json:"refunds"`
	TotalProductCosts float64 `json:"totalProductCosts"`
}

type ProductCosts struct {
	Metrics
	Data []ProductCostsTable `json:"data"`
}

type AdvertisingSpendTable struct {
	Date              time.Time `json:"date"`
	SKU               string    `json:"sku"`
	AdvertisingSpend  float64   `json:"advertisingSpend"`
	PercentageOfSales float64   `json:"percentageOfSales"`
}

type AdvertisingSpend struct {
	Metrics
	Data []AdvertisingSpendTable `json:"data"`
}

type RefundsTable struct {
	Date              time.Time `json:"date"`
	SKU               string    `json:"sku"`
	Refunds           float64   `json:"refunds"`
	PercentageOfSales float64   `json:"percentageOfSales"`
}

type Refunds struct {
	Metrics
	Data []RefundsTable `json:"data"`
}

type ShippingCreditsTable struct {
	Date            time.Time `json:"date"`
	SKU             string    `json:"sku"`
	ShippingCredits float64   `json:"shippingCredits"`
}

type ShippingCredits struct {
	Metrics
	Data []ShippingCreditsTable `json:"data"`
}

type PromotionalRebatesTable struct {
	Date               time.Time `json:"date"`
	SKU                string    `json:"sku"`
	CostOfCoupons      float64   `json:"costOfCoupons"`
	Quantity           int64     `json:"quantity"`
	PromotionalRebates float64   `json:"promotionalRebates"`
}

type PromotionalRebates struct {
	Metrics
	Data []PromotionalRebatesTable `json:"data"`
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
	Metrics
	Data []TotalCostsTable `json:"data"`
}

type GrossMarginTable struct {
	SKU                  string  `json:"sku"`
	QuantitySold         int64   `json:"quantitySold"`
	QuantitySoldCoupons  int64   `json:"quantitySoldCoupons"`
	SalePrice            float64 `json:"salePrice"`
	TotalRevenue         float64 `json:"totalRevenue"`
	AmazonCosts          float64 `json:"amazonCosts"`
	ShippingCredits      float64 `json:"shippingCredits"`
	PromotionalRebates   float64 `json:"promotionalRebates"`
	GrossMargin          float64 `json:"grossMargin"`
	SalesTaxCollected    float64 `json:"salesTaxCollected"`
	TotalAmountCollected float64 `json:"totalAmountCollected"`
}

type GrossMargin struct {
	Metrics
	Data []GrossMarginTable `json:"data"`
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
	Metrics
	Data []NetMarginTable `json:"data"`
}
