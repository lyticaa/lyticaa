package types

type Product struct {
	ProductID   string `json:"productID"`
	SKU         string `json:"sku"`
	Marketplace string `json:"marketplace"`
	Description string `json:"description"`
}

type Currency struct {
	CurrencyID string `json:"currencyID"`
	Symbol     string `json:"symbol"`
	Code       string `json:"code"`
}

type ExpensesTable struct {
	RowID       string  `json:"DT_RowId"`
	ProductID   string  `json:"productID,omitempty"`
	CurrencyID  string  `json:"currencyID,omitempty"`
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
