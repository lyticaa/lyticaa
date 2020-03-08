package payments

import (
	"fmt"
	"os"

	"gitlab.com/getlytica/lytica/internal/models"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
	"github.com/stripe/stripe-go/webhook"
)

func CheckoutSession(user models.User, plan string) (*stripe.CheckoutSession, error) {
	stripe.Key = os.Getenv("STRIPE_SK")

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

func ConstructEvent(body []byte, sig string) (stripe.Event, error) {
	return webhook.ConstructEvent(body, sig, os.Getenv("STRIPE_WHSEC"))
}
