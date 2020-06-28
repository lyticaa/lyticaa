package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

func (a *Amazon) PromotionalRebates(txns *[]models.Transaction) []types.Summary {
	var promotionalRebates []types.Summary

	for _, txn := range *txns {
		promotionalRebates = append(promotionalRebates, types.Summary{
			SKU:                   txn.SKU,
			Description:           txn.Description,
			Marketplace:           *a.marketplace(txn.Marketplace.Id),
			PromotionalRebates:    txn.PromotionalRebates * a.exchangeRate(txn.Marketplace.Id),
			PromotionalRebatesTax: txn.PromotionalRebatesTax * a.exchangeRate(txn.Marketplace.Id),
			Total:                 a.txnPromotionalRebates(txn) * a.exchangeRate(txn.Marketplace.Id),
			Date:                  txn.DateTime,
		})
	}

	return promotionalRebates
}
