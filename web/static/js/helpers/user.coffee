import 'js-cookie'

#
# User.
#
export default class UserHelper
  constructor: ->
    imageUrl = $('.profile-image').attr('rel')
    $('.profile-image-thumb, .profile-image').attr 'src', imageUrl

    $('.log-out').click (e) ->
      Cookies.remove 'auth-session'
      return
    return
