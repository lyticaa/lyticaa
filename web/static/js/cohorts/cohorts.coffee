import $ from 'jquery'
window.jQuery = $
window.$ = $

import AlertsHelper   from '../helpers/alerts'
import TablesHelper   from '../helpers/tables'
import TemplateHelper from '../helpers/template'
import URLHelper      from '../helpers/url'

require('datatables.net') window, $
require('datatables.net-bs4') window, $

#
# Cohorts: Overview.
#
export default class Cohorts
  constructor: ->
    this.alerts = new AlertsHelper()
    this.tables = new TablesHelper()
    this.template = new TemplateHelper()
    this.url = new URLHelper()

  #
  # Initialize.
  #
  init: ->
    this.drawTable()

    return

  #
  # Draw table.
  #
  drawTable: ->
    c = this

    $('button.loading').fadeIn()

    $('table').DataTable
      'serverSide': true,
      'bFilter': false
      'lengthChange': false
      'ajax':
        'url': c.url.clean() + '/filter/all_time'
        'dataSrc': (j) ->
          $('button.loading').fadeOut()

          if j.data.length > 0
            c.alerts.reset()

          return j.data
        'error': (j) ->
          $('.alert.cohorts-load-error').show()
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
      'preDrawCallback': (settings) ->
        c.tables.preDraw($(this), settings)
      'drawCallback': ->
        c.template.renderIcons()

    $('.date-filter').on 'click', (e) ->
      e.preventDefault()
      c.tables.reload($(this), $('table'))

    this.tables.cleanup($('table'))

    return
