package types

import (
	"time"
)

type Summary struct {
	Date        time.Time
	Marketplace string
	Total       float64
}
