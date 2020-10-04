package payments

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"gitlab.com/lyticaa/lyticaa-app/internal/web/pkg/payments/types"

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

type StripeGateway interface {
	CheckoutSession(string, string, string) (*stripe.CheckoutSession, error)
	CustomerRefId(*stripe.CheckoutSession) string
	CustomerId(*stripe.CheckoutSession) *string
	SubscriptionId(*stripe.CheckoutSession) *string
	PlanId(*stripe.CheckoutSession) *string
	ConstructEvent([]byte, string) (stripe.Event, error)
	InvoicesByUser(string) *types.Invoices
	FormatAmount(int64) float64
	CreateSubscription(string, string) (*stripe.Subscription, error)
	ChangePlan(string, string) error
	CancelSubscription(string) error
	EventSession(stripe.Event) (stripe.CheckoutSession, error)
	EventSubscription(stripe.Event) (stripe.Subscription, error)
	EventInvoice(stripe.Event) (stripe.Invoice, error)
}

func NewStripePayments() *Payments {
	stripe.Key = os.Getenv("STRIPE_SK")

	return &Payments{}
}

func (p *Payments) CheckoutSession(userId, email string, plan string) (*stripe.CheckoutSession, error) {
	params := &stripe.CheckoutSessionParams{
		ClientReferenceID: &userId,
		CustomerEmail:     &email,
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			Items: []*stripe.CheckoutSessionSubscriptionDataItemsParams{{
				Plan: stripe.String(p.getPlan(plan)),
			},
			},
			TrialFromPlan: stripe.Bool(true),
		},
		SuccessURL: stripe.String(fmt.Sprintf("%v%v", os.Getenv("BASE_URL"), os.Getenv("STRIPE_SUCCESS_URI"))),
		CancelURL:  stripe.String(fmt.Sprintf("%v%v", os.Getenv("BASE_URL"), os.Getenv("STRIPE_CANCEL_URI"))),
	}

	return session.New(params)
}

func (p *Payments) CustomerRefId(session *stripe.CheckoutSession) string {
	return session.ClientReferenceID
}

func (p *Payments) CustomerId(session *stripe.CheckoutSession) *string {
	if session.Customer != nil {
		return &session.Customer.ID
	}

	return nil
}

func (p *Payments) SubscriptionId(session *stripe.CheckoutSession) *string {
	if session.Subscription != nil {
		return &session.Subscription.ID
	}

	return nil
}

func (p *Payments) PlanId(session *stripe.CheckoutSession) *string {
	if session.Subscription != nil {
		return &session.Subscription.Plan.ID
	}

	return nil
}

func (p *Payments) ConstructEvent(body []byte, sig string) (stripe.Event, error) {
	return webhook.ConstructEvent(body, sig, os.Getenv("STRIPE_WHSEC"))
}

func (p *Payments) InvoicesByUser(customer string) *types.Invoices {
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
				Amount:   p.FormatAmount(list.Invoice().Total),
				Status:   list.Invoice().Status,
				PDF:      list.Invoice().InvoicePDF,
			},
		)
	}

	return &invoices
}

func (p *Payments) FormatAmount(amount int64) float64 {
	var formattedAmount float64
	if amount > 0 {
		formattedAmount = float64(amount / 100)
	}

	return formattedAmount
}

func (p *Payments) CreateSubscription(customerId, planId string) (*stripe.Subscription, error) {
	priceId, ok := p.priceIdByPlan(planId)
	if !ok {
		return nil, errors.New("failed to find the price for the plan")
	}

	method := p.paymentMethodByCustomer(customerId)
	if method == nil {
		return nil, errors.New("failed to find the payment method for the user")
	}

	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerId),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(*priceId),
			},
		},
		DefaultPaymentMethod: stripe.String(*method),
		TrialFromPlan:        stripe.Bool(true),
	}

	subscription, err := sub.New(params)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}

func (p *Payments) ChangePlan(subscriptionId, planId string) error {
	priceId, ok := p.priceIdByPlan(planId)
	if !ok {
		return errors.New("failed to find the price for the plan")
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

func (p *Payments) CancelSubscription(subscriptionId string) error {
	params := &stripe.SubscriptionParams{
		CancelAtPeriodEnd: stripe.Bool(true),
	}
	_, err := sub.Update(subscriptionId, params)
	if err != nil {
		return err
	}

	return nil
}

func (p *Payments) EventSession(e stripe.Event) (stripe.CheckoutSession, error) {
	var session stripe.CheckoutSession
	if err := json.Unmarshal(e.Data.Raw, &session); err != nil {
		return session, err
	}

	return session, nil
}

func (p *Payments) EventSubscription(e stripe.Event) (stripe.Subscription, error) {
	var subscription stripe.Subscription
	if err := json.Unmarshal(e.Data.Raw, &subscription); err != nil {
		return subscription, err
	}

	return subscription, nil
}

func (p *Payments) EventInvoice(e stripe.Event) (stripe.Invoice, error) {
	var inv stripe.Invoice
	if err := json.Unmarshal(e.Data.Raw, &inv); err != nil {
		return inv, err
	}

	return inv, nil
}

func (p *Payments) paymentMethodByCustomer(customerId string) *string {
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

func (p *Payments) priceIdByPlan(planId string) (*string, bool) {
	productId := stripePlanProductMap[planId]
	if productId == "" {
		return nil, false
	}

	priceId := p.priceIdByProduct(productId)
	if priceId != nil {
		return priceId, true
	}

	return nil, false
}

func (p *Payments) priceIdByProduct(productId string) *string {
	params := &stripe.PriceListParams{}
	params.Filters.AddFilter("product", "", productId)

	list := price.List(params)
	for list.Next() {
		return &list.Price().ID
	}

	return nil
}

func (p *Payments) getPlan(plan string) string {
	switch plan {
	case "monthly":
		return p.monthlyPlan()
	case "annual":
		return p.annualPlan()
	}

	return p.monthlyPlan()
}

func (p *Payments) monthlyPlan() string {
	return stripeMonthlyPlanId
}

func (p *Payments) annualPlan() string {
	return stripeAnnualPlanId
}
