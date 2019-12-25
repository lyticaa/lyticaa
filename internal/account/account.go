package account

import (
	"os"

	"github.com/stripe/stripe-go"
)

type Account struct{}

func init() {
	stripe.Key = os.Getenv("STRIPE_SK")
}

func NewAccount() *Account {
	return &Account{}
}
