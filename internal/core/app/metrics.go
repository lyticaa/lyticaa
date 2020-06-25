package app

import (
	"net/http"

	"gitlab.com/getlytica/lytica-app/internal/core/app/metrics"

	"github.com/urfave/negroni"
)

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
