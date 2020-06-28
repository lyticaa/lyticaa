package types

import "time"

type Summary struct {
	SKU                        string
	Description                string
	Marketplace                string
	QuantitySold               int64
	ProductCosts               float64
	AdvertisingSpend           float64
	AdvertisingSpendPercentage float64
	Refunds                    float64
	Total                      float64
	OrderDate                  time.Time
}
