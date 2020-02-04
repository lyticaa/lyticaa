package types

var FlashMessages = map[string]map[string]map[string]string{
	"setup": {
		"subscribe": {
			"success": "Thanks for subscribing.",
			"error":   "There was an issue while processing your subscription. Please try again.",
		},
	},
}
