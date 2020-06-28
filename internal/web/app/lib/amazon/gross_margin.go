package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

func (a *Amazon) GrossMargin(txns *[]models.Transaction) []types.Summary {
	var grossMargin []types.Summary

	for _, txn := range *txns {
		txnTypeId := txn.TransactionType.Id
		if a.isOrder(txnTypeId) {
			grossMargin = append(grossMargin, types.Summary{
				SKU:                txn.SKU,
				Description:        txn.Description,
				Marketplace:        *a.marketplace(txn.Marketplace.Id),
				ProductCosts:       a.txnProductCosts(txn),
				QuantitySold:       txn.Quantity,
				TotalRevenue:       a.txnProductSales(txn) * a.exchangeRate(txn.Marketplace.Id),
				AmazonCosts:        a.txnAmazonCosts(txn) * a.exchangeRate(txn.Marketplace.Id),
				ShippingCredits:    a.txnShippingCredits(txn) * a.exchangeRate(txn.Marketplace.Id),
				PromotionalRebates: a.txnPromotionalRebates(txn) * a.exchangeRate(txn.Marketplace.Id),
				GrossMargin:        a.txnGrossMargin(txn) * a.exchangeRate(txn.Marketplace.Id),
				SalesTaxCollected:  a.txnSalesTaxCollected(txn) * a.exchangeRate(txn.Marketplace.Id),
				Total:              (a.txnGrossMargin(txn) + a.txnSalesTaxCollected(txn)) * a.exchangeRate(txn.Marketplace.Id),
				Date:               txn.DateTime,
			})
		}
	}

	return grossMargin
}
