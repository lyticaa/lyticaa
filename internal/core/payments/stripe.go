package payments

import (
	"errors"
	"fmt"
	"os"
	"time"

	"gitlab.com/getlytica/lytica-app/internal/core/payments/types"

	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/checkout/session"
	"github.com/stripe/stripe-go/v71/invoice"
	"github.com/stripe/stripe-go/v71/price"
	"github.com/stripe/stripe-go/v71/sub"
	"github.com/stripe/stripe-go/v71/webhook"
)

var (
	stripeMonthlyPlanId    = os.Getenv("STRIPE_MONTHLY_PLAN_ID")
	stripeMonthlyProductId = os.Getenv("STRIPE_MONTHLY_PRODUCT_ID")
	stripeAnnualPlanId     = os.Getenv("STRIPE_ANNUAL_PLAN_ID")
	stripeAnnualProductId  = os.Getenv("STRIPE_ANNUAL_PRODUCT_ID")

	stripePlanProductMap = map[string]string{
		stripeMonthlyPlanId: stripeMonthlyProductId,
		stripeAnnualPlanId:  stripeAnnualProductId,
	}
)

func CheckoutSession(userId, email string, plan string) (*stripe.CheckoutSession, error) {
	setStripeKey()

	params := &stripe.CheckoutSessionParams{
		ClientReferenceID: &userId,
		CustomerEmail:     &email,
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

func SubscriptionId(session *stripe.CheckoutSession) string {
	return session.Subscription.ID
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

func ChangePlan(subscriptionId, planId string) error {
	setStripeKey()

	priceId, ok := priceIdByPlan(planId)
	if !ok {
		return errors.New("unable to find the price for the plan")
	}

	subscription, err := sub.Get(subscriptionId, nil)
	if err != nil {
		return err
	}

	params := &stripe.SubscriptionParams{
		CancelAtPeriodEnd: stripe.Bool(false),
		ProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorCreateProrations)),
		Items: []*stripe.SubscriptionItemsParams{
			{
				ID:    stripe.String(subscription.Items.Data[0].ID),
				Price: stripe.String(*priceId),
			},
		},
	}

	if _, err = sub.Update(subscriptionId, params); err != nil {
		return err
	}

	return nil
}

func priceIdByPlan(planId string) (*string, bool) {
	productId := stripePlanProductMap[planId]
	if productId == "" {
		return nil, false
	}

	priceId := priceIdByProduct(productId)
	if priceId != nil {
		return priceId, true
	}

	return nil, false
}

func priceIdByProduct(productId string) *string {
	setStripeKey()

	params := &stripe.PriceListParams{}
	params.Filters.AddFilter("product", "", productId)

	list := price.List(params)
	for list.Next() {
		return &list.Price().ID
	}

	return nil
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

func monthlyPlan() string {
	return stripeMonthlyPlanId
}

func annualPlan() string {
	return stripeAnnualPlanId
}

func setStripeKey() {
	stripe.Key = os.Getenv("STRIPE_SK")
}
