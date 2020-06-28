package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

func (a *Amazon) NetMargin(txns *[]models.Transaction) []types.Summary {
	var netMargin []types.Summary

	for _, txn := range *txns {
		netMargin = append(netMargin, types.Summary{
			SKU:           txn.SKU,
			Description:   txn.Description,
			Marketplace:   *a.marketplace(txn.Marketplace.Id),
			GrossMargin:   a.txnGrossMargin(txn) * a.exchangeRate(txn.Marketplace.Id),
			TotalCosts:    a.txnTotalCosts(txn) * a.exchangeRate(txn.Marketplace.Id),
			NetMargin:     a.txnNetMargin(txn) * a.exchangeRate(txn.Marketplace.Id),
			QuantitySold:  txn.Quantity,
			NetMarginUnit: (a.txnNetMargin(txn) * a.exchangeRate(txn.Marketplace.Id)) / float64(txn.Quantity),
			ROI:           0.0,
			Total:         a.txnNetMargin(txn) * a.exchangeRate(txn.Marketplace.Id),
			Date:          txn.DateTime,
		})
	}

	return netMargin
}
