package types

import (
	"time"
)

type ExpensesTable struct {
	SKU         string    `json:"sku,omitempty"`
	Marketplace string    `json:"marketplace,omitempty"`
	Description string    `json:"description"`
	FromDate    time.Time `json:"fromDate,omitempty"`
	DateTime    time.Time `json:"dateTime,omitempty"`
	Cost        float64   `json:"cost"`
}

type Expenses struct {
	Data            []ExpensesTable `json:"data"`
	Draw            int64           `json:"draw"`
	RecordsTotal    int64           `json:"recordsTotal"`
	RecordsFiltered int64           `json:"recordsFiltered"`
}
