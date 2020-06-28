package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

func (a *Amazon) TotalCosts(txns *[]models.Transaction) []types.Summary {
	var totalCosts []types.Summary

	for _, txn := range *txns {
		var cost float64

		exchangeRate := a.exchangeRate(txn.Marketplace.Id)

		txnTypeId := txn.TransactionType.Id
		if a.isOrder(txnTypeId) {
			cost += a.txnTotalCosts(txn)
		}

		cost += a.txnAdvertisingCosts(txn)

		if a.isRefund(txnTypeId) {
			cost += txn.Total * exchangeRate
		}

		totalCosts = append(totalCosts, types.Summary{
			SKU:                  txn.SKU,
			Description:          txn.Description,
			Marketplace:          *a.marketplace(txn.Marketplace.Id),
			AmazonCosts:          0.0,
			ProductCosts:         0.0,
			ProductCostsUnit:     0.0,
			Total:                cost * exchangeRate,
			TotalCostsPercentage: 0.0,
			Date:                 txn.DateTime,
		})
	}

	return totalCosts
}
