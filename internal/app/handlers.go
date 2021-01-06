package app

import (
	"net/http"

	"github.com/lyticaa/lyticaa/internal/app/handlers/account"
	"github.com/lyticaa/lyticaa/internal/app/handlers/admin"
	"github.com/lyticaa/lyticaa/internal/app/handlers/api"
	"github.com/lyticaa/lyticaa/internal/app/handlers/auth"
	"github.com/lyticaa/lyticaa/internal/app/handlers/cohorts"
	"github.com/lyticaa/lyticaa/internal/app/handlers/dashboard"
	"github.com/lyticaa/lyticaa/internal/app/handlers/expenses"
	"github.com/lyticaa/lyticaa/internal/app/handlers/forecast"
	"github.com/lyticaa/lyticaa/internal/app/handlers/home"
	"github.com/lyticaa/lyticaa/internal/app/handlers/metrics"
	"github.com/lyticaa/lyticaa/internal/app/handlers/profit_loss"
	"github.com/lyticaa/lyticaa/internal/app/handlers/reports"
	"github.com/lyticaa/lyticaa/internal/app/handlers/webhooks"
	"github.com/lyticaa/lyticaa/internal/app/pkg/accounts/payments"

	"github.com/urfave/negroni"
)

func (a *App) WebHandlers() {
	a.account()
	a.admin()
	a.auth()
	a.cohorts()
	a.dashboard()
	a.expenses()
	a.forecast()
	a.home()
	a.metrics()
	a.profitLoss()
	a.reports()

	a.HTTP.Router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./web/dist"))))
}

func (a *App) APIHandlers() {
	a.api()
	a.webhooks()
}

func (a *App) account() {
	acct := account.NewAccount(a.Database.PG, a.Database.Redis, a.Monitoring.Logger, payments.NewStripePayments())

	a.HTTP.Router.Handle("/account/notifications", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(acct.Notifications)),
	))

	a.HTTP.Router.Handle("/account/notifications/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(acct.NotificationsByDate)),
	))

	a.HTTP.Router.Handle("/account/preferences/setup_completed", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(acct.SetupCompleted)),
	)).Methods("PUT")

	a.HTTP.Router.Handle("/account/preferences/mailing_list", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(acct.MailingList)),
	)).Methods("PUT")

	a.HTTP.Router.Handle("/account/subscription", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(acct.Subscription)),
	))

	a.HTTP.Router.Handle("/account/subscription/new/{result}/{sessionID}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(acct.Subscription)),
	))

	a.HTTP.Router.Handle("/account/subscription/update/{planID}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(acct.UpdateSubscription)),
	)).Methods("PUT")

	a.HTTP.Router.Handle("/account/subscription/cancel", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(acct.CancelSubscription)),
	)).Methods("PUT")

	a.HTTP.Router.Handle("/account/subscription/reactivate/{planID}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(acct.ReactivateSubscription)),
	)).Methods("PUT")

	a.HTTP.Router.Handle("/account/subscription/invoices", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(acct.InvoicesByUser)),
	))

	a.HTTP.Router.Handle("/account/change_password", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(acct.ChangePassword)),
	))
}

func (a *App) admin() {
	ad := admin.NewAdmin(a.Database.PG, a.Database.Redis, a.Monitoring.Logger)

	a.HTTP.Router.Handle("/admin", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.HandlerFunc(a.Admin),
		negroni.Wrap(http.HandlerFunc(ad.Overview)),
	))

	a.HTTP.Router.Handle("/admin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.HandlerFunc(a.Admin),
		negroni.Wrap(http.HandlerFunc(ad.UsersByDate)),
	))

	a.HTTP.Router.Handle("/admin/i/{user}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.HandlerFunc(a.Admin),
		negroni.Wrap(http.HandlerFunc(ad.Impersonate)),
	))

	a.HTTP.Router.Handle("/admin/i/{user}/logout", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.HandlerFunc(a.Admin),
		negroni.Wrap(http.HandlerFunc(ad.LogOut)),
	))
}

func (a *App) api() {
	ap := api.NewAPI()

	a.HTTP.Router.Handle("/api/health_check", negroni.New(
		negroni.Wrap(http.HandlerFunc(ap.HealthCheck)),
	))
}

func (a *App) auth() {
	au := auth.NewAuth(a.Database.PG, a.Database.Redis, a.Monitoring.Logger)

	a.HTTP.Router.HandleFunc("/auth/login", au.Login)
	a.HTTP.Router.HandleFunc("/auth/logout", au.Logout)
	a.HTTP.Router.HandleFunc("/auth/callback", au.Callback)
}

func (a *App) cohorts() {
	c := cohorts.NewCohorts(a.Database.PG, a.Database.Redis, a.Monitoring.Logger)

	a.HTTP.Router.Handle("/cohorts/high_margin", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(c.HighMargin)),
	))

	a.HTTP.Router.Handle("/cohorts/high_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(c.HighMarginByDate)),
	))

	a.HTTP.Router.Handle("/cohorts/low_margin", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(c.LowMargin)),
	))

	a.HTTP.Router.Handle("/cohorts/low_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(c.LowMarginByDate)),
	))

	a.HTTP.Router.Handle("/cohorts/negative_margin", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(c.NegativeMargin)),
	))

	a.HTTP.Router.Handle("/cohorts/negative_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(c.NegativeMarginByDate)),
	))
}

func (a *App) dashboard() {
	dashboard := dashboard.NewDashboard(a.Database.PG, a.Database.Redis, a.Monitoring.Logger)

	a.HTTP.Router.Handle("/dashboard", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.HandlerFunc(a.SetupComplete),
		negroni.Wrap(http.HandlerFunc(dashboard.Overview)),
	))

	a.HTTP.Router.Handle("/dashboard/metrics/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(dashboard.MetricsByDate)),
	))
}

func (a *App) expenses() {
	e := expenses.NewExpenses(a.Database.PG, a.Database.Redis, a.Monitoring.Logger)

	a.HTTP.Router.Handle("/expenses/cost_of_goods", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(e.CostOfGoods)),
	))

	a.HTTP.Router.Handle("/expenses/cost_of_goods/all", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(e.CostOfGoodsByUser)),
	))

	a.HTTP.Router.Handle("/expenses/cost_of_goods/products", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(e.Products)),
	))

	a.HTTP.Router.Handle("/expenses/cost_of_goods/new", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(e.NewCostOfGood)),
	))

	a.HTTP.Router.Handle("/expenses/cost_of_goods/{expense}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(e.EditCostOfGood)),
	)).Methods("PUT")

	a.HTTP.Router.Handle("/expenses/cost_of_goods/{expense}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(e.DeleteCostOfGood)),
	)).Methods("DELETE")

	a.HTTP.Router.Handle("/expenses/other", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(e.Other)),
	))

	a.HTTP.Router.Handle("/expenses/other/all", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(e.OtherByUser)),
	))

	a.HTTP.Router.Handle("/expenses/other/currencies", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(e.Currencies)),
	))

	a.HTTP.Router.Handle("/expenses/other/new", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(e.NewOther)),
	))

	a.HTTP.Router.Handle("/expenses/other/{expense}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(e.EditOther)),
	)).Methods("PUT")

	a.HTTP.Router.Handle("/expenses/other/{expense}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(e.DeleteOther)),
	)).Methods("DELETE")
}

func (a *App) forecast() {
	f := forecast.NewForecast(a.Database.PG, a.Database.Redis, a.Monitoring.Logger)

	a.HTTP.Router.Handle("/forecast/total_sales", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(f.TotalSales)),
	))

	a.HTTP.Router.Handle("/forecast/total_sales/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(f.TotalSalesByDate)),
	))

	a.HTTP.Router.Handle("/forecast/units_sold", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(f.UnitsSold)),
	))

	a.HTTP.Router.Handle("/forecast/units_sold/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(f.UnitsSoldByDate)),
	))
}

func (a *App) home() {
	h := home.NewHome(a.Database.PG, a.Database.Redis, a.Monitoring.Logger)

	a.HTTP.Router.Handle("/", negroni.New(
		negroni.Wrap(http.HandlerFunc(h.Login)),
	))

	a.HTTP.Router.Handle("/onboard", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(h.Onboard)),
	))
}

func (a *App) metrics() {
	m := metrics.NewMetrics(a.Database.PG, a.Database.Redis, a.Monitoring.Logger)

	a.HTTP.Router.Handle("/metrics/total_sales", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.TotalSales)),
	))

	a.HTTP.Router.Handle("/metrics/total_sales/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.TotalSalesByDate)),
	))

	a.HTTP.Router.Handle("/metrics/units_sold", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.UnitsSold)),
	))

	a.HTTP.Router.Handle("/metrics/units_sold/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.UnitsSoldByDate)),
	))

	a.HTTP.Router.Handle("/metrics/amazon_costs", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.AmazonCosts)),
	))

	a.HTTP.Router.Handle("/metrics/amazon_costs/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.AmazonCostsByDate)),
	))

	a.HTTP.Router.Handle("/metrics/product_costs", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.ProductCosts)),
	))

	a.HTTP.Router.Handle("/metrics/product_costs/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.ProductCostsByDate)),
	))

	a.HTTP.Router.Handle("/metrics/advertising_spend", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.AdvertisingSpend)),
	))

	a.HTTP.Router.Handle("/metrics/advertising_spend/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.AdvertisingSpendByDate)),
	))

	a.HTTP.Router.Handle("/metrics/refunds", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.Refunds)),
	))

	a.HTTP.Router.Handle("/metrics/refunds/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.RefundsByDate)),
	))

	a.HTTP.Router.Handle("/metrics/shipping_credits", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.ShippingCredits)),
	))

	a.HTTP.Router.Handle("/metrics/shipping_credits/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.ShippingCreditsByDate)),
	))

	a.HTTP.Router.Handle("/metrics/promotional_rebates", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.PromotionalRebates)),
	))

	a.HTTP.Router.Handle("/metrics/promotional_rebates/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.PromotionalRebatesByDate)),
	))

	a.HTTP.Router.Handle("/metrics/total_costs", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.TotalCosts)),
	))

	a.HTTP.Router.Handle("/metrics/total_costs/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.TotalCostsByDate)),
	))

	a.HTTP.Router.Handle("/metrics/gross_margin", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.GrossMargin)),
	))

	a.HTTP.Router.Handle("/metrics/gross_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.GrossMarginByDate)),
	))

	a.HTTP.Router.Handle("/metrics/net_margin", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.NetMargin)),
	))

	a.HTTP.Router.Handle("/metrics/net_margin/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(m.NetMarginByDate)),
	))
}

func (a *App) profitLoss() {
	p := profit_loss.NewProfitLoss(a.Database.PG, a.Database.Redis, a.Monitoring.Logger)

	a.HTTP.Router.Handle("/profit_loss/statement", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(p.Statement)),
	))

	a.HTTP.Router.Handle("/profit_loss/statement/filter/{dateRange}", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(p.StatementByDate)),
	))
}

func (a *App) reports() {
	r := reports.NewReports(a.Database.PG, a.Database.Redis, a.Monitoring.Logger)

	a.HTTP.Router.Handle("/reports/import", negroni.New(
		negroni.HandlerFunc(a.Authenticated),
		negroni.Wrap(http.HandlerFunc(r.Import)),
	)).Methods("PUT")
}

func (a *App) webhooks() {
	wh := webhooks.NewWebhooks(a.Database.PG, a.Monitoring.Logger, payments.NewStripePayments())

	a.HTTP.Router.Handle("/api/v1/webhooks/stripe", negroni.New(
		negroni.Wrap(http.HandlerFunc(wh.Stripe)),
	))
}
