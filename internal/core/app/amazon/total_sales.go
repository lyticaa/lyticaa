package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (a *Amazon) TotalSales(txns *[]models.Transaction) []types.Summary {
	var totalSales []types.Summary

	for _, txn := range *txns {
		if a.isOrder(txn.TransactionType.Id) {
			totalSales = append(totalSales,
				types.Summary{
					SKU:         txn.SKU,
					Description: txn.Description,
					Marketplace: *a.marketplace(txn.Marketplace.Id),
					Total:       a.txnProductSales(txn) * a.exchangeRate(txn.Marketplace.Id),
					OrderDate:   txn.DateTime,
				},
			)
		}
	}

	return totalSales
}
