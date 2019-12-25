package account

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/sub"
)

func (a *Account) CreateCustomer(token, email string) string {
	params := &stripe.CustomerParams{
		Email: stripe.String(email),
	}
	_ = params.SetSource(token)

	stripeCustomer, _ := customer.New(params)

	return stripeCustomer.ID
}

func (a *Account) Subscribe(plan, customerId string) {
	var subParams *stripe.SubscriptionItemsParams
	subParams.Plan = stripe.String(plan)

	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerId),
		Items: []*stripe.SubscriptionItemsParams{subParams},
	}
	_, _ = sub.New(params)
}

func (a *Account) ListPlans(product string) *plan.Iter {
	params := &stripe.PlanListParams{}
	params.Filters.AddFilter("product", "=", product)
	i := plan.List(params)

	return i
}
