import Turbolinks from 'turbolinks'

import $ from 'jquery'
window.jQuery = $
window.$ = $

import 'bootstrap'
import './vendors/nobleui/nobleui'
import './vendors/intercom/intercom'

import AdminOverview              from './admin/overview'
import AccountNotifications       from './account/notifications'
import AccountPassword            from './account/password'
import AccountSubscription        from './account/subscription'
import Cohorts                    from './cohorts/cohorts'
import DashboardOverview          from './dashboard/overview'
import ExpensesCostOfGoods        from './expenses/cost_of_goods'
import ExpensesOther              from './expenses/other'
import ForecastOverview           from './forecast/overview'
import MetricsAdvertisingSpend    from './metrics/advertising_spend'
import MetricsAmazonCosts         from './metrics/amazon_costs'
import MetricsGrossMargin         from './metrics/gross_margin'
import MetricsNetMargin           from './metrics/net_margin'
import MetricsProductCosts        from './metrics/product_costs'
import MetricsPromotionalRebates  from './metrics/promotional_rebates'
import MetricsRefunds             from './metrics/refunds'
import MetricsShippingCredits     from './metrics/shipping_credits'
import MetricsTotalCosts          from './metrics/total_costs'
import MetricsUnitsSold           from './metrics/units_sold'
import MetricsTotalSales          from './metrics/total_sales'
import ProfitLossOverview         from './profit_loss/overview'
import Reports                    from './reports/reports'
import SetupSubscribe             from './setup/subscribe'
import TemplateHelper             from './helpers/template'

#
# Start.
#
Turbolinks.start()
$(document).on 'ready, turbolinks:load', ->
  init()

  return

#
# Initialize.
#
init = ->
  # Admin.
  admin()

  # Account
  accountNotifications()
  accountSubscription()
  accountResetPassword()

  # Cohorts
  cohortsHighMargin()
  cohortsLowMargin()
  cohortsNegativeMargin()

  # Dashboard
  dashboard()

  # Expenses
  expensesCostOfGoods()
  expensesOther()

  # Forecast
  forecast()

  # Metrics
  metricsAdvertisingSpend()
  metricsAmazonCosts()
  metricsGrossMargin()
  metricsNetMargin()
  metricsProductCosts()
  metricsPromotionalRebates()
  metricsRefunds()
  metricsShippingCredits()
  metricsTotalCosts()
  metricsTotalSales()
  metricsUnitsSold()

  # Profit Loss
  profitLoss()

  # Reports.
  reports()

  # Setup
  setupSubscribe()

  # Icons.
  renderIcons()

  return

#
# Admin.
#
admin = ->
  if $('input.location').data('section') != 'admin'
    return

  a = new AdminOverview()
  a.init()

  return

#
# Account: Notifications
#
accountNotifications = ->
  if $('input.location').data('section') != 'account-notifications'
    return

  a = new AccountNotifications()
  a.init()

  return

#
# Account: Reset Password.
#
accountResetPassword = ->
  $('a.account-change-password ').on 'click', (e)->
    e.preventDefault()

    a = new AccountPassword()
    a.reset()

  return

#
# Account: Subscription.
#
accountSubscription = ->
  if $('input.location').data('section') != 'account-subscription'
    return

  a = new AccountSubscription()
  a.init()

  return

#
# Cohorts: High Margin.
#
cohortsHighMargin = ->
  if $('input.location').data('section') != 'cohorts-high-margin'
    return

  c = new Cohorts()
  c.init()

  return

#
# Cohorts: Low Margin.
#
cohortsLowMargin = ->
  if $('input.location').data('section') != 'cohorts-low-margin'
    return

  c = new Cohorts()
  c.init()

  return

#
# Cohorts: Negative Margin.
#
cohortsNegativeMargin = ->
  if $('input.location').data('section') != 'cohorts-negative-margin'
    return

  c = new Cohorts()
  c.init()

  return

#
# Dashboard.
#
dashboard = ->
  if $('input.location').data('section') != 'dashboard'
    return

  d = new DashboardOverview()
  d.init()

  return

#
# Expenses: Cost of Goods
#
expensesCostOfGoods = ->
  if $('input.location').data('section') != 'expenses-cost-of-goods'
    return

  e = new ExpensesCostOfGoods()
  e.init()

  return

#
# Expenses: Other
#
expensesOther = ->
  if $('input.location').data('section') != 'expenses-other'
    return

  e = new ExpensesOther()
  e.init()

  return

#
# Forecast.
#
forecast = ->
  if $('input.location').data('section') != 'forecast'
    return

  f = new ForecastOverview()
  f.init()

  return

#
# Metrics: Advertising Spend.
#
metricsAdvertisingSpend = ->
  if $('input.location').data('section') != 'metrics-advertising-spend'
    return

  m = new MetricsAdvertisingSpend()
  m.init()

  return

#
# Metrics: Amazon Costs.
#
metricsAmazonCosts = ->
  if $('input.location').data('section') != 'metrics-amazon-costs'
    return

  m = new MetricsAmazonCosts()
  m.init()

  return

#
# Metrics: Gross Margin.
#
metricsGrossMargin = ->
  if $('input.location').data('section') != 'metrics-gross-margin'
    return

  m = new MetricsGrossMargin()
  m.init()

  return

#
# Metrics: Net Margin.
#
metricsNetMargin = ->
  if $('input.location').data('section') != 'metrics-net-margin'
    return

  m = new MetricsNetMargin()
  m.init()

  return

#
# Metrics: Product Costs.
#
metricsProductCosts = ->
  if $('input.location').data('section') != 'metrics-product-costs'
    return

  m = new MetricsProductCosts()
  m.init()

  return

#
# Metrics: Promotional Rebates.
#
metricsPromotionalRebates = ->
  if $('input.location').data('section') != 'metrics-promotional-rebates'
    return

  m = new MetricsPromotionalRebates()
  m.init()

  return

#
# Metrics: Refunds.
#
metricsRefunds = ->
  if $('input.location').data('section') != 'metrics-refunds'
    return

  m = new MetricsRefunds()
  m.init()

  return

#
# Metrics: Shipping Credits.
#
metricsShippingCredits = ->
  if $('input.location').data('section') != 'metrics-shipping-credits'
    return

  m = new MetricsShippingCredits()
  m.init()

  return

#
# Metrics: Total Costs.
#
metricsTotalCosts = ->
  if $('input.location').data('section') != 'metrics-total-costs'
    return

  m = new MetricsTotalCosts()
  m.init()

  return

#
# Metrics: Total Sales.
#
metricsTotalSales = ->
  if $('input.location').data('section') != 'metrics-total-sales'
    return

  m = new MetricsTotalSales()
  m.init()

  return

#
# Metrics: Units Sold.
#
metricsUnitsSold = ->
  if $('input.location').data('section') != 'metrics-units-sold'
    return

  m = new MetricsUnitsSold()
  m.init()

  return

#
# Profit & Loss.
#
profitLoss = ->
  if $('input.location').data('section') != 'profit-loss'
    return

  pl = new ProfitLossOverview()
  pl.init()

  return

#
# Reports.
#
reports = ->
  r = new Reports()
  r.init()

  return

#
# Setup: Subscribe.
#
setupSubscribe = ->
  if $('input.location').data('section') != 'setup-subscribe'
    return

  s = new SetupSubscribe()
  s.init()

  return

#
# Render Icons.
#
renderIcons = ->
  th = new TemplateHelper()
  th.renderIcons()

  return
