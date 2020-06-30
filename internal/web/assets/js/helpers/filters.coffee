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
