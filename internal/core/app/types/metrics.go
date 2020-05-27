package types

type Sparkline struct {
	Plots []int64 `json:"plots"`
}

type Flot struct {
	Plots [][]int64 `json:"plots"`
}

type Chart struct {
	Sparkline Sparkline `json:"sparkline,omitempty"`
	Flot      Flot      `json:"flot,omitempty"`
}

type Card struct {
	Total int64  `json:"total"`
	Diff  uint64 `json:"diff"`
	Chart Chart  `json:"chart"`
}

type Metrics struct {
	TotalSales         Card `json:"total_sales,omitempty"`
	UnitsSold          Card `json:"units_sold,omitempty"`
	AmazonCosts        Card `json:"amazon_costs,omitempty"`
	AdvertisingSpend   Card `json:"advertising_spend,omitempty"`
	Refunds            Card `json:"refunds,omitempty"`
	ShippingCredits    Card `json:"shipping_credits,omitempty"`
	PromotionalRebates Card `json:"promotional_rebates,omitempty"`
	TotalCosts         Card `json:"total_costs,omitempty"`
	NetMargin          Card `json:"net_margin,omitempty"`
}
