import AlertsHelper     from '../helpers/alerts'
import FiltersHelper    from '../helpers/filters'
import ModalsHelper     from '../helpers/modals'
import TurbolinksHelper from '../helpers/turbolinks'
import URLHelper        from '../helpers/url'

#
# Expenses
#
export default class Expenses
  constructor: ->
    this.alerts = new AlertsHelper()
    this.filters = new FiltersHelper()
    this.modals = new ModalsHelper()
    this.turbolinks = new TurbolinksHelper()
    this.url = new URLHelper()

    this.confirmOpts = 'deleteExpense': 'Are you sure you want to delete this expense?'

  #
  # Expense.
  #
  populate: ->
    ex = this

    $('a.expenses-edit').on 'click', ->
      $('#expenses-edit-modal input#expense').data('expense', $(this).data('expense'))

      if $('#expenses-edit-modal select#product').length
        $('#expenses-edit-modal select#product').append(
          $('<option/>').val($(this).data('product')).text("#{$(this).data('sku')} - #{$(this).data('marketplace')}")
        ).attr('disabled', 'disabled')

      $('#expenses-edit-modal input#description').val($(this).data('description'))

      if $('#expenses-edit-modal input#from-date').length
        $('#expenses-edit-modal input#from-date').val($(this).data('from-date'))

      if $('#expenses-edit-modal input#date-time').length
        $('#expenses-edit-modal input#date-time').val($(this).data('date-time'))

      $('#expenses-edit-modal input#amount').val($(this).data('amount'))

      if $('#expenses-edit-modal select#currency').length
        ex.loadCurrencies($(this).data('currency'))

    return

  #
  # Validation errors.
  #
  validationErrors: (response) ->
    if response.description
      $('.modal:visible .alert.expenses-description-validation-error').fadeIn()
      $('.modal:visible input#description').addClass('is-invalid')

    if response.fromdate
      $('.modal:visible .alert.expenses-from-date-validation-error').fadeIn()
      $('.modal:visible input#from-date').addClass('is-invalid')

    if response.datetime
      $('.modal:visible .alert.expenses-date-time-validation-error').fadeIn()
      $('.modal:visible input#from-date').addClass('is-invalid')

    if response.amount
      $('.modal:visible .alert.expenses-amount-validation-error').fadeIn()
      $('.modal:visible input#amount').addClass('is-invalid')

    return

  #
  # Currencies.
  #
  loadCurrencies: (selected) ->
    ex = this

    $('select#currency').html('')

    $.ajax(
      type: 'GET'
      url: ex.url.clean() + '/currencies'
      timeout: 10000
      statusCode:
        200: (j) ->
          $.each j, ->
            content = $('<option/>').val(@currencyId).text("#{@code} (#{@symbol})")
            if selected != undefined && @currencyId == selected
              content.attr('selected', 'selected')

            $('select#currency').append content
          return
    ).fail ->
      $('.alert.expenses-other-currencies-load-error').fadeIn()

    return

  #
  # Products.
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
            $('.alert.expenses-products-empty').fadeIn(400, ->
              $('select#product').attr('disabled', 'disabled')
              ex.modals.disable('Submit')
            )
    ).fail ->
      $('.alert.expenses-products-load-error').fadeIn()

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
  # New expense.
  #
  new: (products, currencies)->
    ex = this

    ex.alerts.resetSuccess()
    ex.alerts.resetErrors()

    ex.modals.resetForm()
    ex.modals.reset('Submit')

    if products
      ex.loadProducts()

    if currencies
      ex.loadCurrencies()

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
            $('.alert.expenses-error').fadeIn(400, ->
              ex.stop()
            )
          422: (j) ->
            ex.stop()
            ex.validationErrors(j.responseJSON)
          500: ->
            $('.alert.expenses-error').fadeIn(400, ->
              ex.stop()
            )
      )

    return

  #
  # Edit expense.
  #
  edit: ->
    ex = this

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
            $('.alert.expenses-error').fadeIn(400, ->
              ex.stop()
            )
          422: (j) ->
            ex.stop()
            ex.validationErrors(j.responseJSON)
          500: ->
            $('.alert.expenses-error').fadeIn(400, ->
              ex.stop()
            )
      )

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
