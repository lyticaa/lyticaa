package payments

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/lyticaa/lyticaa-app/internal/web/pkg/payments/types"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v72/invoice"
	"github.com/stripe/stripe-go/v72/paymentmethod"
	"github.com/stripe/stripe-go/v72/plan"
	"github.com/stripe/stripe-go/v72/price"
	"github.com/stripe/stripe-go/v72/sub"
	"github.com/stripe/stripe-go/v72/webhook"
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
	stripeMonthlyPlanID    = os.Getenv("STRIPE_MONTHLY_PLAN_ID")
	stripeMonthlyProductID = os.Getenv("STRIPE_MONTHLY_PRODUCT_ID")
	stripeAnnualPlanID     = os.Getenv("STRIPE_ANNUAL_PLAN_ID")
	stripeAnnualProductID  = os.Getenv("STRIPE_ANNUAL_PRODUCT_ID")

	stripePlanProductMap = map[string]string{
		stripeMonthlyPlanID: stripeMonthlyProductID,
		stripeAnnualPlanID:  stripeAnnualProductID,
	}
)

type StripeGateway interface {
	CheckoutSessions(string, string) (*[]string, error)
	CheckoutSession(string, string, string) (*stripe.CheckoutSession, error)
	CustomerRefID(*stripe.CheckoutSession) string
	CustomerID(*stripe.CheckoutSession) *string
	SubscriptionID(*stripe.CheckoutSession) *string
	PlanID(*stripe.CheckoutSession) *string
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

func init() {
	stripe.Key = os.Getenv("STRIPE_SK")
}

func NewStripePayments() *StripePayments {
	return &StripePayments{
		key: os.Getenv("STRIPE_SK"),
	}
}

func (p *StripePayments) CheckoutSessions(userID, email string) (*[]string, error) {
	monthly, err := p.CheckoutSession(userID, email, monthlyPaymentPlan)
	if err != nil {
		return nil, err
	}

	annual, err := p.CheckoutSession(userID, email, annualPaymentPlan)
	if err != nil {
		return nil, err
	}

	return &[]string{monthly.ID, annual.ID}, nil
}

func (p *StripePayments) CheckoutSession(userID, email string, plan string) (*stripe.CheckoutSession, error) {
	params := &stripe.CheckoutSessionParams{
		ClientReferenceID: &userID,
		CustomerEmail:     &email,
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		SubscriptionData: &stripe.CheckoutSessionSubscriptionDataParams{
			Items: []*stripe.CheckoutSessionSubscriptionDataItemsParams{{
				Plan: stripe.String(p.getPlan(plan)),
			}},
			TrialFromPlan: stripe.Bool(true),
		},
		SuccessURL: stripe.String(fmt.Sprintf("%v%v", os.Getenv("BASE_URL"), os.Getenv("STRIPE_SUCCESS_URI"))),
		CancelURL:  stripe.String(fmt.Sprintf("%v%v", os.Getenv("BASE_URL"), os.Getenv("STRIPE_CANCEL_URI"))),
	}

	return session.New(params)
}

func (p *StripePayments) CustomerRefID(session *stripe.CheckoutSession) string {
	return session.ClientReferenceID
}

func (p *StripePayments) CustomerID(session *stripe.CheckoutSession) *string {
	if session.Customer != nil {
		return &session.Customer.ID
	}

	return nil
}

func (p *StripePayments) SubscriptionID(session *stripe.CheckoutSession) *string {
	if session.Subscription != nil {
		return &session.Subscription.ID
	}

	return nil
}

func (p *StripePayments) PlanID(session *stripe.CheckoutSession) *string {
	if len(session.LineItems.Data) > 0 {
		return p.planIDByProduct(session.LineItems.Data[0].Price.Product.ID)
	}

	return nil
}

func (p *StripePayments) ConstructEvent(body []byte, sig string) (stripe.Event, error) {
	return webhook.ConstructEvent(body, sig, os.Getenv("STRIPE_WHSEC"))
}

func (p *StripePayments) InvoicesByUser(customer string) *types.Invoices {
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

func (p *StripePayments) FormatAmount(amount int64) float64 {
	var formattedAmount float64
	if amount > 0 {
		formattedAmount = float64(amount / 100)
	}

	return formattedAmount
}

func (p *StripePayments) CreateSubscription(customerID, planID string) (*stripe.Subscription, error) {
	priceID, ok := p.priceIDByPlan(planID)
	if !ok {
		return nil, errors.New("failed to find the price for the plan")
	}

	method := p.paymentMethodByCustomer(customerID)
	if method == nil {
		return nil, errors.New("failed to find the payment method for the user")
	}

	params := &stripe.SubscriptionParams{
		Customer: stripe.String(customerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(*priceID),
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

func (p *StripePayments) ChangePlan(subscriptionID, planID string) error {
	priceID, ok := p.priceIDByPlan(planID)
	if !ok {
		return errors.New("failed to find the price for the plan")
	}

	subscription, err := sub.Get(subscriptionID, nil)
	if err != nil {
		return err
	}

	params := &stripe.SubscriptionParams{
		CancelAtPeriodEnd: stripe.Bool(false),
		ProrationBehavior: stripe.String(string(stripe.SubscriptionProrationBehaviorCreateProrations)),
		Items: []*stripe.SubscriptionItemsParams{
			{
				ID:    stripe.String(subscription.Items.Data[0].ID),
				Price: stripe.String(*priceID),
			},
		},
	}

	if _, err = sub.Update(subscriptionID, params); err != nil {
		return err
	}

	return nil
}

func (p *StripePayments) CancelSubscription(subscriptionID string) error {
	params := &stripe.SubscriptionParams{
		CancelAtPeriodEnd: stripe.Bool(true),
	}
	_, err := sub.Update(subscriptionID, params)
	if err != nil {
		return err
	}

	return nil
}

func (p *StripePayments) EventSession(e stripe.Event) (stripe.CheckoutSession, error) {
	var session stripe.CheckoutSession
	if err := json.Unmarshal(e.Data.Raw, &session); err != nil {
		return session, err
	}

	return session, nil
}

func (p *StripePayments) EventSubscription(e stripe.Event) (stripe.Subscription, error) {
	var subscription stripe.Subscription
	if err := json.Unmarshal(e.Data.Raw, &subscription); err != nil {
		return subscription, err
	}

	return subscription, nil
}

func (p *StripePayments) EventInvoice(e stripe.Event) (stripe.Invoice, error) {
	var inv stripe.Invoice
	if err := json.Unmarshal(e.Data.Raw, &inv); err != nil {
		return inv, err
	}

	return inv, nil
}

func (p *StripePayments) paymentMethodByCustomer(customerID string) *string {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(customerID),
		Type:     stripe.String("card"),
	}
	i := paymentmethod.List(params)

	for i.Next() {
		return &i.PaymentMethod().ID
	}

	return nil
}

func (p *StripePayments) priceIDByPlan(planID string) (*string, bool) {
	productID := stripePlanProductMap[planID]
	if productID == "" {
		return nil, false
	}

	priceID := p.priceIDByProduct(productID)
	if priceID != nil {
		return priceID, true
	}

	return nil, false
}

func (p *StripePayments) priceIDByProduct(productID string) *string {
	params := &stripe.PriceListParams{}
	params.Filters.AddFilter("product", "", productID)

	list := price.List(params)
	for list.Next() {
		return &list.Price().ID
	}

	return nil
}

func (p *StripePayments) planIDByProduct(productID string) *string {
	params := &stripe.PlanListParams{}
	params.Filters.AddFilter("product", "", productID)

	list := plan.List(params)
	for list.Next() {
		return &list.Plan().ID
	}

	return nil
}

func (p *StripePayments) getPlan(plan string) string {
	switch plan {
	case monthlyPaymentPlan:
		return p.monthlyPlan()
	case annualPaymentPlan:
		return p.annualPlan()
	}

	return p.monthlyPlan()
}

func (p *StripePayments) monthlyPlan() string {
	return stripeMonthlyPlanID
}

func (p *StripePayments) annualPlan() string {
	return stripeAnnualPlanID
}
