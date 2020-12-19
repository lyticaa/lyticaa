package helpers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

var (
	cwd, _     = os.Getwd()
	appTpl     = "app"
	appLayouts = []string{
		filepath.Join(cwd, "./web/dist/"+appTpl+".html"),
		filepath.Join(cwd, "./web/templates/partials/_nav.gohtml"),
		filepath.Join(cwd, "./web/templates/partials/_footer.gohtml"),
	}
	bareTpl     = "bare"
	bareLayouts = []string{
		filepath.Join(cwd, "./web/dist/"+bareTpl+".html"),
	}

	PartialsFlash                   = "partials/_flash"
	PartialsAdminImpersonate        = "partials/admin/_impersonate"
	PartialsFiltersFilters          = "partials/filters/_filters"
	PartialsFiltersDate             = "partials/filters/_date"
	PartialsFiltersImport           = "partials/filters/_import"
	PartialsCohortsMargin           = "partials/cohorts/_margin"
	PartialsExpensesCostOfGoodsForm = "partials/expenses/cost_of_goods/_form"
	PartialsExpensesOtherForm       = "partials/expenses/other/_form"

	AccountNotifications      = "account/notifications"
	AccountSubscription       = "account/subscription"
	AdminOverview             = "admin/overview"
	CohortsHighMargin         = "cohorts/high_margin"
	CohortsLowMargin          = "cohorts/low_margin"
	CohortsNegativeMargin     = "cohorts/negative_margin"
	DashboardOverview         = "dashboard/overview"
	ExpensesCostOfGoods       = "expenses/cost_of_goods"
	ExpensesOther             = "expenses/other"
	ForecastTotalSales        = "forecast/total_sales"
	ForecastUnitsSold         = "forecast/units_sold"
	HomeLogin                 = "home/login"
	MetricsAdvertisingSpend   = "metrics/advertising_spend"
	MetricsAmazonCosts        = "metrics/amazon_costs"
	MetricsGrossMargin        = "metrics/gross_margin"
	MetricsNetMargin          = "metrics/net_margin"
	MetricsProductCosts       = "metrics/product_costs"
	MetricsPromotionalRebates = "metrics/promotional_rebates"
	MetricsRefunds            = "metrics/refunds"
	MetricsShippingCredits    = "metrics/shipping_credits"
	MetricsTotalCosts         = "metrics/total_costs"
	MetricsTotalSales         = "metrics/total_sales"
	MetricsUnitsSold          = "metrics/units_sold"
	ProfitLossStatement       = "profit_loss/statement"
	SetupComplete             = "setup/complete"
	SetupSubscribe            = "setup/subscribe"

	DefaultWithImpersonate = append(DefaultNav, []string{
		PartialsAdminImpersonate,
	}...)
	DefaultWithFilters = append(DefaultNav, []string{
		PartialsAdminImpersonate,
		PartialsFiltersFilters,
		PartialsFiltersDate,
		PartialsFiltersImport,
	}...)
	Cohorts = append(DefaultNav, []string{
		PartialsFiltersFilters,
		PartialsFiltersDate,
		PartialsFiltersImport,
		PartialsCohortsMargin,
	}...)
	Expenses = append(DefaultNav, []string{
		PartialsAdminImpersonate,
		PartialsFiltersFilters,
		PartialsFiltersImport,
	}...)
)

const (
	AppLayout  = "app"
	BareLayout = "bare"
)

func TemplateList(page string) []string {
	switch page {
	case AccountNotifications:
		return append(DefaultWithFilters, []string{AccountNotifications}...)
	case AccountSubscription:
		return append(DefaultWithImpersonate, []string{PartialsFlash, AccountSubscription}...)
	case AdminOverview:
		return append(DefaultWithImpersonate, []string{PartialsFiltersFilters, AdminOverview}...)
	case CohortsHighMargin:
		return append(Cohorts, []string{CohortsHighMargin}...)
	case CohortsLowMargin:
		return append(Cohorts, []string{CohortsLowMargin}...)
	case CohortsNegativeMargin:
		return append(Cohorts, []string{CohortsNegativeMargin}...)
	case DashboardOverview:
		return append(DefaultWithFilters, []string{DashboardOverview}...)
	case ExpensesCostOfGoods:
		return append(Expenses, []string{PartialsExpensesCostOfGoodsForm, ExpensesCostOfGoods}...)
	case ExpensesOther:
		return append(Expenses, []string{PartialsExpensesOtherForm, ExpensesOther}...)
	case ForecastTotalSales:
		return append(DefaultWithFilters, []string{ForecastTotalSales}...)
	case ForecastUnitsSold:
		return append(DefaultWithFilters, []string{ForecastUnitsSold}...)
	case HomeLogin:
		return []string{HomeLogin}
	case MetricsAdvertisingSpend:
		return append(DefaultWithFilters, []string{MetricsAdvertisingSpend}...)
	case MetricsAmazonCosts:
		return append(DefaultWithFilters, []string{MetricsAmazonCosts}...)
	case MetricsGrossMargin:
		return append(DefaultWithFilters, []string{MetricsGrossMargin}...)
	case MetricsNetMargin:
		return append(DefaultWithFilters, []string{MetricsNetMargin}...)
	case MetricsProductCosts:
		return append(DefaultWithFilters, []string{MetricsProductCosts}...)
	case MetricsPromotionalRebates:
		return append(DefaultWithFilters, []string{MetricsPromotionalRebates}...)
	case MetricsRefunds:
		return append(DefaultWithFilters, []string{MetricsRefunds}...)
	case MetricsShippingCredits:
		return append(DefaultWithFilters, []string{MetricsShippingCredits}...)
	case MetricsTotalCosts:
		return append(DefaultWithFilters, []string{MetricsTotalCosts}...)
	case MetricsTotalSales:
		return append(DefaultWithFilters, []string{MetricsTotalSales}...)
	case MetricsUnitsSold:
		return append(DefaultWithFilters, []string{MetricsUnitsSold}...)
	case ProfitLossStatement:
		return append(DefaultWithFilters, []string{ProfitLossStatement}...)
	case SetupComplete:
		return append(SetupNav, []string{SetupComplete}...)
	case SetupSubscribe:
		return append(SetupNav, []string{PartialsFlash, SetupSubscribe}...)
	default:
		return DefaultNav
	}
}

func RenderTemplate(w http.ResponseWriter, layout string, templates []string, data interface{}) {
	baseTpl, layouts := layoutFiles(layout)

	files := compileList(layouts, templates)
	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, *baseTpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func layoutFiles(layout string) (*string, []string) {
	switch layout {
	case AppLayout:
		return &appTpl, appLayouts
	case BareLayout:
		return &bareTpl, bareLayouts
	}

	return nil, nil
}

func compileList(layouts, fileList []string) []string {
	var container []string
	container = append(container, layouts...)

	for _, file := range fileList {
		container = append(container, filepath.Join(cwd, "./web/templates/"+file+".gohtml"))
	}

	return container
}
