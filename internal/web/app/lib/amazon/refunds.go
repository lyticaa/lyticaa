package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

func (a *Amazon) Refunds(txns *[]models.Transaction) []types.Summary {
	var refunds []types.Summary

	for _, txn := range *txns {
		txnTypeId := txn.TransactionType.Id
		if a.isRefund(txnTypeId) {
			refunds = append(refunds, types.Summary{
				Total:       txn.Total * a.exchangeRate(txn.Marketplace.Id),
				Marketplace: *a.marketplace(txn.Marketplace.Id),
				OrderDate:   txn.DateTime,
			})
		}
	}

	return refunds
}
