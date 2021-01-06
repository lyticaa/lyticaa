package custom_transactions

import (
	"context"

	"github.com/lyticaa/lyticaa/internal/models"
)

const (
	unknown = int64(1)
)

func (c *CustomTransactions) Process(rows []map[string]string, userID string) []error {
	var errors []error

	formatted := c.Format(rows, userID)
	for _, item := range formatted {
		if err := c.Create(item); err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func (c *CustomTransactions) Format(rows []map[string]string, userID string) []models.AmazonCustomTransactionModel {
	var transactions []models.AmazonCustomTransactionModel
	for _, row := range rows {
		transactions = append(transactions, models.AmazonCustomTransactionModel{
			DateTime:                   c.dateTime(row),
			UserID:                     userID,
			SettlementID:               c.settlementID(row),
			AmazonTransactionTypeID:    c.amazonTransactionType(row),
			OrderID:                    c.orderID(row),
			SKU:                        c.sku(row),
			Quantity:                   c.quantity(row),
			AmazonMarketplaceID:        c.amazonMarketplace(row),
			AmazonFulfillmentID:        c.amazonFulfillment(row),
			AmazonTaxCollectionModelID: c.amazonTaxCollectionModel(row),
			ProductSales:               c.productSales(row),
			ProductSalesTax:            c.productSalesTax(row),
			ShippingCredits:            c.shippingCredits(row),
			ShippingCreditsTax:         c.shippingCreditsTax(row),
			GiftwrapCredits:            c.giftwrapCredits(row),
			GiftwrapCreditsTax:         c.giftwrapCreditsTax(row),
			PromotionalRebates:         c.promotionalRebates(row),
			PromotionalRebatesTax:      c.promotionalRebatesTax(row),
			MarketplaceWithheldTax:     c.marketplaceWithheldTax(row),
			SellingFees:                c.sellingFees(row),
			FBAFees:                    c.fbaFees(row),
			OtherTransactionFees:       c.otherTransactionFees(row),
			Other:                      c.other(row),
			Total:                      c.total(row),
		})
	}

	return transactions
}

func (c *CustomTransactions) Create(transaction models.AmazonCustomTransactionModel) error {
	if err := transaction.Create(context.TODO(), c.db); err != nil {
		return err
	}

	return nil
}

func (c *CustomTransactions) AmazonTransactionTypes() *[]models.AmazonTransactionTypeModel {
	var amazonTransactionTypeModel models.AmazonTransactionTypeModel
	amazonTransactionTypes := amazonTransactionTypeModel.FetchAll(context.TODO(), nil, nil, c.db).([]models.AmazonTransactionTypeModel)

	return &amazonTransactionTypes
}

func (c *CustomTransactions) amazonTransactionTypeIDByName(name string) (int64, bool) {
	txnTypes := *c.AmazonTransactionTypes()
	for _, txnType := range txnTypes {
		if txnType.Name == name {
			return txnType.ID, true
		}
	}

	return unknown, false
}

func (c *CustomTransactions) AmazonMarketplaces() *[]models.AmazonMarketplaceModel {
	var amazonMarketplaceModel models.AmazonMarketplaceModel
	amazonMarketplaces := amazonMarketplaceModel.FetchAll(context.TODO(), nil, nil, c.db).([]models.AmazonMarketplaceModel)

	return &amazonMarketplaces
}

func (c *CustomTransactions) amazonMarketplaceIDByName(name string) (int64, bool) {
	marketplaces := *c.AmazonMarketplaces()
	for _, marketplace := range marketplaces {
		if marketplace.Name == name {
			return marketplace.ID, true
		}
	}

	return unknown, false
}

func (c *CustomTransactions) AmazonFulfillments() *[]models.AmazonFulfillmentModel {
	var amazonFulfillmentModel models.AmazonFulfillmentModel
	amazonFulfillments := amazonFulfillmentModel.FetchAll(context.TODO(), nil, nil, c.db).([]models.AmazonFulfillmentModel)

	return &amazonFulfillments
}

func (c *CustomTransactions) amazonFulfillmentIDByName(name string) (int64, bool) {
	fulfillments := *c.AmazonFulfillments()
	for _, fulfillment := range fulfillments {
		if fulfillment.Name == name {
			return fulfillment.ID, true
		}
	}

	return unknown, false
}

func (c *CustomTransactions) AmazonTaxCollectionModels() *[]models.AmazonTaxCollectionModelModel {
	var amazonTaxCollectionModelModel models.AmazonTaxCollectionModelModel
	amazonTaxCollectionModels := amazonTaxCollectionModelModel.FetchAll(context.TODO(), nil, nil, c.db).([]models.AmazonTaxCollectionModelModel)

	return &amazonTaxCollectionModels
}

func (c *CustomTransactions) amazonTaxCollectionModelIDByName(name string) (int64, bool) {
	taxCollectionModels := *c.AmazonTaxCollectionModels()
	for _, taxCollectionModel := range taxCollectionModels {
		if taxCollectionModel.Name == name {
			return taxCollectionModel.ID, true
		}
	}

	return unknown, false
}
