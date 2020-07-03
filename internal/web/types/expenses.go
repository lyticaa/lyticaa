package types

type Product struct {
	ProductId   string `json:"productId"`
	SKU         string `json:"sku"`
	Marketplace string `json:"marketplace"`
	Description string `json:"description"`
}

type Currency struct {
	CurrencyId string `json:"currencyId"`
	Symbol     string `json:"symbol"`
	Code       string `json:"code"`
}

type ExpensesTable struct {
	RowId       string  `json:"DT_RowId"`
	ProductId   string  `json:"productId,omitempty"`
	CurrencyId  string  `json:"currencyId,omitempty"`
	SKU         string  `json:"sku,omitempty"`
	Marketplace string  `json:"marketplace,omitempty"`
	Description string  `json:"description"`
	FromDate    string  `json:"fromDate,omitempty"`
	DateTime    string  `json:"dateTime,omitempty"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency,omitempty"`
}

type Expenses struct {
	Data            []ExpensesTable `json:"data"`
	Draw            int64           `json:"draw"`
	RecordsTotal    int64           `json:"recordsTotal"`
	RecordsFiltered int64           `json:"recordsFiltered"`
}
