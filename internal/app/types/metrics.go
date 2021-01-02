package types

type Total struct {
	Value float64 `json:"value"`
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
	SKU         string `json:"sku"`
	Description string `json:"description"`
	Marketplace string `json:"marketplace"`
	Quantity    int64  `json:"quantity"`
}

type UnitsSold struct {
	Metrics
	Data []UnitsSoldTable `json:"data"`
}

type AmazonCostsTable struct {
	SKU         string  `json:"sku"`
	Description string  `json:"type"`
	Marketplace string  `json:"marketplace"`
	AmazonCosts float64 `json:"amazonCosts"`
}

type AmazonCosts struct {
	Metrics
	Data []AmazonCostsTable `json:"data"`
}

type ProductCostsTable struct {
	SKU              string  `json:"sku"`
	Description      string  `json:"description"`
	Marketplace      string  `json:"marketplace"`
	Quantity         int64   `json:"quantity"`
	ProductCosts     float64 `json:"productCosts"`
	AdvertisingSpend float64 `json:"advertisingSpend"`
	Refunds          float64 `json:"refunds"`
	TotalCosts       float64 `json:"totalCosts"`
}

type ProductCosts struct {
	Metrics
	Data []ProductCostsTable `json:"data"`
}

type AdvertisingSpendTable struct {
	SKU                        string  `json:"sku"`
	Description                string  `json:"description"`
	Marketplace                string  `json:"marketplace"`
	AdvertisingSpend           float64 `json:"advertisingSpend"`
	AdvertisingSpendPercentage float64 `json:"advertisingSpendPercentage"`
}

type AdvertisingSpend struct {
	Metrics
	Data []AdvertisingSpendTable `json:"data"`
}

type RefundsTable struct {
	SKU               string  `json:"sku"`
	Description       string  `json:"description"`
	Marketplace       string  `json:"marketplace"`
	Refunds           float64 `json:"refunds"`
	RefundsPercentage float64 `json:"refundsPercentage"`
}

type Refunds struct {
	Metrics
	Data []RefundsTable `json:"data"`
}

type ShippingCreditsTable struct {
	SKU                string  `json:"sku"`
	Description        string  `json:"description"`
	Marketplace        string  `json:"marketplace"`
	ShippingCredits    float64 `json:"shippingCredits"`
	ShippingCreditsTax float64 `json:"shippingCreditsTax"`
}

type ShippingCredits struct {
	Metrics
	Data []ShippingCreditsTable `json:"data"`
}

type PromotionalRebatesTable struct {
	SKU                   string  `json:"sku"`
	Description           string  `json:"description"`
	Marketplace           string  `json:"marketplace"`
	PromotionalRebates    float64 `json:"promotionalRebates"`
	PromotionalRebatesTax float64 `json:"promotionalRebatesTax"`
}

type PromotionalRebates struct {
	Metrics
	Data []PromotionalRebatesTable `json:"data"`
}

type TotalCostsTable struct {
	SKU                  string  `json:"sku"`
	Description          string  `json:"description"`
	Marketplace          string  `json:"marketplace"`
	AmazonCosts          float64 `json:"amazonCosts"`
	ProductCosts         float64 `json:"productCosts"`
	ProductCostPerUnit   float64 `json:"productCostPerUnit"`
	TotalCosts           float64 `json:"totalCosts"`
	TotalCostsPercentage float64 `json:"totalCostsPercentage"`
}

type TotalCosts struct {
	Metrics
	Data []TotalCostsTable `json:"data"`
}

type GrossMarginTable struct {
	SKU                string  `json:"sku"`
	Description        string  `json:"description"`
	Marketplace        string  `json:"marketplace"`
	ProductCosts       float64 `json:"productCosts"`
	Quantity           int64   `json:"quantity"`
	TotalSales         float64 `json:"totalSales"`
	AmazonCosts        float64 `json:"amazonCosts"`
	ShippingCredits    float64 `json:"shippingCredits"`
	PromotionalRebates float64 `json:"promotionalRebates"`
	GrossMargin        float64 `json:"grossMargin"`
	SalesTaxCollected  float64 `json:"salesTaxCollected"`
	TotalCollected     float64 `json:"totalCollected"`
}

type GrossMargin struct {
	Metrics
	Data []GrossMarginTable `json:"data"`
}

type NetMarginTable struct {
	SKU           string  `json:"sku"`
	Description   string  `json:"description"`
	Marketplace   string  `json:"marketplace"`
	GrossMargin   float64 `json:"grossMargin"`
	TotalCosts    float64 `json:"totalCosts"`
	NetMargin     float64 `json:"netMargin"`
	Quantity      int64   `json:"quantity"`
	NetMarginUnit float64 `json:"netMarginUnit"`
	ROI           float64 `json:"roi"`
}

type NetMargin struct {
	Metrics
	Data []NetMarginTable `json:"data"`
}
