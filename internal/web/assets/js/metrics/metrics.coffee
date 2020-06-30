import AlertsHelper from '../helpers/alerts'
import TablesHelper from '../helpers/tables'

#
# Metrics.
#
export default class Metrics
  constructor: ->
    this.alerts = new AlertsHelper()
    this.tables = new TablesHelper()

  #
  # Reload.
  #
  reload: ->
    p = this
    p.alerts.reset()

    $('button.reload').on 'click', (e) ->
      p.tables.reload($('.date-filter.active'), $('table'))
      return

    return
