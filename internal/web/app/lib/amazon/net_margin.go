package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

func (a *Amazon) NetMargin(txns *[]models.Transaction) []types.Summary {
	var netMargin []types.Summary

	for _, txn := range *txns {
		netMargin = append(netMargin, types.Summary{
			Total:       a.txnNetMargin(txn) * a.exchangeRate(txn.Marketplace.Id),
			Marketplace: *a.marketplace(txn.Marketplace.Id),
			OrderDate:   txn.DateTime,
		})
	}

	return netMargin
}
