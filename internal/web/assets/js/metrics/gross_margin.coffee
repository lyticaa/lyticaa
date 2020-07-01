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
# Gross Margin.
#
export default class MetricsGrossMargin
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
              $('.alert.metrics-gross-margin-chart-error').fadeIn()
            else
              m.charts.line(
                'metrics-gross-margin-chart',
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
          $('.alert.metrics-gross-margin-load-error').show()
      'columns': [
        { 'data': 'marketplace' }
        { 'data': 'sku' }
        { 'data': 'description' }
        { 'data': 'productCosts' }
        { 'data': 'quantity' }
        { 'data': 'totalSales' }
        { 'data': 'amazonCosts' }
        { 'data': 'shippingCredits' }
        { 'data': 'promotionalRebates' }
        { 'data': 'grossMargin' }
        { 'data': 'salesTaxCollected' }
        { 'data': 'totalAmountCollected' }
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
