import $ from 'jquery'
window.jQuery = $
window.$ = $

import AlertsHelper   from '../helpers/alerts'
import FiltersHelper  from '../helpers/filters'
import ModalsHelper   from '../helpers/modals'
import TablesHelper   from '../helpers/tables'
import TemplateHelper from '../helpers/template'
import URLHelper      from '../helpers/url'

require('datatables.net') window, $
require('datatables.net-bs4') window, $

#
# Expenses.
#
export default class ExpensesCostOfGoods
  constructor: ->
    this.alerts = new AlertsHelper()
    this.filters = new FiltersHelper()
    this.modals = new ModalsHelper()
    this.tables = new TablesHelper()
    this.template = new TemplateHelper()
    this.url = new URLHelper()

  #
  # Initialize.
  #
  init: ->
    this.drawTable()
    this.new()

    return

  #
  # Draw table.
  #
  drawTable: ->
    ex = this

    $('button.loading').fadeIn()

    $('table').DataTable
      'serverSide': true,
      'bFilter': false
      'lengthChange': false
      'ajax':
        'url': ex.url.clean() + '/filter/all_time'
        'dataSrc': (j) ->
          $('button.loading').fadeOut()

          if j.data.length > 0
            ex.alerts.reset()

          return j.data
        'error': (j) ->
          $('.alert.expenses-cost-of-goods-load-error').show()
      'columns': [
        { 'data': 'marketplace' }
        { 'data': 'sku' }
        { 'data': 'description' }
        { 'data': 'fromDate' }
        { 'data': 'amount' }
      ]
      'language': {
        'infoFiltered': ''
      }
      'preDrawCallback': (settings) ->
        ex.tables.preDraw($(this), settings)
      'drawCallback': ->
        ex.template.renderIcons()

    $('.date-filter').on 'click', (e) ->
      e.preventDefault()
      ex.tables.reload($(this), $('table'))

    this.tables.cleanup($('table'))

    return

  #
  # New cost of good.
  #
  new: ->
    ex = this

    $('#expenses-cost-of-goods-modal').on 'shown.bs.modal', ->
      ex.alerts.resetSuccess()
      ex.alerts.resetErrors()
      ex.modals.resetForm()

      ex.loadProducts()

      ex.filters.datePicker('#expenses-cost-of-goods-modal .datepicker')

      $('form#expenses-cost-of-goods').on 'submit', (e) ->
        e.preventDefault()

        return

    return

  #
  # Currencies.
  #
  loadProducts: ->
    ex = this

    $('select#product').removeAttr('disabled').html('')

    $.ajax(
      type: 'GET'
      url: ex.url.clean() + '/products'
      timeout: 10000
      statusCode:
        200: (j) ->
          if j.length > 0
            $.each j, ->
              $('select#product').append $('<option/>').val(@productId).text("#{@sku} - #{@marketplace} - #{@descriptoion}")
            return
          else
            $('.alert.expenses-cost-of-goods-products-empty').fadeIn(400, ->
              $('select#product').attr('disabled', 'disabled')
            )
    ).fail ->
      $('.alert.expenses-other-currencies-load-error').fadeIn()

    return

  #
  # Start.
  #
  start: (text) ->
    this.alerts.resetErrors()
    this.turbolinks.start()
    this.modals.disable(text)

    return

  #
  # Stop.
  #
  stop: ->
    this.turbolinks.stop()
    this.modals.reset('Add')

    return
