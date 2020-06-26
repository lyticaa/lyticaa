package types

import (
	"time"
)

type NotificationTable struct {
	Notification string    `json:"notification"`
	Date         time.Time `json:"date"`
}

type Notifications struct {
	Data         []NotificationTable `json:"data"`
	Draw         int64               `json:"draw"`
	RecordsTotal int64               `json:"recordsTotal"`
}

type InvoiceTable struct {
	Number      string `json:"number"`
	Date        string `json:"date"`
	Amount      string `json:"amount"`
	Status      string `json:"status"`
	StatusClass string `json:"statusClass"`
	PDF         string `json:"pdf"`
}

type Invoices struct {
	Data            []InvoiceTable `json:"data"`
	Draw            int64          `json:"draw"`
	RecordsTotal    int64          `json:"recordsTotal"`
	RecordsFiltered int64          `json:"recordsFiltered"`
}

type Subscription struct {
	Invoices `json:"invoices"`
}
