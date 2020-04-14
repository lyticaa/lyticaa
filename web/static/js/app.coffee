import Turbolinks from 'turbolinks'

import $ from 'jquery'
window.jQuery = $
window.$ = $

import 'js-cookie'
import 'dropzone'

initialize = ->
  user()
  uploads()
  payments()
  
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
# Init
#
Turbolinks.start()
$(document).on 'ready, turbolinks:load', ->
  initialize()
  return
