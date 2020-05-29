package payments

import (
	"fmt"
	"os"
	"time"

	"gitlab.com/getlytica/lytica-app/internal/core/payments/types"
	"gitlab.com/getlytica/lytica-app/internal/models"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
	"github.com/stripe/stripe-go/invoice"
	"github.com/stripe/stripe-go/webhook"
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
					Plan: stripe.String(GetPlan(plan)),
				},
			},
			TrialFromPlan: stripe.Bool(true),
		},
		SuccessURL: stripe.String(fmt.Sprintf("%v%v", os.Getenv("BASE_URL"), os.Getenv("STRIPE_SUCCESS_URI"))),
		CancelURL:  stripe.String(fmt.Sprintf("%v%v", os.Getenv("BASE_URL"), os.Getenv("STRIPE_CANCEL_URI"))),
	}

	return session.New(params)
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

		invoices = append(invoices, types.Invoice{
			Number:   list.Invoice().Number,
			Date:     time.Unix(list.Invoice().Created, 0),
			Currency: list.Invoice().Currency,
			Amount:   formattedAmount,
			Status:   list.Invoice().Status,
			PDF:      list.Invoice().InvoicePDF,
		})
	}

	return &invoices
}

func GetPlan(plan string) string {
	switch plan {
	case "monthly":
		return Monthly()
	case "annual":
		return Annual()
	}

	return Monthly()
}

func Monthly() string {
	return os.Getenv("STRIPE_MONTHLY_PLAN_ID")
}

func Annual() string {
	return os.Getenv("STRIPE_ANNUAL_PLAN_ID")
}

func CustomerRefId(session *stripe.CheckoutSession) string {
	return session.ClientReferenceID
}

func CustomerId(session *stripe.CheckoutSession) string {
	return session.Customer.ID
}

func SubscriptionId(session *stripe.CheckoutSession) string {
	return session.Subscription.ID
}

func PlanId(session *stripe.CheckoutSession) string {
	return session.Subscription.Plan.ID
}

func ConstructEvent(body []byte, sig string) (stripe.Event, error) {
	return webhook.ConstructEvent(body, sig, os.Getenv("STRIPE_WHSEC"))
}

func setStripeKey() {
	stripe.Key = os.Getenv("STRIPE_SK")
}
