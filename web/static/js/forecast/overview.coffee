import AlertsHelper   from '../helpers/alerts'
import TemplateHelper from '../helpers/template'
import URLHelper      from '../helpers/url'
import FiltersHelper  from '../helpers/filters'

#
# Forecast: Overview.
#
export default class ForecastOverview
  constructor: ->
    this.alerts = new AlertsHelper()
    this.template = new TemplateHelper()
    this.url = new URLHelper()
    this.filters = new FiltersHelper()

  #
  # Initialize.
  #
  init: ->
    this.load()
    this.filters.watchDate(this)

    return

  #
  # Load.
  #
  load: ->
    this.alerts.reset()

    d = this

    $('button.loading').fadeIn()

    $.ajax(
      type: 'GET'
      url: d.url.clean() + '/filter/' + $('.date-filter.active').data('range')
      timeout: 10000
      statusCode:
        200: (j) ->
          $('button.loading').fadeOut(400, ->
          )
        422: ->
          $('button.loading').fadeOut(400, ->
            $('.alert.forecast-load-error').fadeIn()
          )
        500: ->
          $('button.loading').fadeOut(400, ->
            $('.alert.forecast-load-error').fadeIn()
          )
    ).fail ->
      $('button.loading').fadeOut(400, ->
        $('.alert.forecast-load-error').fadeIn()
      )

    return
