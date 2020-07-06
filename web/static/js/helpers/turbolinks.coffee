import Turbolinks from 'turbolinks'

#
# Turbolinks.
#
export default class TurbolinksHelper
  constructor: ->

#
# Start.
#
  start: ->
    Turbolinks.controller.adapter.progressBar.setValue 0
    Turbolinks.controller.adapter.progressBar.show()

    return

#
# Stop Turbolinks progress bar.
#
  stop: ->
    Turbolinks.controller.adapter.progressBar.hide()

    return
