import AlertsHelper     from '../helpers/alerts'
import ModalsHelper     from '../helpers/modals'
import TurbolinksHelper from '../helpers/turbolinks'

import HomeOnboard      from '../home/onboard'

#
# Reports.
#
export default class Reports
  constructor: ->
    this.alerts = new AlertsHelper()
    this.modals = new ModalsHelper()
    this.turbolinks = new TurbolinksHelper()

  #
  # Initialize.
  #
  init: ->
    this.import()

    return

  #
  # Import.
  #
  import: ->
    r = this

    $('#import-modal').on 'shown.bs.modal', ->
      r.alerts.resetSuccess()
      r.alerts.resetErrors()

      r.modals.resetForm()
      r.modals.reset('Submit')
      $('.dropify-clear').trigger 'click'

      $('.modal:visible form').on 'submit', (e) ->
        e.preventDefault()
        r.start('Importing')

        $.ajax(
          type: 'PUT'
          enctype: 'multipart/form-data'
          url: '/reports/import'
          processData: false
          contentType: false
          data: new FormData($('.modal:visible form')[0])
          timeout: 10000
          statusCode:
            200: (j) ->
              $('.alert.import-success').fadeIn(400, ->
                r.stop()
                $('.dropify-clear').trigger 'click'
                r.setupCompleted()
              )
        ).fail ->
          $('.alert.import-error').fadeIn(400, ->
            r.stop()
            $('.dropify-clear').trigger 'click'
          )

    return

  #
  # Setup Completed.
  #
  setupCompleted: ->
    if $('input.location').data('section') != 'home-onboard'
      return

    csrfToken = $('input[name="gorilla.csrf.Token"]').val()

    o = new HomeOnboard()
    o.setupCompleted(csrfToken)

    window.location = '/'

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
    this.modals.resetForm()
    this.modals.reset('Submit')

    return