import $ from 'jquery'
window.jQuery = $
window.$ = $

require('datatables.net') window, $
require('datatables.net-bs4') window, $

#
# Tables.
#
export default class TablesHelper
  constructor: ->

#
# DT pre-draw.
#
  preDraw: (obj, settings) ->
    api = new ($.fn.dataTable.Api)(settings)
    pagination = $(obj).closest('.dataTables_wrapper').find('.dataTables_paginate')
    pagination.toggle api.page.info().pages > 1

    return

#
# DT Cleanup.
#
  cleanup: (table) ->
    table.each ->
      datatable = $(this)
      length_sel = datatable.closest('.dataTables_wrapper').find('div[id$=_length] select')
      length_sel.removeClass('form-control-sm')
      return

    return

#
# DT Reload.
#
  reload: (obj, table) ->
    $('button.loading').fadeIn()

    $('.date-filter.active').removeClass('active')
    $(obj).addClass('active')
    table.DataTable().ajax.url(window.location.href + '/filter/' + $(obj).data('range')).load()

    return
