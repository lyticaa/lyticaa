#
# Modals.
#
export default class ModalsHelper
  constructor: ->

  #
  # Disable.
  #
  disable: (text) ->
    $('.modal').each ->
      if $(this).is(':visible')
        $('button.close-modal').attr('disabled')
        $('button[type="submit"]').html(text).attr('disabled')
    return

  #
  # Reset.
  #
  reset: (text) ->
    $('.modal').each ->
      if $(this).is(':visible')
        $('button.close-modal').removeAttr('disabled')
        $('button[type="submit"]').html(text).removeAttr('disabled')
    return

  #
  # Reset form.
  #
  resetForm: ->
    $('.modal').each ->
      if $(this).is(':visible')
        $('form').trigger 'reset'
    return
