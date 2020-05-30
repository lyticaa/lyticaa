package payments

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"gitlab.com/getlytica/lytica-app/internal/core/payments/types"

	"github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/checkout/session"
	"github.com/stripe/stripe-go/v71/invoice"
	"github.com/stripe/stripe-go/v71/paymentmethod"
	"github.com/stripe/stripe-go/v71/price"
	"github.com/stripe/stripe-go/v71/sub"
	"github.com/stripe/stripe-go/v71/webhook"
)

const (
	CheckoutSessionCompleted    = "checkout.session.completed"
	CustomerSubscriptionCreated = "customer.subscription.created"
	CustomerSubscriptionDeleted = "customer.subscription.deleted"
	InvoiceCreated              = "invoice.created"
	InvoicePaymentFailed        = "invoice.payment_failed"
	InvoicePaymentSucceeded     = "invoice.payment_succeeded"
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
		invoices = append(
			invoices,
			types.Invoice{
				Number:   list.Invoice().Number,
				Date:     time.Unix(list.Invoice().Created, 0),
				Currency: list.Invoice().Currency,
				Amount:   FormatAmount(list.Invoice().Total),
				Status:   list.Invoice().Status,
				PDF:      list.Invoice().InvoicePDF,
			},
		)
	}

	return &invoices
}

func FormatAmount(amount int64) float64 {
	var formattedAmount float64
	if amount > 0 {
		formattedAmount = float64(amount / 100)
	}

	return formattedAmount
}

func CreateSubscription(customerId, planId string) (*stripe.Subscription, error) {
	setStripeKey()

	priceId, ok := priceIdByPlan(planId)
	if !ok {
		return nil, errors.New("unable to find the price for the plan")
	}

	method := paymentMethodByCustomer(customerId)
	if method == nil {
		return nil, errors.New("unable to find the payment method for the user")
	}

	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerId),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(*priceId),
			},
		},
		DefaultPaymentMethod: stripe.String(*method),
	}

	subscription, err := sub.New(params)
	if err != nil {
		return nil, err
	}

	return subscription, nil
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

func CancelSubscription(subscriptionId string) error {
	params := &stripe.SubscriptionParams{
		CancelAtPeriodEnd: stripe.Bool(true),
	}
	_, err := sub.Update(subscriptionId, params)
	if err != nil {
		return err
	}

	return nil
}

func EventSession(e stripe.Event) (stripe.CheckoutSession, error) {
	var session stripe.CheckoutSession
	if err := json.Unmarshal(e.Data.Raw, &session); err != nil {
		return session, err
	}

	return session, nil
}

func EventSubscription(e stripe.Event) (stripe.Subscription, error) {
	var subscription stripe.Subscription
	if err := json.Unmarshal(e.Data.Raw, &subscription); err != nil {
		return subscription, err
	}

	return subscription, nil
}

func EventInvoice(e stripe.Event) (stripe.Invoice, error) {
	var inv stripe.Invoice
	if err := json.Unmarshal(e.Data.Raw, &inv); err != nil {
		return inv, err
	}

	return inv, nil
}

func paymentMethodByCustomer(customerId string) *string {
	setStripeKey()

	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(customerId),
		Type:     stripe.String("card"),
	}
	i := paymentmethod.List(params)

	for i.Next() {
		return &i.PaymentMethod().ID
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
