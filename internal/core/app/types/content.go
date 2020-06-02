package types

var (
	FlashMessages = map[string]map[string]map[string]string{
		"setup": {
			"subscribe": {
				"success": "Thanks for subscribing. Please click Next to continue.",
				"error":   "There was an issue while processing your subscription. Please try again.",
			},
		},
		"account": {
			"subscription": {
				"cancel": "Your cancellation request was processed successfully.",
			},
		},
	}
	NotificationMessages = map[string]string{
		"customer.subscription.created": "Subscribed to the plan %v.",
		"customer.subscription.deleted": "Unsubscribed from the plan %v.",
		"invoice.created":               "A new invoice %v for %v was created.",
		"invoice.payment_failed":        "Payment of the invoice %v for %v was unsuccessful.",
		"invoice.payment_succeeded":     "Payment of the invoice %v for %v was successful.",
	}
)
