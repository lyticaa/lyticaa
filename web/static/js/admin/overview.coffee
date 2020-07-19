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
# Admin: Overview.
#
export default class AdminOverview
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
    this.reload()

    return

  #
  # Draw table.
  #
  drawTable: ->
    a = this

    $('button.loading').fadeIn()

    $('table').DataTable
      'serverSide': true,
      'bFilter': false
      'lengthChange': false
      'ajax':
        'url': a.url.clean() + '/filter/all_time'
        'dataSrc': (j) ->
          $('button.loading').fadeOut()

          if j.data.length > 0
            a.alerts.reset()

          return j.data
        'error': (j) ->
          $('button.loading').fadeOut(400, ->
            $('.alert.admin-load-error').fadeIn()
          )
      'columns': [
        { 'data': 'email' }
        { 'data': 'date' }
        {
          'data': null
          'className': 'text-center pr-0 pl-0 w-15'
          'fnCreatedCell': (nTd, sData, oData, iRow) ->
            content = """
              <a href='/admin/i/#{oData.DT_RowId}' data-turbolinks='false'>
                <i class='impersonate' data-feather='maximize' data-toggle='tooltip' data-placement='top' title='Impersonate this user.' ></i>
              </a>
              """
            $(nTd).html content
            return
        }
      ]
      'language': {
        'infoFiltered': ''
      }
      'preDrawCallback': (settings) ->
        a.tables.preDraw($(this), settings)
      'drawCallback': ->
        a.template.renderIcons()

    $('.date-filter').on 'click', (e) ->
      e.preventDefault()
      a.tables.reload($(this), $('table'))

    this.tables.cleanup($('table'))

    return

  #
  # Reload.
  #
  reload: ->
    a = this
    $('button.reload').on 'click', (e) ->
      a.tables.reload($('.date-filter.active'), $('table'))
      return

    return
