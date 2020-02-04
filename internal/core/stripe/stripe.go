package stripe

import (
	"fmt"
	"github.com/stripe/stripe-go/checkout/session"
	"os"

	"github.com/stripe/stripe-go"
)

func CheckoutSession(plan string) (*stripe.CheckoutSession, error) {
	stripe.Key = os.Getenv("STRIPE_SK")

	params := &stripe.CheckoutSessionParams{
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
