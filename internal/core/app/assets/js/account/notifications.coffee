import $ from 'jquery'
window.jQuery = $
window.$ = $

import AlertsHelper   from '../helpers/alerts'
import TablesHelper   from '../helpers/tables'
import TemplateHelper from '../helpers/template'
import URLHelper      from '../helpers/url'

require('datatables.net') window, $
require('datatables.net-bs4') window, $

#
# Account: Notifications.
#
export default class AccountNotifications
  constructor: ->
    this.alerts = new AlertsHelper()
    this.tables = new TablesHelper()
    this.template = new TemplateHelper()
    this.url = new URLHelper()

  #
  # Initialize.
  #
  init: ->
    this.drawTable()

    return

  #
  # Draw table.
  #
  drawTable: ->
    n = this

    $('button.loading').fadeIn()

    $('table').DataTable
      'serverSide': true,
      'bFilter': false
      'lengthChange': false
      'ajax':
        'url': n.url.clean() + '/filter/all_time'
        'dataSrc': (j) ->
          $('button.loading').fadeOut()

          if j.data.length > 0
            n.alerts.reset()

          return j.data
        'error': (j) ->
          $('button.loading').fadeOut(400, ->
            $('.alert.account-notifications-load-error').fadeIn()
          )
      'columns': [
        { 'data': 'notification' }
        { 'data': 'date' }
      ]
      'language': {
        'infoFiltered': ''
      }
      'preDrawCallback': (settings) ->
        n.tables.preDraw($(this), settings)
      'drawCallback': ->
        n.template.renderIcons()

    $('.date-filter').on 'click', (e) ->
      e.preventDefault()
      n.tables.reload($(this), $('table'))

    this.tables.cleanup($('table'))

    return
