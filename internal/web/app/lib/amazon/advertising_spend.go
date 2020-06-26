package amazon

import (
	"regexp"

	"gitlab.com/getlytica/lytica-app/internal/models"
	"gitlab.com/getlytica/lytica-app/internal/web/app/types"
)

var (
	advertisingSpendLineItem = regexp.MustCompile(`^Cost of Advertising+$`).MatchString
)

func (a *Amazon) AdvertisingSpend(txns *[]models.Transaction) []types.Summary {
	var advertisingSpend []types.Summary

	for _, txn := range *txns {
		txnTypeId := txn.TransactionType.Id
		if a.isServiceFee(txnTypeId) && advertisingSpendLineItem(txn.Description) {
			advertisingSpend = append(advertisingSpend, types.Summary{
				Total:       a.txnAdvertisingCosts(txn) * a.exchangeRate(txn.Marketplace.Id),
				Marketplace: *a.marketplace(txn.Marketplace.Id),
				OrderDate:   txn.DateTime,
			})
		}
	}

	return advertisingSpend
}
