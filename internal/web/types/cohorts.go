package types

type CohortSummary struct {
	TotalSales       Card `json:"totalSales"`
	AmazonCosts      Card `json:"amazonCosts"`
	ProductCosts     Card `json:"productCosts"`
	AdvertisingSpend Card `json:"advertisingSpend"`
	Margin           Card `json:"margin"`
}

type CohortTable struct {
	SKU                string  `json:"sku"`
	ASIN               string  `json:"asin"`
	ProductName        string  `json:"productName"`
	TotalSales         float64 `json:"totalSales"`
	GrossQuantitySold  int64   `json:"grossQuantitySold"`
	NetQuantitySold    int64   `json:"netQuantitySold"`
	AmazonCosts        float64 `json:"amazonCosts"`
	ProductCosts       float64 `json:"productCosts"`
	CostOfCoupons      float64 `json:"costOfCoupons"`
	AdvertisingSpend   float64 `json:"advertisingSpend"`
	Coupons            int64   `json:"coupons"`
	Refunds            float64 `json:"refunds"`
	ShippingCredits    float64 `json:"shippingCredits"`
	PromotionalRebates float64 `json:"promotionalRebates"`
	TotalCosts         float64 `json:"totalCosts"`
	NetMargin          float64 `json:"netMargin"`
}

type Cohort struct {
	CohortSummary   `json:"cohortSummary"`
	Data            []CohortTable `json:"data"`
	Draw            int64         `json:"draw"`
	RecordsTotal    int64         `json:"recordsTotal"`
	RecordsFiltered int64         `json:"recordsFiltered"`
}
