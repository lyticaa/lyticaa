require('datatables.net') window, $
require('datatables.net-bs4') window, $

$(document).on 'ready, turbolinks:load', ->
  if $("#dataTable").length > 0
    $('#dataTable').DataTable
      'bFilter': false
      'lengthChange': false
      preDrawCallback: (settings) ->
        api = new ($.fn.dataTable.Api)(settings)
        pagination = $(this).closest('.dataTables_wrapper').find('.dataTables_paginate')
        pagination.toggle api.page.info().pages > 1

    $('#dataTable').each ->
      datatable = $(this)
      length_sel = datatable.closest('.dataTables_wrapper').find('div[id$=_length] select')
      length_sel.removeClass 'form-control-sm'
      return
