package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (a *Amazon) TotalCosts(txns *[]models.Transaction) []types.Summary {
	var totalCosts []types.Summary

	for _, txn := range *txns {
		var cost float64

		exchangeRate := a.exchangeRate(txn.Marketplace.Id)

		txnTypeId := txn.TransactionType.Id
		if a.isOrder(txnTypeId) {
			cost += a.txnTotalOrderCosts(txn)
		}

		cost += a.txnAdvertisingCosts(txn)

		if a.isRefund(txnTypeId) {
			cost += txn.Total * exchangeRate
		}

		totalCosts = append(totalCosts, types.Summary{
			Total:       cost * exchangeRate,
			Marketplace: *a.marketplace(txn.Marketplace.Id),
			OrderDate:   txn.DateTime,
		})
	}

	return totalCosts
}
