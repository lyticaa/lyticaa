import Turbolinks from 'turbolinks'

import $ from 'jquery'
window.jQuery = $
window.$ = $

#
# Intercom.
#
Turbolinks.start()
$(document).on 'ready, turbolinks:load', ->
  intercomId = $('.intercom').data('intercom-id')

  window.intercomSettings =
    app_id: intercomId
    name: $('.intercom').data('name')
    email: $('.intercom').data('email')
    created_at: $('.intercom').data('created-at')

  do ->
    w = window
    ic = w.Intercom
    if typeof ic == 'function'
      ic 'reattach_activator'
      ic 'update', w.intercomSettings
    else
      d = document

      i = ->
        i.c arguments
        return

      i.q = []

      i.c = (args) ->
        i.q.push args
        return

      w.Intercom = i

      l = ->
        s = d.createElement('script')
        s.type = 'text/javascript'
        s.async = true
        s.src = 'https://widget.intercom.io/widget/' + intercomId
        x = d.getElementsByTagName('script')[0]
        x.parentNode.insertBefore s, x
        return

      if w.attachEvent
        w.attachEvent 'onload', l
      else
        w.addEventListener 'load', l, false
    return

  $('a.contact-support').on 'click', (e) ->
    e.preventDefault()
    window.Intercom('show')
    return

  return
