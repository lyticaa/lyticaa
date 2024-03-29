import $ from 'jquery'
window.jQuery = $
window.$ = $

import AlertsHelper     from '../helpers/alerts'
import ModalsHelper     from '../helpers/modals'
import TablesHelper     from '../helpers/tables'
import TemplateHelper   from '../helpers/template'
import TurbolinksHelper from '../helpers/turbolinks'
import URLHelper        from '../helpers/url'

require('datatables.net') window, $
require('datatables.net-bs4') window, $

#
# Account: Subscription.
#
export default class AccountSubscription
  constructor: ->
    this.alerts = new AlertsHelper()
    this.modals = new ModalsHelper()
    this.tables = new TablesHelper()
    this.template = new TemplateHelper()
    this.turbolinks = new TurbolinksHelper()
    this.url = new URLHelper()

  #
  # Initialize.
  #
  init: ->
    this.drawTable()
    this.subscribe()
    this.reactivate()
    this.change()
    this.cancel()

    return

  #
  # Draw table.
  #
  drawTable: ->
    s = this

    $('button.loading').fadeIn()

    $('table#account-subscription-invoice-table').DataTable
      'serverSide': true,
      'bFilter': false
      'ordering': false
      'lengthChange': false
      'paging': false
      'initComplete': ->
        s.template.renderIcons()
      'ajax':
        'url': s.url.clean() + '/invoices'
        'dataSrc': (j) ->
          $('button.loading').fadeOut()

          if j.data.length > 0
            s.alerts.reset()

          return j.data
        'error': ->
          $('.alert.account-subscription-invoices-load-error').fadeIn()
      'columns': [
        { 'data': 'number' }
        { 'data': 'date' }
        { 'data': 'amount' }
        {
          'data': 'status'
          'fnCreatedCell': (nTd, sData, oData, iRow, iCol) ->
            content = """
              <span class='badge #{oData.statusClass} text-uppercase'>#{oData.status}</span>
            """
            $(nTd).html content
            return
        }
        {
          'data': 'pdf'
          'fnCreatedCell': (nTd, sData, oData, iRow, iCol) ->
            content = """
              <a href='#{oData.pdf}' class='btn btn-primary float-right' target='_blank'><i data-feather='download' class='mr-3 icon-md'></i>Download</a>
            """
            $(nTd).html content
            return
        }
      ]
      'language': {
        'infoFiltered': ''
      }
      'preDrawCallback': (settings) ->
        s.tables.preDraw($(this), settings)
      'drawCallback': ->
        s.template.renderIcons()

    this.tables.cleanup($('table#account-subscription-invoice-table'))

    return

  #
  # New.
  #
  subscribe: ->
    stripe = Stripe($('.stripe-pk').data('stripe-pk'))
    $('button.subscribe').on 'click', ->
      $('button.subscribe').attr('disabled', true)

      stripe.redirectToCheckout(sessionId: $(this).data('stripe-session')).then (result) ->
        $('button.subscribe').removeAttr('disabled')
        alert result.error.message
      return

    return

  #
  # (Re)Subscribe.
  #
  reactivate: ->
    s = this

    $('button.reactivate').on 'click', (e) ->
      e.preventDefault()

      s.alerts.reset()
      s.turbolinks.start()

      $('button.processing').fadeIn()
      $('button.reactivate').attr('disabled', true)

      csrfToken = $('input[name="gorilla.csrf.Token"]').val()
      planId = $(this).data('stripe-plan')

      $.ajax(
        type: 'PUT'
        url: s.url.clean() + '/reactivate/' + planId
        timeout: 10000
        headers: 'X-CSRF-Token': csrfToken
        statusCode:
          200: ->
            s.turbolinks.stop()

            $('button.processing').fadeOut()
            $('button.reactivate').removeAttr('disabled')

            $('button.reactivate').each () ->
              $(this).addClass('hide')

              if $(this).data('stripe-plan') == planId
                $(this).next().next().removeClass('hide')
              else
                $(this).next().removeClass('hide')
              return

            $('.alert.account-subscription-change-success').fadeIn()

            $('button.loading').fadeIn()
            $('#account-subscription-invoice-table').DataTable().ajax.reload()
          500: ->
            s.turbolinks.stop()

            $('.alert.account-subscription-change-error').fadeIn(400, ->
              $('button.processing').fadeOut()
              $('button.reactivate').removeAttr('disabled')
            )
      ).fail ->
        s.turbolinks.stop()

        $('.alert.account-subscription-change-error').fadeIn(400, ->
          $('button.processing').fadeOut()
          $('button.reactivate').removeAttr('disabled')
        )
      return

    return

  #
  # Change.
  #
  change: ->
    s = this

    $('button.change').on 'click', (e) ->
      e.preventDefault()

      s.alerts.reset()
      s.turbolinks.start()

      $('button.processing').fadeIn()
      $('button.change, button.cancel').attr('disabled', true)

      planId = $(this).data('stripe-plan')
      csrfToken = $('input[name="gorilla.csrf.Token"]').val()

      $.ajax(
        type: 'PUT'
        url: s.url.clean() + '/update/' + planId
        timeout: 10000
        headers: 'X-CSRF-Token': csrfToken
        statusCode:
          200: ->
            s.turbolinks.stop()

            $('button.processing').fadeOut()
            $('button.change, button.cancel').removeAttr('disabled')

            $('button.change').each () ->
              if $(this).data('stripe-plan') == planId
                $(this).addClass('hide')
                $(this).next().removeClass('hide')
              else if $(this).next().is(':visible')
                $(this).removeClass('hide')
                $(this).next().addClass('hide')
              return

            $('.alert.account-subscription-change-success').fadeIn()

            $('button.loading').fadeIn(400, ->
              $('#account-subscription-invoice-table').DataTable().ajax.reload()
            )
          500: ->
            s.turbolinks.stop()

            $('.alert.account-subscription-change-error').fadeIn(400, ->
              $('button.processing').fadeOut()
              $('button.change, button.cancel').removeAttr('disabled')
            )
      ).fail ->
        s.turbolinks.stop()

        $('.alert.account-subscription-change-error').fadeIn(400, ->
          $('button.processing').fadeOut()
          $('button.change, button.cancel').removeAttr('disabled')
        )

      return

    return

  #
  # Cancel.
  #
  cancel: ->
    s = this

    $('#account-subscription-cancel-modal').on 'shown.bs.modal', ->
      $('form#account-subscription-cancel').submit (e) ->
        e.preventDefault()

        s.alerts.resetErrors()
        s.turbolinks.start()

        $('button.close-modal').attr('disabled', 'true')
        $('button.submit').html('Processing...').attr('disabled', 'true')

        csrfToken = $('input[name="gorilla.csrf.Token"]').val()

        $.ajax(
          type: 'PUT'
          url: s.url.clean() + '/cancel'
          data: $('form#account-subscription-cancel').serialize()
          timeout: 10000
          headers: 'X-CSRF-Token': csrfToken
          statusCode:
            200: ->
              s.turbolinks.stop()

              location.reload()
            422: ->
              s.stop()

              $('.alert.account-subscription-cancel-validation-error').fadeIn()
              $('input#cancel').addClass('is-invalid')
            500: ->
              s.stop()

              $('.alert.account-subscription-cancel-error').fadeIn()
        ).fail ->
          s.stop()

          $('.alert.account-subscription-cancel-error').fadeIn()

      return

    return

  #
  # Stop.
  #
  stop: ->
    this.turbolinks.stop()
    this.modals.reset('Submit')

    return
