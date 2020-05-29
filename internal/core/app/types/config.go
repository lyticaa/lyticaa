package types

import "os"

type Stripe struct {
	AnnualPlanId  string
	MonthlyPlanId string
}

type Config struct {
	Stripe
	IntercomId   string
	SupportEmail string
}

func NewConfig() *Config {
	return &Config{
		Stripe: Stripe{
			AnnualPlanId:  os.Getenv("STRIPE_ANNUAL_PLAN_ID"),
			MonthlyPlanId: os.Getenv("STRIPE_MONTHLY_PLAN_ID"),
		},
		IntercomId:   os.Getenv("INTERCOM_ID"),
		SupportEmail: os.Getenv("SUPPORT_EMAIL"),
	}
}
