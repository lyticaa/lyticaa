package types

var Ignore = []string{
	"Includes Amazon Marketplace, Fulfillment by Amazon (FBA), and Amazon Webstore transactions",
	"All amounts in USD, unless specified",
	"Definitions:",
	"Sales tax collected: Includes sales tax collected from buyers for product sales, shipping, and gift wrap.",
	"Selling fees: Includes variable closing fees and referral fees.",
	"Other transaction fees: Includes shipping chargebacks, shipping holdbacks, per-item fees  and sales tax collection fees.",
	"Other: Includes non-order transaction amounts. For more details, see the \"Type\" and \"Description\" columns for each order ID.",
}

func ShouldIgnore(val string) bool {
	for _, item := range Ignore {
		if item == val {
			return true
		}
	}

	return false
}
