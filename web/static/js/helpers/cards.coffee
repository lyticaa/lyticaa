#
# Cards helper.
#
export default class CardsHelper
  constructor: ->

  #
  # Paint.
  #
  paint: (cssClass, total, diff) ->
    $("h3.#{cssClass}").text(total)
    $("span.diff.#{cssClass}").text("#{diff}%")
    $("span.diff.#{cssClass}").next().remove()

    if diff > 0
      arrow = 'arrow-up'
      $("span.diff.#{cssClass}").parent().removeClass().addClass 'text-success'
    else if diff == 0
      arrow = 'minus'
      $("span.diff.#{cssClass}").parent().removeClass().addClass 'text-warning'
    else
      arrow = 'arrow-down'
      $("span.diff.#{cssClass}").parent().removeClass().addClass 'text-danger'

    $("span.diff.#{cssClass}").parent().append(
      $('<i>')
        .attr('data-feather', arrow)
        .addClass('icon-sm')
        .addClass('mb-1')
    )

    return
