import 'bootstrap-datepicker'

#
# Filters
#
export default class FiltersHelper
  constructor: ->

  #
  # Watch the Date Filter.
  #
  watchDate: (obj) ->
    $('.date-filter').on 'click', (e) ->
      e.preventDefault()
      $('.date-filter.active').removeClass('active')
      $(this).addClass('active')

      obj.load()

      return

    return

  #
  # Watch the view filter.
  #
  watchView: (obj) ->
    $('.view-filter').on 'click', (e) ->
      e.preventDefault()
      $('.view-filter.btn-primary').removeClass('btn-primary').removeClass('text-white').addClass('btn-outline-primary')
      $(this).addClass('btn-primary text-white')

      obj.load()

      return

    return

  #
  # Datepicker.
  #
  datePicker: (container) ->
    date = new Date
    today = new Date(date.getFullYear(), date.getMonth(), date.getDate())
    $(container).datepicker
      format: 'mm/dd/yyyy'
      todayHighlight: true
      autoclose: true
    $(container).datepicker 'setDate', today

    return
