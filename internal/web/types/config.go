package types

import "os"

type Stripe struct {
	PK            string
	AnnualPlanID  string
	MonthlyPlanID string
}

type Config struct {
	Stripe
	IntercomID   string
	SupportEmail string
}

func NewConfig() *Config {
	return &Config{
		Stripe: Stripe{
			PK:            os.Getenv("STRIPE_PK"),
			AnnualPlanID:  os.Getenv("STRIPE_ANNUAL_PLAN_ID"),
			MonthlyPlanID: os.Getenv("STRIPE_MONTHLY_PLAN_ID"),
		},
		IntercomID:   os.Getenv("INTERCOM_ID"),
		SupportEmail: os.Getenv("SUPPORT_EMAIL"),
	}
}
