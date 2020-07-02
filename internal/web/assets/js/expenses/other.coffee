import $ from 'jquery'
window.jQuery = $
window.$ = $

import AlertsHelper   from '../helpers/alerts'
import FiltersHelper  from '../helpers/filters'
import ModalsHelper   from '../helpers/modals'
import TablesHelper   from '../helpers/tables'
import TemplateHelper from '../helpers/template'
import URLHelper      from '../helpers/url'

require('datatables.net') window, $
require('datatables.net-bs4') window, $

#
# Expenses.
#
export default class ExpensesOther
  constructor: ->
    this.alerts = new AlertsHelper()
    this.filters = new FiltersHelper()
    this.modals = new ModalsHelper()
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
    ex = this

    $('button.loading').fadeIn()

    $('table').DataTable
      'serverSide': true,
      'bFilter': false
      'lengthChange': false
      'ajax':
        'url': ex.url.clean() + '/filter/all_time'
        'dataSrc': (j) ->
          $('button.loading').fadeOut()

          if j.data.length > 0
            ex.alerts.reset()

          return j.data
        'error': (j) ->
          $('.alert.expenses-other-load-error').show()
      'columns': [
        { 'data': 'description' }
        { 'data': 'dateTime' }
        { 'data': 'cost' }
      ]
      'language': {
        'infoFiltered': ''
      }
      'preDrawCallback': (settings) ->
        ex.tables.preDraw($(this), settings)
      'drawCallback': ->
        ex.template.renderIcons()

    $('.date-filter').on 'click', (e) ->
      e.preventDefault()
      ex.tables.reload($(this), $('table'))

    this.tables.cleanup($('table'))

    return

  #
  # New cost of good.
  #
  new: ->
    ex = this

    $('#expenses-other-modal').on 'shown.bs.modal', ->
      ex.alerts.resetSuccess()
      ex.alerts.resetErrors()
      ex.modals.resetForm()

      ex.filters.datePicker('#expenses-other-modal .datepicker')

      $('form#expenses-other').on 'submit', (e) ->
        e.preventDefault()

        return

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
    this.modals.reset('Submit')

    return