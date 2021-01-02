package types

type CohortTable struct {
	SKU                string  `json:"sku"`
	Description        string  `json:"description"`
	Marketplace        string  `json:"marketplace"`
	TotalSales         float64 `json:"totalSales"`
	Quantity           int64   `json:"quantity"`
	AmazonCosts        float64 `json:"amazonCosts"`
	ProductCosts       float64 `json:"productCosts"`
	AdvertisingSpend   float64 `json:"advertisingSpend"`
	Refunds            float64 `json:"refunds"`
	ShippingCredits    float64 `json:"shippingCredits"`
	PromotionalRebates float64 `json:"promotionalRebates"`
	TotalCosts         float64 `json:"totalCosts"`
	NetMargin          float64 `json:"netMargin"`
}

type Cohort struct {
	TotalSales       Card          `json:"totalSales"`
	AmazonCosts      Card          `json:"amazonCosts"`
	ProductCosts     Card          `json:"productCosts"`
	AdvertisingSpend Card          `json:"advertisingSpend"`
	NetMargin        Card          `json:"netMargin"`
	Data             []CohortTable `json:"data"`
	Draw             int64         `json:"draw"`
	RecordsTotal     int64         `json:"recordsTotal"`
	RecordsFiltered  int64         `json:"recordsFiltered"`
}
