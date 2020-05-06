import Turbolinks from 'turbolinks'

import $ from 'jquery'
window.jQuery = $
window.$ = $

import 'js-cookie'
import 'dropzone'

require('datatables.net') window, $
require('datatables.net-bs4') window, $

initialize = ->
  user()
  support()
  uploads()
  payments()
  pageComponents()

  # Logout.
  $(document).ready ->
    $('.log-out').click (e) ->
      Cookies.remove 'auth-session'
      return
    return
  return

#
# User
#
user = ->
  imageUrl = $('.profile-image').attr('rel')
  $('.profile-image-thumb, .profile-image').attr 'src', imageUrl
  return

#
# Support (intercom)
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
# Uploads
#
uploads = ->
  if $('#dropzone').length > 0
    $('#dropzone').dropzone
  return

#
# Payments
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
# Page components
#
pageComponents = ->
  if $("#dataTable").length > 0
    $('#dataTable').DataTable
      'bFilter': false
      'lengthChange': false
      preDrawCallback: (settings) ->
        api = new ($.fn.dataTable.Api)(settings)
        pagination = $(this).closest('.dataTables_wrapper').find('.dataTables_paginate')
        pagination.toggle api.page.info().pages > 1

    $('#dataTable').each ->
      datatable = $(this)
      length_sel = datatable.closest('.dataTables_wrapper').find('div[id$=_length] select')
      length_sel.removeClass 'form-control-sm'
      return
  return

#
# Init
#
Turbolinks.start()
$(document).on 'ready, turbolinks:load', ->
  initialize()
  return
