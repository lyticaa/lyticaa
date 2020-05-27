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

  loadDashboardMetrics()

  $('.date-filter').on 'click', (e) ->
    e.preventDefault()
    $('.date-filter.active').removeClass 'active'
    $(this).addClass 'active'

    loadDashboardMetrics()

    return
  return

#
# Load Metrics.
#
loadDashboardMetrics = ->
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

        # feather.replace()
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
# Reset errors.
#
resetErrors = ->
  $('.alert.alert-danger').each ->
    if $(this).is(':visible')
      $(this).hide()
    return

#
# Reset warnings.
#
resetWarnings = ->
  $('.alert.alert-warning').each ->
    if $(this).is(':visible')
      $(this).hide()
    return

#
# Start Turbolinks progress bar.
#
tbStart = ->
  Turbolinks.controller.adapter.progressBar.setValue 0
  Turbolinks.controller.adapter.progressBar.show()

#
# Stop Turbolinks progress bar.
#
tbStop = ->
  Turbolinks.controller.adapter.progressBar.hide()

#
# Init
#
Turbolinks.start()
$(document).on 'ready, turbolinks:load', ->
  initialize()
  return
