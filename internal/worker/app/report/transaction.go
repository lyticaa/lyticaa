package report

import (
	"strconv"
	"strings"
	"time"

	"gitlab.com/getlytica/lytica-app/internal/models"
)

func (r *Report) getTransactionTypes() []models.TransactionType {
	return models.GetTransactionTypes(r.Db)
}

func (r *Report) getTransactionTypeIdByName(name string, txnTypes []models.TransactionType) (int64, bool) {
	for _, txnType := range txnTypes {
		if txnType.Name == name {
			return txnType.Id, true
		}
	}

	return unknown, false
}

func (r *Report) getMarketplaces() []models.Marketplace {
	return models.GetMarketplaces(r.Db)
}

func (r *Report) getMarketplaceIdByName(name string, marketplaces []models.Marketplace) (int64, bool) {
	for _, marketplace := range marketplaces {
		if marketplace.Name == name {
			return marketplace.Id, true
		}
	}

	return unknown, false
}

func (r *Report) getFulfillments() []models.Fulfillment {
	return models.GetFulfillments(r.Db)
}

func (r *Report) getFulfillmentIdByName(name string, fulfillments []models.Fulfillment) (int64, bool) {
	for _, fulfillment := range fulfillments {
		if fulfillment.Name == name {
			return fulfillment.Id, true
		}
	}

	return unknown, false
}

func (r *Report) getTaxCollectionModels() []models.TaxCollectionModel {
	return models.GetTaxCollectionModels(r.Db)
}

func (r *Report) getTaxCollectionModelIdByName(name string, taxCollectionModels []models.TaxCollectionModel) (int64, bool) {
	for _, taxCollectionModel := range taxCollectionModels {
		if taxCollectionModel.Name == name {
			return taxCollectionModel.Id, true
		}
	}

	return unknown, false
}

func (r *Report) formatTransactions(rows []map[string]string, username string) []models.Transaction {
	user := models.FindUser(username, r.Db)
	txnTypes := r.getTransactionTypes()
	marketplaces := r.getMarketplaces()
	fulfillments := r.getFulfillments()
	taxCollectionModels := r.getTaxCollectionModels()

	var txns []models.Transaction
	settlementIdx := int64(1)

	for idx, row := range rows {
		if idx > 0 {
			if row["settlement id"] != rows[idx-1]["settlement id"] {
				settlementIdx = 1
			} else {
				settlementIdx++
			}
		}

		txnType, ok := r.getTransactionTypeIdByName(row["type"], txnTypes)
		if !ok && row["type"] != "" {
			r.Logger.Error().Msgf("Transaction Type %v not found", row["type"])
		}

		marketplace, ok := r.getMarketplaceIdByName(strings.ToLower(row["marketplace"]), marketplaces)
		if !ok && row["marketplace"] != "" {
			r.Logger.Error().Msgf("Marketplace %v not found", row["marketplace"])
		}

		fulfillment, ok := r.getFulfillmentIdByName(row["fulfillment"], fulfillments)
		if !ok && row["fulfillment"] != "" {
			r.Logger.Error().Msgf("Fulfillment %v not found", row["fulfillment"])
		}

		taxCollectionModel, ok := r.getTaxCollectionModelIdByName(row["tax collection model"], taxCollectionModels)
		if !ok && row["tax collection model"] != "" {
			r.Logger.Error().Msgf("Tax Collection Model %v not found", row["tax collection model"])
		}

		dateTime, _ := time.Parse("Jan 2, 2006 3:04:05 PM MST", row["date/time"])
		settlementId, _ := strconv.ParseInt(row["settlement id"], 10, 64)
		quantity, _ := strconv.ParseInt(row["quantity"], 10, 64)
		productSales, _ := strconv.ParseFloat(row["product sales"], 64)
		productSalesTax, _ := strconv.ParseFloat(row["product sales tax"], 64)
		shippingCredits, _ := strconv.ParseFloat(row["shipping credits"], 64)
		shippingCreditsTax, _ := strconv.ParseFloat(row["shipping credits tax"], 64)
		giftwrapCredits, _ := strconv.ParseFloat(row["gift wrap credits"], 64)
		giftwrapCreditsTax, _ := strconv.ParseFloat(row["giftwrap credits tax"], 64)
		promotionalRebates, _ := strconv.ParseFloat(row["promotional rebates"], 64)
		promotionalRebatesTax, _ := strconv.ParseFloat(row["promotional rebates tax"], 64)
		marketplaceWithheldTax, _ := strconv.ParseFloat(row["marketplace withheld tax"], 64)
		sellingFees, _ := strconv.ParseFloat(row["selling fees"], 64)
		fbaFees, _ := strconv.ParseFloat(row["fba fees"], 64)
		otherTransactionFees, _ := strconv.ParseFloat(row["other transaction fees"], 64)
		other, _ := strconv.ParseFloat(row["other"], 64)
		total, _ := strconv.ParseFloat(row["total"], 64)

		txn := models.Transaction{
			DateTime:               dateTime,
			User:                   *user,
			SettlementId:           settlementId,
			SettlementIdx:          settlementIdx,
			TransactionType:        models.TransactionType{Id: txnType},
			OrderId:                row["order id"],
			Sku:                    row["sku"],
			Quantity:               quantity,
			Marketplace:            models.Marketplace{Id: marketplace},
			Fulfillment:            models.Fulfillment{Id: fulfillment},
			TaxCollectionModel:     models.TaxCollectionModel{Id: taxCollectionModel},
			ProductSales:           productSales,
			ProductSalesTax:        productSalesTax,
			ShippingCredits:        shippingCredits,
			ShippingCreditsTax:     shippingCreditsTax,
			GiftwrapCredits:        giftwrapCredits,
			GiftwrapCreditsTax:     giftwrapCreditsTax,
			PromotionalRebates:     promotionalRebates,
			PromotionalRebatesTax:  promotionalRebatesTax,
			MarketplaceWithheldTax: marketplaceWithheldTax,
			SellingFees:            sellingFees,
			FBAFees:                fbaFees,
			OtherTransactionFees:   otherTransactionFees,
			Other:                  other,
			Total:                  total,
		}

		txns = append(txns, txn)
	}

	return txns
}

func (r *Report) saveTransaction(txn models.Transaction) error {
	err := models.SaveTransaction(txn, r.Db)
	if err != nil {
		return err
	}

	return nil
}
