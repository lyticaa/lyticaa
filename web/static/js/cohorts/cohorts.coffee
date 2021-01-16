import $ from 'jquery'
window.jQuery = $
window.$ = $

import AlertsHelper   from '../helpers/alerts'
import CardsHelper    from '../helpers/cards'
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
    this.cards = new CardsHelper()
    this.tables = new TablesHelper()
    this.template = new TemplateHelper()
    this.url = new URLHelper()

  #
  # Initialize.
  #
  init: ->
    this.drawTable()
    this.reload()

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
          $('button.loading').fadeOut(400, ->
            if j.totalSales.chart.sparkline.data
              if j.totalSales.chart.sparkline.data.length > 0
                c.charts.sparkline('cohorts-total-sales-sparkline', [j.totalSales.chart.sparkline])
              else
                $('#cohorts-total-sales-sparkline').html('')

            if j.amazonCosts.chart.sparkline.data
              if j.amazonCosts.chart.sparkline.data.length > 0
                c.charts.sparkline('cohorts-amazon-costs-sparkline', [j.amazonCosts.chart.sparkline])
              else
                $('#cohorts-amazon-costs-sparkline').html('')

            if j.productCosts.chart.sparkline.data
              if j.productCosts.chart.sparkline.data.length > 0
                c.charts.sparkline('cohorts-product-costs-sparkline', [j.productCosts.chart.sparkline])
              else
                $('#cohorts-product-costs-sparkline').html('')

            if j.advertisingSpend.chart.sparkline.data
              if j.advertisingSpend.chart.sparkline.data.length > 0
                c.charts.sparkline('cohorts-advertising-spend-sparkline', [j.advertisingSpend.chart.sparkline])
              else
                $('#cohorts-advertising-spend-sparkline').html('')

            if j.netMargin.chart.sparkline.data
              if j.netMargin.chart.sparkline.data.length > 0
                c.charts.sparkline('cohorts-net-margin-sparkline', [j.netMargin.chart.sparkline])
              else
                $('#cohorts-net-margin-sparkline').html('')

            c.cards.paint('cohorts-total-sales', j.totalSales.total.value, j.totalSales.total.diff)
            c.cards.paint('cohorts-amazon-costs', j.amazonCosts.total.value, j.amazonCosts.total.diff)
            c.cards.paint('cohorts-product-costs', j.productCosts.total.value, j.productCosts.total.diff)
            c.cards.paint('cohorts-advertising-spend', j.advertisingSpend.total.value, j.advertisingSpend.total.diff)
            c.cards.paint('cohorts-net-margin', j.netMargin.total.value, j.netMargin.total.diff)
          )

          if j.data.length > 0
            c.alerts.reset()

          return j.data
        'error': (j) ->
          $('.alert.cohorts-load-error').show()
      'columns': [
        { 'data': 'marketplace' }
        { 'data': 'sku' }
        { 'data': 'description' }
        { 'data': 'totalSales' }
        { 'data': 'quantity' }
        { 'data': 'amazonCosts' }
        { 'data': 'productCosts' }
        { 'data': 'advertisingSpend' }
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

  #
  # Reload.
  #
  reload: ->
    p = this
    p.alerts.reset()

    $('button.reload').on 'click', (e) ->
      p.tables.reload($('.date-filter.active'), $('table'))
      return

    return
