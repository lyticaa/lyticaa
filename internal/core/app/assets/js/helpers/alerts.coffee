#
# Alerts helper.
#
export default class AlertsHelper
  constructor: ->

  #
  # Reset.
  #
  reset: ->
    this.resetSuccess()
    this.resetErrors()
    this.resetWarnings()

  #
  # Reset success.
  #
  resetSuccess: ->
    $('input').each ->
      if $(this).hasClass('is-valid')
        $(this).removeClass('is-valid')
      return

    $('.alert.alert-success, .alert.alert-icon-success').each ->
      if $(this).is(':visible') && $(this).not('.flash')
        $(this).fadeOut()
      return

    return

  #
  # Reset errors.
  #
  resetErrors: ->
    $('input').each ->
      if $(this).hasClass('is-invalid')
        $(this).removeClass('is-invalid')
      return

    $('.alert.alert-danger, .alert.alert-icon-danger').each ->
      if $(this).is(':visible') && $(this).not('.flash')
        $(this).fadeOut()
      return

    return

  #
  # Reset warnings.
  #
  resetWarnings: ->
    $('.alert.alert-warning, .alert.alert-icon-warning').each ->
      if $(this).is(':visible') && $(this).not('.flash')
        $(this).fadeOut()
      return

    return
