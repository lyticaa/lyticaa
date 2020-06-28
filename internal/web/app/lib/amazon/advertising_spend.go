package amazon

import (
	"regexp"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

var (
	advertisingSpendLineItem = regexp.MustCompile(`^Cost of Advertising+$`).MatchString
)

func (a *Amazon) AdvertisingSpend(txns *[]models.Transaction, sps *[]models.SponsoredProduct) []types.Summary {
	var advertisingSpend []types.Summary

	for _, sp := range *sps {
		advertisingSpend = append(advertisingSpend, types.Summary{
			SKU:                        sp.SKU,
			Description:                "",
			Marketplace:                "",
			AdvertisingSpend:           0.0,
			AdvertisingSpendPercentage: 0.0,
		})
	}

	return advertisingSpend
}
