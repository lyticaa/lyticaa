package types

import "time"

type Summary struct {
	Total       float64
	Marketplace string
	OrderDate   time.Time
}
