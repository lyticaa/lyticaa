package webhooks

import (
	"os"
	"testing"
	"time"

	"github.com/lyticaa/lyticaa-app/internal/web/pkg/payments/types"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/stripe/stripe-go/v71"
	. "gopkg.in/check.v1"
	"syreclabs.com/go/faker"
)

type mockGateway struct{}

func (m *mockGateway) CheckoutSessions(string, string) (*[]string, error) { return &[]string{}, nil }

func (m *mockGateway) CheckoutSession(string, string, string) (*stripe.CheckoutSession, error) {
	return &stripe.CheckoutSession{}, nil
}

func (m *mockGateway) CustomerRefId(session *stripe.CheckoutSession) string   { return "" }
func (m *mockGateway) CustomerId(session *stripe.CheckoutSession) *string     { return nil }
func (m *mockGateway) SubscriptionId(session *stripe.CheckoutSession) *string { return nil }
func (m *mockGateway) PlanId(session *stripe.CheckoutSession) *string         { return nil }

func (m *mockGateway) ConstructEvent(body []byte, sig string) (stripe.Event, error) {
	return stripe.Event{}, nil
}

func (m *mockGateway) InvoicesByUser(subscriptionId string) *types.Invoices {
	invoices := types.Invoices{{
		Number:   faker.RandomString(10),
		Date:     time.Now(),
		Currency: "USD",
		Amount:   float64(faker.Commerce().Price()),
		Status:   "paid",
		PDF:      faker.Internet().Url(),
	}}

	return &invoices
}

func (m *mockGateway) FormatAmount(amount int64) float64 { return 0.0 }

func (m *mockGateway) CreateSubscription(customerId string, planId string) (*stripe.Subscription, error) {
	return &stripe.Subscription{}, nil
}

func (m *mockGateway) ChangePlan(customerId string, planId string) error { return nil }
func (m *mockGateway) CancelSubscription(subscriptionId string) error    { return nil }

func (m *mockGateway) EventSession(event stripe.Event) (stripe.CheckoutSession, error) {
	return stripe.CheckoutSession{}, nil
}

func (m *mockGateway) EventSubscription(event stripe.Event) (stripe.Subscription, error) {
	return stripe.Subscription{}, nil
}

func (m *mockGateway) EventInvoice(event stripe.Event) (stripe.Invoice, error) {
	return stripe.Invoice{}, nil
}

type webhooksSuite struct {
	w *Webhooks
}

var _ = Suite(&webhooksSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *webhooksSuite) SetUpSuite(c *C) {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	c.Assert(err, IsNil)

	s.w = NewWebhooks(db, log.With().Logger(), &mockGateway{})
}

func (s *webhooksSuite) TestWebhooks(c *C) {
	c.Assert(s.w.db, NotNil)
	c.Assert(s.w.logger, NotNil)
	c.Assert(s.w.stripe, NotNil)
}

func (s *webhooksSuite) TearDownSuite(c *C) {}
