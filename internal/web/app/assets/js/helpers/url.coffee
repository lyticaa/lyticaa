#
# URL Utils.
#
export default class URLHelper
  constructor: ->

  clean: ->
    return window.location.href.split("#")[0]
