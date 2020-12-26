import $ from 'jquery'
window.jQuery = $
window.$ = $

#
# Onboard.
#
export default class HomeOnboard
  constructor: ->

  #
  # Init.
  #
  init: ->
    this.mailingList()

  #
  # Setup completed.
  #
  setupCompleted: (csrfToken) ->
    $.ajax(
      type: 'PUT'
      url: '/account/preferences/setup_completed'
      timeout: 10000
      data:
        setup_completed: true
      headers: 'X-CSRF-Token': csrfToken
      statusCode:
        200: (j) ->
          $('.preferences').hide()
    ).fail ->

    return

  #
  # Mailing list.
  #
  mailingList: ->
    $('button.mailing-list').on 'click', (e) ->
      e.preventDefault()

      subscribe = $(this).data('subscribe')
      csrfToken = $('input[name="gorilla.csrf.Token"]').val()

      $.ajax(
        type: 'PUT'
        url: '/account/preferences/mailing_list'
        timeout: 10000
        data:
          mailing_list: subscribe
        headers: 'X-CSRF-Token': csrfToken
        statusCode:
          200: (j) ->
            $('.preferences').hide()
      ).fail ->

      return
