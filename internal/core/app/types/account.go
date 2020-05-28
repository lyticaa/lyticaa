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
