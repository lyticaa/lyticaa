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
# Cohorts.
#
export default class MetricsAmazonCosts
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
    m = this

    $('button.loading').fadeIn()

    $('table').DataTable
      'serverSide': true,
      'bFilter': false
      'lengthChange': false
      'ajax':
        'url': m.url.clean() + '/filter/all_time'
        'dataSrc': (j) ->
          $('button.loading').fadeOut()

          if j.data.length > 0
            m.alerts.reset()

          return j.data
        'error': (j) ->
          $('.alert.metrics-amazon-costs-load-error').show()
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
      'preDrawCallback': (settings) ->
        m.tables.preDraw($(this), settings)
      'drawCallback': ->
        m.template.renderIcons()

    $('.date-filter').on 'click', (e) ->
      e.preventDefault()
      m.tables.reload($(this), $('table'))

    this.tables.cleanup($('table'))

    return
