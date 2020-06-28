package types

import "time"

type Summary struct {
	SKU                        string
	Description                string
	Marketplace                string
	QuantitySold               int64
	ProductCosts               float64
	ProductCostsUnit           float64
	AmazonCosts                float64
	AdvertisingSpend           float64
	AdvertisingSpendPercentage float64
	Refunds                    float64
	RefundsPercentage          float64
	ShippingCredits            float64
	ShippingCreditsTax         float64
	PromotionalRebates         float64
	PromotionalRebatesTax      float64
	TotalCosts                 float64
	TotalCostsPercentage       float64
	TotalRevenue               float64
	GrossMargin                float64
	NetMargin                  float64
	NetMarginUnit              float64
	SalesTaxCollected          float64
	ROI                        float64
	Total                      float64
	Date                       time.Time
}
