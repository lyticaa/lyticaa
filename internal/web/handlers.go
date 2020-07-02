package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/web/controllers/account"
	"gitlab.com/getlytica/lytica-app/internal/web/controllers/api"
	"gitlab.com/getlytica/lytica-app/internal/web/controllers/auth"
	"gitlab.com/getlytica/lytica-app/internal/web/controllers/cohorts"
	"gitlab.com/getlytica/lytica-app/internal/web/controllers/dashboard"
	"gitlab.com/getlytica/lytica-app/internal/web/controllers/expenses"
	"gitlab.com/getlytica/lytica-app/internal/web/controllers/forecast"
	"gitlab.com/getlytica/lytica-app/internal/web/controllers/metrics"
	"gitlab.com/getlytica/lytica-app/internal/web/controllers/profit_loss"
	"gitlab.com/getlytica/lytica-app/internal/web/controllers/setup"
	"gitlab.com/getlytica/lytica-app/internal/web/controllers/webhooks"
	"gitlab.com/getlytica/lytica-app/internal/web/lib/payments"

	"github.com/urfave/negroni"
)

func (a *App) accountHandlers() {
	acct := account.NewAccount(a.Db, a.SessionStore, a.Logger, payments.NewStripePayments())

	a.Router.Handle("/account/notifications", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.Notifications)),
	))
	a.Router.Handle("/account/notifications/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.NotificationsByDate)),
	))
	a.Router.Handle("/account/subscription", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.Subscription)),
	))
	a.Router.Handle("/account/subscription/subscribe/{planId}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.Subscribe)),
	))
	a.Router.Handle("/account/subscription/cancel", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.CancelSubscription)),
	))
	a.Router.Handle("/account/subscription/change/{planId}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.ChangePlan)),
	))
	a.Router.Handle("/account/subscription/invoices", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.InvoicesByUser)),
	))
	a.Router.Handle("/account/change_password", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(acct.ChangePassword)),
	))
}

func (a *App) apiHandlers() {
	ap := api.NewAPI()

	a.Router.Handle("/api/health_check", negroni.New(
		negroni.Wrap(http.HandlerFunc(ap.HealthCheck)),
	))
}

func (a *App) authHandlers() {
	au := auth.NewAuth(a.Db, a.SessionStore, a.Logger)

	a.Router.HandleFunc("/auth/login", au.Login)
	a.Router.HandleFunc("/auth/logout", au.Logout)
	a.Router.HandleFunc("/auth/callback", au.Callback)
}

func (a *App) cohortsHandlers() {
	c := cohorts.NewCohorts(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/cohorts/high_margin", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.HighMargin)),
	))
	a.Router.Handle("/cohorts/high_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.HighMarginByDate)),
	))
	a.Router.Handle("/cohorts/low_margin", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.LowMargin)),
	))
	a.Router.Handle("/cohorts/low_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.LowMarginByDate)),
	))
	a.Router.Handle("/cohorts/negative_margin", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.NegativeMargin)),
	))
	a.Router.Handle("/cohorts/negative_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(c.NegativeMarginByDate)),
	))
}

func (a *App) dashboardHandlers() {
	dashboard := dashboard.NewDashboard(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(dashboard.Overview)),
	))
	a.Router.Handle("/dashboard/metrics/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(dashboard.MetricsByDate)),
	))
}

func (a *App) expensesHandlers() {
	e := expenses.NewExpenses(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/expenses/cost_of_goods", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(e.CostOfGoods)),
	))
	a.Router.Handle("/expenses/cost_of_goods/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(e.CostOfGoodsByDate)),
	))
	a.Router.Handle("/expenses/cost_of_goods/products", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(e.Products)),
	))
	a.Router.Handle("/expenses/other", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(e.Other)),
	))
	a.Router.Handle("/expenses/other/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(e.OtherByDate)),
	))
	a.Router.Handle("/expenses/other/currencies", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(e.Currencies)),
	))
}

func (a *App) forecastHandlers() {
	f := forecast.NewForecast(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/forecast", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(f.Overview)),
	))
	a.Router.Handle("/forecast/filter/{dateRange}/view/{view}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(f.ForecastByDate)),
	))
}

func (a *App) metricsHandlers() {
	m := metrics.NewMetrics(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/metrics/total_sales", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.TotalSales)),
	))
	a.Router.Handle("/metrics/total_sales/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.TotalSalesByDate)),
	))
	a.Router.Handle("/metrics/units_sold", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.UnitsSold)),
	))
	a.Router.Handle("/metrics/units_sold/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.UnitsSoldByDate)),
	))
	a.Router.Handle("/metrics/amazon_costs", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.AmazonCosts)),
	))
	a.Router.Handle("/metrics/amazon_costs/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.AmazonCostsByDate)),
	))
	a.Router.Handle("/metrics/product_costs", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.ProductCosts)),
	))
	a.Router.Handle("/metrics/product_costs/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.ProductCostsByDate)),
	))
	a.Router.Handle("/metrics/advertising_spend", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.AdvertisingSpend)),
	))
	a.Router.Handle("/metrics/advertising_spend/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.AdvertisingSpendByDate)),
	))
	a.Router.Handle("/metrics/refunds", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.Refunds)),
	))
	a.Router.Handle("/metrics/refunds/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.RefundsByDate)),
	))
	a.Router.Handle("/metrics/shipping_credits", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.ShippingCredits)),
	))
	a.Router.Handle("/metrics/shipping_credits/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.ShippingCreditsByDate)),
	))
	a.Router.Handle("/metrics/promotional_rebates", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.PromotionalRebates)),
	))
	a.Router.Handle("/metrics/promotional_rebates/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.PromotionalRebatesByDate)),
	))
	a.Router.Handle("/metrics/total_costs", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.TotalCosts)),
	))
	a.Router.Handle("/metrics/total_costs/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.TotalCostsByDate)),
	))
	a.Router.Handle("/metrics/gross_margin", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.GrossMargin)),
	))
	a.Router.Handle("/metrics/gross_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.GrossMarginByDate)),
	))
	a.Router.Handle("/metrics/net_margin", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.NetMargin)),
	))
	a.Router.Handle("/metrics/net_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(m.NetMarginByDate)),
	))
}

func (a *App) profitLossHandlers() {
	p := profit_loss.NewProfitLoss(a.Db, a.SessionStore, a.Logger)

	a.Router.Handle("/profit_loss", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(p.Overview)),
	))
	a.Router.Handle("/profit_loss/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.HandlerFunc(a.setupComplete),
		negroni.Wrap(http.HandlerFunc(p.ProfitLossByDate)),
	))
}

func (a *App) setupHandlers() {
	s := setup.NewSetup(a.Db, a.SessionStore, a.Logger, payments.NewStripePayments())

	a.Router.Handle("/setup", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.Subscribe)),
	))
	a.Router.Handle("/setup/subscribe", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.Subscribe)),
	))
	a.Router.Handle("/setup/subscribe/success", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.SubscribeSuccess)),
	))
	a.Router.Handle("/setup/subscribe/cancel", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.SubscribeCancel)),
	))
	a.Router.Handle("/setup/complete", negroni.New(
		negroni.HandlerFunc(a.isAuthenticated),
		negroni.Wrap(http.HandlerFunc(s.Complete)),
	))
}

func (a *App) webhookHandlers() {
	wh := webhooks.NewWebhooks(a.Db, a.Logger, payments.NewStripePayments())

	a.Router.Handle("/webhooks/stripe", negroni.New(
		negroni.Wrap(http.HandlerFunc(wh.Stripe)),
	))
}
