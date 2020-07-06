import AlertsHelper     from '../helpers/alerts'
import TurbolinksHelper from '../helpers/turbolinks'

#
# Account: Password
#
export default class AccountPassword
  constructor: ->
    this.alerts = new AlertsHelper()
    this.turbolinks = new TurbolinksHelper()

  #
  # Init.
  #
  init: ->
    return

  #
  # Reset.
  #
  reset: ->
    this.alerts.reset()
    this.turbolinks.start()

    p = this

    $.ajax
      type: 'GET'
      url: '/account/change_password'
      statusCode:
        200: ->
          p.turbolinks.stop()

          $('.alert.alert-success.account-change-password').fadeIn()
        500: ->
          p.turbolinks.stop()

          $('.alert.alert-danger.account-change-password').fadeIn()

    return
