package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (a *Amazon) AmazonCosts(txns *[]models.Transaction) []types.Summary {
	var amazonCosts []types.Summary

	for _, txn := range *txns {
		txnTypeId := txn.TransactionType.Id
		if a.isOrder(txnTypeId) || a.isFBAInventoryFee(txnTypeId) || a.isServiceFee(txnTypeId) {
			amazonCosts = append(amazonCosts,
				types.Summary{
					Total:       a.txnAmazonCosts(txn) * a.exchangeRate(txn.Marketplace.Id),
					Marketplace: *a.marketplace(txn.Marketplace.Id),
					OrderDate:   txn.DateTime,
				},
			)
		}
	}

	return amazonCosts
}
