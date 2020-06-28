package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

func (a *Amazon) ShippingCredits(txns *[]models.Transaction) []types.Summary {
	var shippingCredits []types.Summary

	for _, txn := range *txns {
		shippingCredits = append(shippingCredits, types.Summary{
			SKU:                txn.SKU,
			Description:        txn.Description,
			Marketplace:        *a.marketplace(txn.Marketplace.Id),
			ShippingCredits:    txn.ShippingCredits * a.exchangeRate(txn.Marketplace.Id),
			ShippingCreditsTax: txn.ShippingCreditsTax * a.exchangeRate(txn.Marketplace.Id),
			Total:              a.txnShippingCredits(txn) * a.exchangeRate(txn.Marketplace.Id),
			Date:               txn.DateTime,
		})
	}

	return shippingCredits
}
