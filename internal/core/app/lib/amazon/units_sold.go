package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (a *Amazon) UnitsSold(txns *[]models.Transaction) []types.Summary {
	var unitsSold []types.Summary

	for _, txn := range *txns {
		if a.isOrder(txn.TransactionType.Id) {
			unitsSold = append(unitsSold, types.Summary{
				SKU:         txn.SKU,
				Description: txn.Description,
				Marketplace: *a.marketplace(txn.Marketplace.Id),
				Total:       a.txnUnitsSold(txn),
				OrderDate:   txn.DateTime,
			})
		}
	}

	return unitsSold
}
