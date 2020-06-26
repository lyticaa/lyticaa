#
# Setup: Subscribe.
#
export default class SetupSubscribe
  constructor: ->

  #
  # Initialize.
  #
  init: ->
    stripe = Stripe($('.stripe-pk').data('stripe-pk'))
    $('button.subscribe').on 'click', ->
      stripe.redirectToCheckout(sessionId: $(this).data('stripe-session')).then (result) ->
        alert result.error.message
      return

    return
