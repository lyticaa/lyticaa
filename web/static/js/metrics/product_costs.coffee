import $ from 'jquery'
window.jQuery = $
window.$ = $

import Metrics        from './metrics'
import AlertsHelper   from '../helpers/alerts'
import ChartsHelper   from '../helpers/charts'
import TablesHelper   from '../helpers/tables'
import TemplateHelper from '../helpers/template'
import URLHelper      from '../helpers/url'

require('datatables.net') window, $
require('datatables.net-bs4') window, $

#
# Product Costs.
#
export default class MetricsProductCosts
  constructor: ->
    this.metrics = new Metrics()
    this.alerts = new AlertsHelper()
    this.charts = new ChartsHelper()
    this.tables = new TablesHelper()
    this.template = new TemplateHelper()
    this.url = new URLHelper()

  #
  # Initialize.
  #
  init: ->
    this.drawTable()
    this.metrics.reload()

    return

  #
  # Draw table.
  #
  drawTable: ->
    m = this

    $('button.loading').fadeIn()

    $('table').DataTable
      'serverSide': true,
      'bFilter': false
      'lengthChange': false
      'ajax':
        'url': m.url.clean() + '/filter/all_time'
        'dataSrc': (j) ->
          $('button.loading').fadeOut(400, ->
            if j.chart.line.categories[0].category.length == 0
              $('.alert.metrics-product-costs-chart-error').fadeIn()
            else
              m.charts.line(
                'metrics-product-costs-chart',
                'AMOUNT',
                'DATE',
                j.chart.line.categories,
                j.chart.line.dataSets
              )
          )

          if j.data.length > 0
            m.alerts.reset()

          return j.data
        'error': (j) ->
          $('.alert.metrics-product-costs-load-error').show()
      'columns': [
        { 'data': 'sku' }
        { 'data': 'cost' }
        { 'data': 'quantitySold' }
        { 'data': 'quantitySoldCoupons' }
        { 'data': 'netQuantitySold' }
        { 'data': 'totalCostSold' }
        { 'data': 'totalCostCoupons' }
        { 'data': 'advertisingCosts' }
        { 'data': 'refunds' }
        { 'data': 'totalCosts' }
      ]
      'language': {
        'infoFiltered': ''
      }
      'preDrawCallback': (settings) ->
        m.tables.preDraw($(this), settings)
      'drawCallback': ->
        m.template.renderIcons()

    $('.date-filter').on 'click', (e) ->
      e.preventDefault()
      m.tables.reload($(this), $('table'))

    this.tables.cleanup($('table'))

    return
