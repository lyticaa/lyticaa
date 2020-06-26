package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/core/app/types"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (a *Amazon) PromotionalRebates(txns *[]models.Transaction) []types.Summary {
	var promotionalRebates []types.Summary

	for _, txn := range *txns {
		promotionalRebates = append(promotionalRebates, types.Summary{
			Total:       a.txnPromotionalRebates(txn) * a.exchangeRate(txn.Marketplace.Id),
			Marketplace: *a.marketplace(txn.Marketplace.Id),
			OrderDate:   txn.DateTime,
		})
	}

	return promotionalRebates
}
