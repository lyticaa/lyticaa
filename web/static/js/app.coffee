$ = require('jquery')
Turbolinks = require('turbolinks')

require 'bootstrap-datepicker'
require 'js-cookie'

initialize = ->
  # Update profile image.
  imageUrl = $('.profile-image').attr('rel')
  $('.profile-image-thumb, .profile-image').attr 'src', imageUrl

  # Datepicker
  if $('#dashboardDate').length
    date = new Date
    today = new Date(date.getFullYear(), date.getMonth(), date.getDate())
    $('#dashboardDate').datepicker
      format: 'dd-MM-yyyy'
      todayHighlight: true
      autoclose: true
    $('#dashboardDate').datepicker 'setDate', today
  
  # Logout
  $(document).ready ->
    $('.log-out').click (e) ->
      Cookies.remove 'auth-session'
      return
    return
  return


Turbolinks.start()
$(document).on 'turbolinks:load', ->
  initialize()
  return
