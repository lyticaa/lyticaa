package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (a *Amazon) UnitsSold(txns *[]models.Transaction) []types.Summary {
	var unitsSold []types.Summary

	for _, txn := range *txns {
		if a.isOrder(txn.TransactionType.Id) {
			unitsSold = append(unitsSold,
				types.Summary{
					Total:       float64(txn.Quantity),
					Marketplace: *a.marketplace(txn.Marketplace.Id),
					OrderDate:   txn.DateTime,
				},
			)
		}
	}

	return unitsSold
}
