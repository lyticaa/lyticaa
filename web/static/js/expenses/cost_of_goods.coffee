import $ from 'jquery'
window.jQuery = $
window.$ = $

import Expenses from './expenses'

import AlertsHelper     from '../helpers/alerts'
import TablesHelper     from '../helpers/tables'
import TemplateHelper   from '../helpers/template'
import URLHelper        from '../helpers/url'

require('datatables.net') window, $
require('datatables.net-bs4') window, $

#
# Expenses.
#
export default class ExpensesCostOfGoods
  constructor: ->
    this.e = new Expenses()
    this.alerts = new AlertsHelper()
    this.tables = new TablesHelper()
    this.template = new TemplateHelper()
    this.url = new URLHelper()

  #
  # Initialize.
  #
  init: ->
    this.drawTable()
    this.new()

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
        'url': ex.url.clean() + '/all'
        'dataSrc': (j) ->
          $('button.loading').fadeOut()

          if j.data.length > 0
            ex.alerts.reset()

          return j.data
        'error': (j) ->
          $('button.loading').fadeOut(400, ->
            $('.alert.expenses-load-error').show()
          )
      'columns': [
        { 'data': 'marketplace' }
        { 'data': 'sku' }
        { 'data': 'description' }
        { 'data': 'fromDate' }
        { 'data': 'amount' }
        {
          'data': null
          'className': 'text-center pr-0 pl-0 w-15'
          'fnCreatedCell': (nTd, sData, oData, iRow) ->
            content = """
                <a href='#' class='expenses-edit' data-toggle='modal' data-target='#expenses-edit-modal' data-expense='#{oData.DT_RowId}' data-product='#{oData.productId}' data-sku='#{oData.sku}' data-marketplace='#{oData.marketplace}' data-description='#{oData.description}' data-from-date='#{oData.fromDate}' data-amount='#{oData.amount}' target='_blank'>
                  <i class='edit' data-feather='edit' data-toggle='tooltip' data-placement='top' title='Edit the expense.'></i>
                </a>
                <i class='delete' data-expense='#{oData.DT_RowId}' data-idx='#{iRow}' data-feather='trash' data-toggle='tooltip' data-placement='top' title='Delete the expense.'></i>
              """
            $(nTd).html content
            return
        }
      ]
      'language': {
        'infoFiltered': ''
      }
      'preDrawCallback': (settings) ->
        ex.tables.preDraw($(this), settings)
      'drawCallback': ->
        ex.template.renderIcons()

        ex.e.populate()
        ex.edit()
        ex.delete()

    $('.date-filter').on 'click', (e) ->
      e.preventDefault()
      ex.tables.reload($(this), $('table'))

    this.tables.cleanup($('table'))

    return

  #
  # Delete expense.
  #
  delete: ->
    this.e.delete()

    return

  #
  # New expense.
  #
  new: ->
    ex = this

    $('#expenses-new-modal').on 'shown.bs.modal', ->
      ex.e.new(true, false)

    return

  #
  # Edit expense.
  #
  edit: ->
    ex = this

    $('#expenses-edit-modal').on 'shown.bs.modal', ->
      ex.e.edit()

    return
