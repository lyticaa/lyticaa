import AlertsHelper     from '../helpers/alerts'
import ModalsHelper     from '../helpers/modals'
import TurbolinksHelper from '../helpers/turbolinks'

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
              )
        ).fail ->
          $('.alert.import-error').fadeIn(400, ->
            r.stop()
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
    this.modals.resetForm()
    this.modals.reset('Submit')

    return