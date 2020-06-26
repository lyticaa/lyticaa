package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (a *Amazon) ShippingCredits(txns *[]models.Transaction) []types.Summary {
	var shippingCredits []types.Summary

	for _, txn := range *txns {
		shippingCredits = append(shippingCredits, types.Summary{
			Total:       a.txnShippingCredits(txn) * a.exchangeRate(txn.Marketplace.Id),
			Marketplace: *a.marketplace(txn.Marketplace.Id),
			OrderDate:   txn.DateTime,
		})
	}

	return shippingCredits
}
