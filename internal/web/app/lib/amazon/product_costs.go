package amazon

import (
	"time"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

const (
	defaultCostOfGood = 0.0
)

// ((cost of product) * quantity sold) + ((cost of product) * quantity sold as coupons) + SKU ads + refunds

func (a *Amazon) ProductCosts(txns *[]models.Transaction) []types.Summary {
	var productCosts []types.Summary

	for _, txn := range *txns {
		txnTypeId := txn.TransactionType.Id
		if a.isOrder(txnTypeId) {
			productCosts = append(productCosts, types.Summary{
				SKU:              txn.SKU,
				Description:      txn.Description,
				Marketplace:      *a.marketplace(txn.Marketplace.Id),
				QuantitySold:     txn.Quantity,
				ProductCosts:     a.txnProductCosts(txn) * a.exchangeRate(txn.Marketplace.Id),
				AdvertisingSpend: 0.0,
				Refunds:          0.0,
				Total:            a.txnProductCosts(txn) * a.exchangeRate(txn.Marketplace.Id),
				Date:             txn.DateTime,
			})

		}
	}

	return productCosts
}

func (a *Amazon) costOfGoods(userId int64, sku string, orderDate time.Time) float64 {
	costOfGoods := models.LoadCostOfGood(userId, sku, a.db)
	for _, cost := range *costOfGoods {
		if orderDate.After(cost.StartAt) && orderDate.Before(cost.EndAt) {
			return cost.Cost
		}
	}

	return defaultCostOfGood
}
