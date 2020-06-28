package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

func (a *Amazon) AmazonCosts(txns *[]models.Transaction) []types.Summary {
	var amazonCosts []types.Summary

	for _, txn := range *txns {
		txnTypeId := txn.TransactionType.Id
		if a.isOrder(txnTypeId) || a.isFBAInventoryFee(txnTypeId) || a.isServiceFee(txnTypeId) {
			amazonCosts = append(amazonCosts, types.Summary{
				SKU:         txn.SKU,
				Description: txn.Description,
				Marketplace: *a.marketplace(txn.Marketplace.Id),
				AmazonCosts: a.txnAmazonCosts(txn) * a.exchangeRate(txn.Marketplace.Id),
				Total:       a.txnAmazonCosts(txn) * a.exchangeRate(txn.Marketplace.Id),
				Date:        txn.DateTime,
			})
		}
	}

	return amazonCosts
}
