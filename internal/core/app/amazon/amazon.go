package amazon

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/getlytica/lytica-app/internal/models"
)

const (
	typeOrder           = "Order"
	typeServiceFee      = "Service Fee"
	typeFBAInventoryFee = "FBA Inventory Fee"
	defaultExchangeRate = 1.0
)

type Amazon struct {
	db               *sqlx.DB
	transactionTypes []models.TransactionType
	exchangeRates    []models.ExchangeRate
}

func NewAmazon(db *sqlx.DB) *Amazon {
	return &Amazon{
		db:               db,
		transactionTypes: models.LoadTransactionTypes(db),
		exchangeRates:    models.LoadExchangeRates(db),
	}
}

func (a *Amazon) LoadTransactions(userId int64, dateRange string) *[]models.Transaction {
	return models.LoadTransactionsByDateRange(userId, dateRange, a.db)
}

func (a *Amazon) marketplace(marketplaceId int64) *string {
	mp := models.LoadMarketplaces(a.db)
	for _, m := range mp {
		if m.Id == marketplaceId {
			return &m.Name
		}
	}

	return nil
}

func (a *Amazon) isOrder(txnId int64) bool {
	isOrder := false
	for _, tt := range a.transactionTypes {
		if tt.Id == txnId && tt.Name == typeOrder {
			isOrder = true
		}
	}

	return isOrder
}

func (a *Amazon) isServiceFee(txnId int64) bool {
	isServiceFee := false
	for _, tt := range a.transactionTypes {
		if tt.Id == txnId && tt.Name == typeServiceFee {
			isServiceFee = true
		}
	}

	return isServiceFee
}

func (a *Amazon) isFBAInventoryFee(txnId int64) bool {
	isServiceFee := false
	for _, tt := range a.transactionTypes {
		if tt.Id == txnId && tt.Name == typeFBAInventoryFee {
			isServiceFee = true
		}
	}

	return isServiceFee
}

func (a *Amazon) exchangeRate(marketplaceId int64) float64 {
	for _, rate := range a.exchangeRates {
		if rate.MarketplaceId == marketplaceId {
			return rate.Rate
		}
	}

	return defaultExchangeRate
}
