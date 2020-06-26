package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (a *Amazon) GrossMargin(txns *[]models.Transaction) []types.Summary {
	var grossMargin []types.Summary

	for _, txn := range *txns {
		txnTypeId := txn.TransactionType.Id
		if a.isOrder(txnTypeId) {
			grossMargin = append(grossMargin, types.Summary{
				Total:       a.txnGrossMargin(txn) * a.exchangeRate(txn.Marketplace.Id),
				Marketplace: *a.marketplace(txn.Marketplace.Id),
				OrderDate:   txn.DateTime,
			})
		}
	}

	return grossMargin
}
