package payments

import (
	"fmt"
	"os"
	"time"

	"gitlab.com/getlytica/lytica-app/internal/core/payments/types"
	"gitlab.com/getlytica/lytica-app/internal/models"

	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/checkout/session"
	"github.com/stripe/stripe-go/v71/invoice"
	"github.com/stripe/stripe-go/v71/price"
	"github.com/stripe/stripe-go/v71/webhook"
)

func CheckoutSession(user models.User, plan string) (*stripe.CheckoutSession, error) {
	setStripeKey()

	params := &stripe.CheckoutSessionParams{
		ClientReferenceID: &user.UserId,
		CustomerEmail:     &user.Email,
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			Items: []*stripe.CheckoutSessionSubscriptionDataItemsParams{
				&stripe.CheckoutSessionSubscriptionDataItemsParams{
					Plan: stripe.String(getPlan(plan)),
				},
			},
			TrialFromPlan: stripe.Bool(true),
		},
		SuccessURL: stripe.String(fmt.Sprintf("%v%v", os.Getenv("BASE_URL"), os.Getenv("STRIPE_SUCCESS_URI"))),
		CancelURL:  stripe.String(fmt.Sprintf("%v%v", os.Getenv("BASE_URL"), os.Getenv("STRIPE_CANCEL_URI"))),
	}

	return session.New(params)
}

func CustomerRefId(session *stripe.CheckoutSession) string {
	return session.ClientReferenceID
}

func CustomerId(session *stripe.CheckoutSession) string {
	return session.Customer.ID
}

func PlanId(session *stripe.CheckoutSession) string {
	return session.Subscription.Plan.ID
}

func ConstructEvent(body []byte, sig string) (stripe.Event, error) {
	return webhook.ConstructEvent(body, sig, os.Getenv("STRIPE_WHSEC"))
}

func InvoicesByUser(customer string) *types.Invoices {
	setStripeKey()
	var invoices types.Invoices

	params := &stripe.InvoiceListParams{Customer: &customer}
	list := invoice.List(params)

	for list.Next() {
		var formattedAmount float64
		total := list.Invoice().Total
		if total > 0 {
			formattedAmount = float64(total / 100)
		}

		invoices = append(
			invoices,
			types.Invoice{
				Number:   list.Invoice().Number,
				Date:     time.Unix(list.Invoice().Created, 0),
				Currency: list.Invoice().Currency,
				Amount:   formattedAmount,
				Status:   list.Invoice().Status,
				PDF:      list.Invoice().InvoicePDF,
			},
		)
	}

	return &invoices
}

func Plans(user models.User) *types.Plans {
	var plans types.Plans

	for _, plan := range getPlans() {
		p := types.Plan{ID: plan}

		for _, product := range getProducts() {
			p.Products = append(p.Products,
				types.Product{
					ID:     product,
					Prices: *pricesByProduct(product),
				},
			)
		}

		plans = append(plans, p)
	}

	return &plans
}

func pricesByProduct(product string) *[]types.Price {
	setStripeKey()

	params := &stripe.PriceListParams{}
	params.Filters.AddFilter("product", "=", product)

	var prices []types.Price

	list := price.List(params)
	for list.Next() {
		prices = append(
			prices,
			types.Price{
				ID:     list.Price().ID,
				Amount: list.Price().UnitAmountDecimal,
			},
		)
	}

	return &prices
}

func getPlan(plan string) string {
	switch plan {
	case "monthly":
		return monthlyPlan()
	case "annual":
		return annualPlan()
	}

	return monthlyPlan()
}

func getPlans() []string {
	return []string{
		monthlyPlan(),
		annualPlan(),
	}
}

func getProducts() []string {
	return []string{
		monthlyProduct(),
		annualProduct(),
	}
}

func monthlyPlan() string {
	return os.Getenv("STRIPE_MONTHLY_PLAN_ID")
}

func annualPlan() string {
	return os.Getenv("STRIPE_ANNUAL_PLAN_ID")
}

func monthlyProduct() string {
	return os.Getenv("STRIPE_MONTHLY_PRODUCT_ID")
}

func annualProduct() string {
	return os.Getenv("STRIPE_ANNUAL_PRODUCT_ID")
}

func setStripeKey() {
	stripe.Key = os.Getenv("STRIPE_SK")
}
