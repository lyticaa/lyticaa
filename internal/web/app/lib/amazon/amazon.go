package amazon

import (
	"gitlab.com/getlytica/lytica-app/internal/models"

	"github.com/jmoiron/sqlx"
)

const (
	typeOrder           = "Order"
	typeRefund          = "Refund"
	typeServiceFee      = "Service Fee"
	typeFBAInventoryFee = "FBA Inventory Fee"
	defaultExchangeRate = 1.0
)

type Amazon struct {
	db               *sqlx.DB
	transactionTypes *[]models.TransactionType
	exchangeRates    *[]models.ExchangeRate
}

func NewAmazon(db *sqlx.DB) *Amazon {
	return &Amazon{
		db:               db,
		transactionTypes: models.LoadTransactionTypes(db),
		exchangeRates:    models.LoadExchangeRates(db),
	}
}

func (a *Amazon) LoadTransactions(userId int64, dateRange string) *[]models.Transaction {
	return models.LoadTransaction(userId, dateRange, a.db)
}

func (a *Amazon) LoadSponsoredProducts(userId int64, dateRange string) *[]models.SponsoredProduct {
	return models.LoadSponsoredProducts(userId, dateRange, a.db)
}

func (a *Amazon) marketplace(marketplaceId int64) *string {
	marketplaces := models.LoadMarketplaces(a.db)
	for _, m := range *marketplaces {
		if m.Id == marketplaceId {
			return &m.Name
		}
	}

	return nil
}

func (a *Amazon) isOrder(txnId int64) bool {
	isOrder := false
	for _, tt := range *a.transactionTypes {
		if tt.Id == txnId && tt.Name == typeOrder {
			isOrder = true
		}
	}

	return isOrder
}

func (a *Amazon) isRefund(txnId int64) bool {
	isRefund := false
	for _, tt := range *a.transactionTypes {
		if tt.Id == txnId && tt.Name == typeRefund {
			isRefund = true
		}
	}

	return isRefund
}

func (a *Amazon) isServiceFee(txnId int64) bool {
	isServiceFee := false
	for _, tt := range *a.transactionTypes {
		if tt.Id == txnId && tt.Name == typeServiceFee {
			isServiceFee = true
		}
	}

	return isServiceFee
}

func (a *Amazon) isFBAInventoryFee(txnId int64) bool {
	isServiceFee := false
	for _, tt := range *a.transactionTypes {
		if tt.Id == txnId && tt.Name == typeFBAInventoryFee {
			isServiceFee = true
		}
	}

	return isServiceFee
}

func (a *Amazon) exchangeRate(marketplaceId int64) float64 {
	for _, rate := range *a.exchangeRates {
		if rate.MarketplaceId == marketplaceId {
			return rate.Rate
		}
	}

	return defaultExchangeRate
}

func (a *Amazon) txnProductSales(txn models.Transaction) float64 {
	return txn.ProductSales + txn.ProductSalesTax
}

func (a *Amazon) txnUnitsSold(txn models.Transaction) float64 {
	return float64(txn.Quantity)
}

func (a *Amazon) txnAmazonCosts(txn models.Transaction) float64 {
	return txn.SellingFees + txn.FBAFees + txn.Other
}

func (a *Amazon) txnProductCosts(txn models.Transaction) float64 {
	return float64(txn.Quantity) * a.costOfGoods(txn.User.Id, txn.SKU, txn.DateTime)
}

func (a *Amazon) txnAdvertisingCosts(txn models.Transaction) float64 {
	var cost float64
	if a.isServiceFee(txn.TransactionType.Id) {
		if advertisingSpendLineItem(txn.Description) {
			cost = txn.OtherTransactionFees
		}
	}

	return cost
}

func (a *Amazon) txnShippingCredits(txn models.Transaction) float64 {
	return txn.ShippingCredits + txn.ShippingCreditsTax
}

func (a *Amazon) txnPromotionalRebates(txn models.Transaction) float64 {
	return txn.PromotionalRebates + txn.PromotionalRebatesTax
}

func (a *Amazon) txnTotalOrderCosts(txn models.Transaction) float64 {
	return a.txnAmazonCosts(txn) + a.txnProductCosts(txn) + a.txnShippingCredits(txn) + a.txnPromotionalRebates(txn)
}

func (a *Amazon) txnGrossMargin(txn models.Transaction) float64 {
	return a.txnProductSales(txn) + a.txnShippingCredits(txn) + a.txnPromotionalRebates(txn) + a.txnAmazonCosts(txn)
}

func (a *Amazon) txnNetMargin(txn models.Transaction) float64 {
	return a.txnGrossMargin(txn) - a.txnProductCosts(txn)
}
