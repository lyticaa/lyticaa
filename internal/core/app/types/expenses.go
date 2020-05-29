package types

import (
	"time"
)

type ExpensesTable struct {
	SKU         string    `json:"sku,omitempty"`
	ASIN        string    `json:"asin,omitempty"`
	ProductName string    `json:"productName,omitempty"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	Type        string    `json:"type"`
	Cost        float64   `json:"cost"`
	Currency    string    `json:"currency"`
}

type Expenses struct {
	Data            []ExpensesTable `json:"data"`
	Draw            int64           `json:"draw"`
	RecordsTotal    int64           `json:"recordsTotal"`
	RecordsFiltered int64           `json:"recordsFiltered"`
}
