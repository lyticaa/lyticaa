package types

import "time"

type Parsed struct {
	DateRange                  string    `json:"dateRange,omitempty"`
	UserID                     string    `json:"userID,omitempty"`
	DateTime                   time.Time `json:"dateTime,omitempty"`
	Marketplace                string    `json:"marketplace,omitempty"`
	SKU                        string    `json:"sku,omitempty"`
	Description                string    `json:"description,omitempty"`
	TotalSales                 float64   `json:"totalSales,omitempty"`
	UnitsSold                  int64     `json:"unitsSold,omitempty"`
	AmazonCosts                float64   `json:"amazonCosts,omitempty"`
	CostOfGoods                float64   `json:"costOfGoods,omitempty"`
	ProductCosts               float64   `json:"productCosts,omitempty"`
	ProductCostsUnit           float64   `json:"productCostsUnit,omitempty"`
	AdvertisingSpend           float64   `json:"advertisingSpend,omitempty"`
	AdvertisingSpendPercentage float64   `json:"advertisingSpendPercentage,omitempty"`
	Refunds                    float64   `json:"refunds,omitempty"`
	RefundsPercentage          float64   `json:"refundsPercentage,omitempty"`
	TotalCosts                 float64   `json:"totalCosts,omitempty"`
	TotalCostsPercentage       float64   `json:"totalCostsPercentage,omitempty"`
	ShippingCredits            float64   `json:"shippingCredits,omitempty"`
	ShippingCreditsTax         float64   `json:"shippingCreditsTax,omitempty"`
	PromotionalRebates         float64   `json:"promotionalRebates,omitempty"`
	PromotionalRebatesTax      float64   `json:"promotionalRebatesTax,omitempty"`
	SalesTaxCollected          float64   `json:"salesTaxCollected,omitempty"`
	TotalCollected             float64   `json:"totalCollected,omitempty"`
	GrossMargin                float64   `json:"grossMargin,omitempty"`
	NetMargin                  float64   `json:"netMargin,omitempty"`
	NetMarginUnit              float64   `json:"netMarginUnit,omitempty"`
	ROI                        float64   `json:"roi,omitempty"`
}

type Data struct {
	Overview []Parsed `json:"overview"`
	Product  []Parsed `json:"product"`
}
