import $ from 'jquery'
window.jQuery = $
window.$ = $

import AlertsHelper     from '../helpers/alerts'
import FiltersHelper    from '../helpers/filters'
import ModalsHelper     from '../helpers/modals'
import TablesHelper     from '../helpers/tables'
import TemplateHelper   from '../helpers/template'
import TurbolinksHelper from '../helpers/turbolinks'
import URLHelper        from '../helpers/url'

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
    this.turbolinks = new TurbolinksHelper()
    this.url = new URLHelper()

    this.confirmOpts = 'deleteExpense': 'Are you sure you want to delete this expense?'

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
        'url': ex.url.clean() + '/all'
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
        {
          'data': null
          'className': 'text-center pr-0 pl-0 w-15'
          'fnCreatedCell': (nTd, sData, oData, iRow) ->
            content = """
                <a href='#' class='expenses-cost-of-goods-edit' data-toggle='modal' data-target='#expenses-cost-of-goods-edit-modal' data-expense='#{oData.DT_RowId}' data-product='#{oData.productId}' data-sku='#{oData.sku}' data-marketplace='#{oData.marketplace}' data-description='#{oData.description}' data-from-date='#{oData.fromDate}' data-amount='#{oData.amount}' target='_blank'>
                  <i class='edit' data-feather='edit' data-toggle='tooltip' data-placement='top' title='Edit the expense.'></i>
                </a>
                <i class='delete' data-expense='#{oData.DT_RowId}' data-idx='#{iRow}' data-feather='trash' data-toggle='tooltip' data-placement='top' title='Delete the expense.'></i>
              """
            $(nTd).html content
            return
        }
      ]
      'language': {
        'infoFiltered': ''
      }
      'preDrawCallback': (settings) ->
        ex.tables.preDraw($(this), settings)
      'drawCallback': ->
        ex.template.renderIcons()

        ex.expense()
        ex.edit()
        ex.delete()

    $('.date-filter').on 'click', (e) ->
      e.preventDefault()
      ex.tables.reload($(this), $('table'))

    this.tables.cleanup($('table'))

    return

  #
  # Expense.
  #
  expense: ->
    $('a.expenses-cost-of-goods-edit').on 'click', ->
      $('#expenses-cost-of-goods-edit-modal input#expense').data('expense', $(this).data('expense'))
      $('#expenses-cost-of-goods-edit-modal select#product').append(
        $('<option/>').val($(this).data('product')).text("#{$(this).data('sku')} - #{$(this).data('marketplace')}")
      ).attr('disabled', 'disabled')
      $('#expenses-cost-of-goods-edit-modal input#description').val($(this).data('description'))
      $('#expenses-cost-of-goods-edit-modal input#from-date').val($(this).data('from-date'))
      $('#expenses-cost-of-goods-edit-modal input#amount').val($(this).data('amount'))

    return

  #
  # Delete expense.
  #
  delete: ->
    ex = this

    $('.delete').on 'click', ->
      r = confirm ex.confirmOpts.deleteExpense

      if r == true
        ex.alerts.reset()
        ex.turbolinks.start()

        $('button.loading').fadeIn()

        $.ajax(
          type: 'DELETE'
          url: ex.url.clean() + '/' + $(this).data('expense')
          timeout: 10000
          statusCode:
            200: (j) ->
              $('button.loading').fadeOut(400, ->
                ex.turbolinks.stop()
                $('table').DataTable().ajax.reload()
              )
            400: ->
              $('button.loading').fadeOut(400, ->
                ex.turbolinks.stop()
              )
            500: ->
              $('button.loading').fadeOut(400, ->
                ex.turbolinks.stop()
              )
        )

    return

  #
  # Validation errors.
  #
  validationErrors: (response) ->
    if response.description
      $('.modal:visible .alert.expenses-cost-of-goods-description-validation-error').fadeIn()
      $('.modal:visible input#description').addClass('is-invalid')

    if response.fromdate
      $('.modal:visible .alert.expenses-cost-of-goods-from-date-validation-error').fadeIn()
      $('.modal:visible input#from-date').addClass('is-invalid')

    if response.amount
      $('.modal:visible .alert.expenses-cost-of-goods-amount-validation-error').fadeIn()
      $('.modal:visible input#amount').addClass('is-invalid')

    return

  #
  # New expense.
  #
  new: ->
    ex = this

    $('#expenses-cost-of-goods-modal').on 'shown.bs.modal', ->
      ex.alerts.resetSuccess()
      ex.alerts.resetErrors()

      ex.modals.resetForm()
      ex.modals.reset('Submit')

      ex.loadProducts()

      ex.filters.datePicker('.modal:visible .datepicker')

      $('.modal:visible form').on 'submit', (e) ->
        e.preventDefault()
        ex.start()

        $.ajax(
          type: 'POST'
          url: ex.url.clean() + '/new'
          timeout: 10000
          data: $('.modal:visible form').serialize()
          statusCode:
            200: ->
              ex.stop()

              $('.modal:visible button.close').trigger('click')
              $('table').DataTable().ajax.reload()
            400: ->
              $('.alert.expenses-cost-of-goods-error').fadeIn(400, ->
                ex.stop()
              )
            422: (j) ->
              ex.stop()
              ex.validationErrors(j.responseJSON)
            500: ->
              $('.alert.expenses-cost-of-goods-error').fadeIn(400, ->
                ex.stop()
              )
        )

        return

    return

  #
  # Edit expense.
  #
  edit: ->
    ex = this

    $('#expenses-cost-of-goods-edit-modal').on 'shown.bs.modal', ->
      ex.alerts.resetSuccess()
      ex.alerts.resetErrors()

      ex.modals.reset('Submit')

      ex.filters.datePicker('.modal:visible .datepicker')

      $('.modal:visible form').on 'submit', (e) ->
        e.preventDefault()
        ex.start()

        $.ajax(
          type: 'PUT'
          url: ex.url.clean() + '/' + $('.modal:visible input#expense').data('expense')
          timeout: 10000
          data: $('.modal:visible form').serialize()
          statusCode:
            200: ->
              ex.stop()

              $('.modal:visible button.close').trigger('click')
              $('table').DataTable().ajax.reload()
            400: ->
              $('.alert.expenses-cost-of-goods-error').fadeIn(400, ->
                ex.stop()
              )
            422: (j) ->
              ex.stop()
              ex.validationErrors(j.responseJSON)
            500: ->
              $('.alert.expenses-cost-of-goods-error').fadeIn(400, ->
                ex.stop()
              )
        )

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
              $('select#product').append $('<option/>').val(@productId).text("#{@sku} - #{@marketplace} - #{@description}")
            return
          else
            $('.alert.expenses-cost-of-goods-products-empty').fadeIn(400, ->
              $('select#product').attr('disabled', 'disabled')
              ex.modals.disable('Add')
              # $('button[type="submit"]').attr('disabled', 'disabled')
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
    this.modals.reset('Submit')

    return
