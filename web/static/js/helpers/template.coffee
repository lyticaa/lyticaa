import $ from 'jquery'
window.jQuery = $
window.$ = $

import 'bootstrap'
import feather from 'feather-icons'

#
# Template.
#
export default class TemplateHelper
  constructor: ->

#
# Render icons.
#
  renderIcons: ->
    feather.replace()
    $('[data-toggle="tooltip"]').tooltip()

    return
