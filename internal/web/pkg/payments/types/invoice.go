package types

import (
	"time"

	"github.com/stripe/stripe-go/v72"
)

type Invoices []Invoice
type Invoice struct {
	Number   string
	Date     time.Time
	Currency stripe.Currency
	Amount   float64
	Status   stripe.InvoiceStatus
	PDF      string
}
