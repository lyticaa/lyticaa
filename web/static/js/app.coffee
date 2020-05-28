import Turbolinks from 'turbolinks'

import $ from 'jquery'
window.jQuery = $
window.$ = $

import 'js-cookie'
import 'dropzone'

require('datatables.net') window, $
require('datatables.net-bs4') window, $

#
# Initialize.
#
initialize = ->
  user()
  dashboard()

  # Metrics
  metricsTotalSales()
  metricsUnitsSold()
  metricsAmazonCosts()
  metricsAdvertisingSpend()
  metricsRefunds()
  metricsShippingCredits()
  metricsPromotionalRebates()
  metricsTotalCosts()
  metricsNetMargin()

  # Cohorts
  cohortsHighMargin()
  cohortsLowMargin()
  cohortsNegativeMargin()

  # Forecast
  forecast()

  # Expenses
  expensesCostOfGoods()
  expensesOther()

  # Profit Loss
  profitLoss()

  # Account
  accountNotifications()

  support()
  uploads()
  payments()
  return

#
# User.
#
user = ->
  imageUrl = $('.profile-image').attr('rel')
  $('.profile-image-thumb, .profile-image').attr 'src', imageUrl

  $('.log-out').click (e) ->
    Cookies.remove 'auth-session'
    return
  return

#
# Dashboard.
#
dashboard = ->
  if $('input.location').data('section') != 'dashboard'
    return

  loadDashboard()

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    $('.date-filter.active').removeClass 'active'
    $(this).addClass 'active'

    loadDashboard()

    return
  return

#
# Load Summary.
#
loadDashboard = ->
  tbStart()
  resetErrors()

  $('button.loading').show()

  $.ajax
    type: 'GET'
    url: window.location.href + 'dashboard/metrics/filter/' + $('.date-filter.active').data('range')
    timeout: 10000
    statusCode:
      200: (data) ->
        tbStop()
        $('button.loading').hide()

        #feather.replace()
      500: ->
        tbStop()
        $('button.loading').hide()
        $('.alert.alert-danger.metrics').show()
    error: ->
      tbStop()
      $('button.loading').hide()
      $('.alert.alert-danger.metrics').show()
  return

#
# Metrics: Total Sales.
#
metricsTotalSales = ->
  if $('input.location').data('section') != 'metrics-total-sales'
    return

  $('button.loading').show()

  $('#metrics-total-sales-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.metrics-error').show()
    'columns': [
      { 'data': 'date' }
      { 'data': 'sku' }
      { 'data': 'asin' }
      { 'data': 'productName' }
      { 'data': 'sales' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#metrics-total-sales-table'))

  dtCleanup($('#metrics-total-sales-table'))

#
# Metrics: Units Sold.
#
metricsUnitsSold = ->
  if $('input.location').data('section') != 'metrics-units-sold'
    return

  $('button.loading').show()

  $('#metrics-units-sold-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.metrics-error').show()
    'columns': [
      { 'data': 'date' }
      { 'data': 'sku' }
      { 'data': 'asin' }
      { 'data': 'productName' }
      { 'data': 'grossQuantitySold' }
      { 'data': 'netQuantitySold' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#metrics-units-sold-table'))

  dtCleanup($('#metrics-units-sold-table'))

#
# Metrics: Amazon Costs.
#
metricsAmazonCosts = ->
  if $('input.location').data('section') != 'metrics-amazon-costs'
    return

  $('button.loading').show()

  $('#metrics-amazon-costs-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.metrics-error').show()
    'columns': [
      { 'data': 'date' }
      { 'data': 'sku' }
      { 'data': 'asin' }
      { 'data': 'productName' }
      { 'data': 'type' }
      { 'data': 'amazonCosts' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#metrics-amazon-costs-table'))

  dtCleanup($('#metrics-amazon-costs-table'))

#
# Metrics: Advertising Spend.
#
metricsAdvertisingSpend = ->
  if $('input.location').data('section') != 'metrics-advertising-spend'
    return

  $('button.loading').show()

  $('#metrics-advertising-spend-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.metrics-error').show()
    'columns': [
      { 'data': 'date' }
      { 'data': 'sku' }
      { 'data': 'asin' }
      { 'data': 'productName' }
      { 'data': 'advertisingSpend' }
      { 'data': 'percentageOfSales' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#metrics-advertising-spend-table'))

  dtCleanup($('#metrics-advertising-spend-table'))

#
# Metrics: Refunds.
#
metricsRefunds = ->
  if $('input.location').data('section') != 'metrics-refunds'
    return

  $('button.loading').show()

  $('#metrics-refunds-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.metrics-error').show()
    'columns': [
      { 'data': 'date' }
      { 'data': 'sku' }
      { 'data': 'asin' }
      { 'data': 'productName' }
      { 'data': 'refunds' }
      { 'data': 'percentageOfSales' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#metrics-refunds-table'))

  dtCleanup($('#metrics-refunds-table'))

#
# Metrics: Shipping Credits.
#
metricsShippingCredits = ->
  if $('input.location').data('section') != 'metrics-shipping-credits'
    return

  $('button.loading').show()

  $('#metrics-shipping-credits-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.metrics-error').show()
    'columns': [
      { 'data': 'date' }
      { 'data': 'sku' }
      { 'data': 'asin' }
      { 'data': 'productName' }
      { 'data': 'shippingCredits' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#metrics-shipping-credits-table'))

  dtCleanup($('#metrics-shipping-credits-table'))

#
# Metrics: Promotional Rebates.
#
metricsPromotionalRebates = ->
  if $('input.location').data('section') != 'metrics-promotional-rebates'
    return

  $('button.loading').show()

  $('#metrics-promotional-rebates-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.metrics-error').show()
    'columns': [
      { 'data': 'date' }
      { 'data': 'sku' }
      { 'data': 'asin' }
      { 'data': 'productName' }
      { 'data': 'costOfCoupons' }
      { 'data': 'quantity' }
      { 'data': 'promotionalRebates' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#metrics-promotional-rebates-table'))

  dtCleanup($('#metrics-promotional-rebates-table'))

#
# Metrics: Total Costs.
#
metricsTotalCosts = ->
  if $('input.location').data('section') != 'metrics-total-costs'
    return

  $('button.loading').show()

  $('#metrics-total-costs-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.metrics-error').show()
    'columns': [
      { 'data': 'date' }
      { 'data': 'sku' }
      { 'data': 'asin' }
      { 'data': 'productName' }
      { 'data': 'amazonCosts' }
      { 'data': 'productCosts' }
      { 'data': 'productCostPerUnit' }
      { 'data': 'totalCosts' }
      { 'data': 'percentage' }
      { 'data': 'percentageOfSales' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#metrics-total-costs-table'))

  dtCleanup($('#metrics-total-costs-table'))

#
# Metrics: Net Margin.
#
metricsNetMargin = ->
  if $('input.location').data('section') != 'metrics-net-margin'
    return

  $('button.loading').show()

  $('#metrics-net-margin-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.metrics-error').show()
    'columns': [
      { 'data': 'date' }
      { 'data': 'sku' }
      { 'data': 'asin' }
      { 'data': 'productName' }
      { 'data': 'netMargin' }
      { 'data': 'percentage' }
      { 'data': 'netMarginPerUnit' }
      { 'data': 'roi' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#metrics-net-margin-table'))

  dtCleanup($('#metrics-net-margin-table'))

#
# Cohorts: High Margin.
#
cohortsHighMargin = ->
  if $('input.location').data('section') != 'cohorts-high-margin'
    return

  loadCohort()

#
# Cohorts: Low Margin.
#
cohortsLowMargin = ->
  if $('input.location').data('section') != 'cohorts-low-margin'
    return

  loadCohort()

#
# Cohorts: Negative Margin.
#
cohortsNegativeMargin = ->
  if $('input.location').data('section') != 'cohorts-negative-margin'
    return

  loadCohort()

#
# Load Cohorts Data.
#
loadCohort = ->
  $('button.loading').show()

  $('#cohorts-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.metrics-error').show()
    'columns': [
      { 'data': 'sku' }
      { 'data': 'asin' }
      { 'data': 'productName' }
      { 'data': 'totalSales' }
      { 'data': 'grossQuantitySold' }
      { 'data': 'netQuantitySold' }
      { 'data': 'amazonCosts' }
      { 'data': 'productCosts' }
      { 'data': 'costOfCoupons' }
      { 'data': 'advertisingSpend' }
      { 'data': 'coupons' }
      { 'data': 'refunds' }
      { 'data': 'shippingCredits' }
      { 'data': 'promotionalRebates' }
      { 'data': 'totalCosts' }
      { 'data': 'netMargin' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#cohorts-table'))

  dtCleanup($('#cohorts-table'))

  return

#
# Forecast.
#
forecast = ->
  if $('input.location').data('section') != 'forecast'
    return

  loadForecast($('.date-filter.active'))

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    loadForecast($(this))

  return

#
# Load Forecast.
#
loadForecast = (obj) ->
  resetErrors()
  resetWarnings()
  tbStart()

  $('.date-filter.active').removeClass 'active'
  $(obj).addClass 'active'

  $.ajax
    type: 'POST'
    url: window.location.href + '/filter/' + $(obj).data('range')
    statusCode:
      200: ->
        tbStop()
        $('button.loading').hide()
      500: ->
        tbStop()
        $('.alert.alert-danger.forecast').show()
        $('button.loading').hide()

  return

#
# Expenses: Cost of Goods
#
expensesCostOfGoods = ->
  if $('input.location').data('section') != 'expenses-cost-of-goods'
    return

  $('button.loading').show()

  $('#expenses-cost-of-goods-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.expenses-error').show()
    'columns': [
      { 'data': 'sku' }
      { 'data': 'asin' }
      { 'data': 'productName' }
      { 'data': 'description' }
      { 'data': 'startDate' }
      { 'data': 'endDate' }
      { 'data': 'type' }
      { 'data': 'cost' }
      { 'data': 'currency' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#expenses-cost-of-goods-table'))

  dtCleanup($('#expenses-cost-of-goods-table'))

#
# Expenses: Other
#
expensesOther = ->
  if $('input.location').data('section') != 'expenses-other'
    return

  $('button.loading').show()

  $('#expenses-other-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.expenses-error').show()
    'columns': [
      { 'data': 'description' }
      { 'data': 'startDate' }
      { 'data': 'endDate' }
      { 'data': 'type' }
      { 'data': 'cost' }
      { 'data': 'currency' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#expenses-other-table'))

  dtCleanup($('#expenses-other-table'))

#
# Profit & Loss.
#
profitLoss = ->
  if $('input.location').data('section') != 'profit-loss'
    return

  $('button.loading').show()

  $('#profit-loss-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.profit-loss-error').show()
    'columns': [
      { 'data': 'particulars' }
      { 'data': 'amount' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#profit-loss-table'))

  dtCleanup($('#profit-loss-table'))

#
# Account: Notifications
#
accountNotifications = ->
  if $('input.location').data('section') != 'account-notifications'
    return

  $('button.loading').show()

  $('#account-notifications-table').DataTable
    'serverSide': true,
    'bFilter': false
    'lengthChange': false
    'ajax':
      'url': window.location.href + '/filter/today'
      'dataSrc': (j) ->
        $('button.loading').hide()

        if j.data.length > 0
          resetErrors()
          resetWarnings()

        return j.data
      'error': (j) ->
        $('.alert.alert-error.account-error').show()
    'columns': [
      { 'data': 'notification' }
      { 'data': 'date' }
    ]
    'language': {
      'infoFiltered': ''
    }
    preDrawCallback: (settings) ->
      dtPreDrawCallback(this, settings)

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    dtReload(this, $('#account-notifications-table'))

  dtCleanup($('#account-notifications-table'))

#
# Support (intercom).
#
support = ->
  intercomId = $('.intercom').data('intercom-id')

  window.intercomSettings =
    app_id: intercomId
    name: $('.intercom').data('name')
    email: $('.intercom').data('email')
    created_at: $('.intercom').data('created-at')

  do ->
    w = window
    ic = w.Intercom
    if typeof ic == 'function'
      ic 'reattach_activator'
      ic 'update', w.intercomSettings
    else
      d = document

      i = ->
        i.c arguments
        return

      i.q = []

      i.c = (args) ->
        i.q.push args
        return

      w.Intercom = i

      l = ->
        s = d.createElement('script')
        s.type = 'text/javascript'
        s.async = true
        s.src = 'https://widget.intercom.io/widget/' + intercomId
        x = d.getElementsByTagName('script')[0]
        x.parentNode.insertBefore s, x
        return

      if w.attachEvent
        w.attachEvent 'onload', l
      else
        w.addEventListener 'load', l, false

    return

#
# Uploads.
#
uploads = ->
  if $('#dropzone').length > 0
    $('#dropzone').dropzone

  return

#
# Payments.
#
payments = ->
  if $('a.stripe').length > 0
    stripe = Stripe($('.stripe-pk').data('stripe-pk'))
    $('a.stripe').on 'click', ->
      stripe.redirectToCheckout(sessionId: $(this).attr('rel')).then (result) ->
        alert result.error.message
      return

  return

#
# Datatable Pre-draw callback.
#
dtPreDrawCallback = (obj, settings) ->
  api = new ($.fn.dataTable.Api)(settings)
  pagination = $(obj).closest('.dataTables_wrapper').find('.dataTables_paginate')
  pagination.toggle api.page.info().pages > 1

  return

#
# Datatable Cleanup.
#
dtCleanup = (table) ->
  table.each ->
    datatable = $(this)
    length_sel = datatable.closest('.dataTables_wrapper').find('div[id$=_length] select')
    length_sel.removeClass 'form-control-sm'
    return

  return

#
# Datatable Reload.
#
dtReload = (obj, table) ->
  $('button.loading').show()

  $('.date-filter.active').removeClass 'active'
  $(obj).addClass 'active'
  table.DataTable().ajax.url(window.location.href + '/filter/' + $(obj).data('range')).load()

  return

#
# Reset errors.
#
resetErrors = ->
  $('.alert.alert-danger').each ->
    if $(this).is(':visible')
      $(this).hide()
    return

  return

#
# Reset warnings.
#
resetWarnings = ->
  $('.alert.alert-warning').each ->
    if $(this).is(':visible')
      $(this).hide()
    return

  return
#
# Start Turbolinks progress bar.
#
tbStart = ->
  Turbolinks.controller.adapter.progressBar.setValue 0
  Turbolinks.controller.adapter.progressBar.show()

  return

#
# Stop Turbolinks progress bar.
#
tbStop = ->
  Turbolinks.controller.adapter.progressBar.hide()

  return

#
# Init
#
Turbolinks.start()
$(document).on 'ready, turbolinks:load', ->
  initialize()

  return
