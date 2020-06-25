import TablesHelper from '../helpers/tables'

#
# Metrics.
#
export default class Metrics
  constructor: ->
    this.tables = new TablesHelper()

  #
  # Reload.
  #
  reload: ->
    p = this
    $('button.reload').on 'click', (e) ->
      p.tables.reload($('.date-filter.active'), $('table'))
      return

    return
