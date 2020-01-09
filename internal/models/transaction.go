package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Transaction struct {
	Id                     int64
	Idx                    int64
	User                   User `db:"user_id"`
	DateTime               time.Time
	SettlementId           int64
	TransactionType        TransactionType `db:"transaction_type_id"`
	OrderId                string
	Sku                    string
	Quantity               int64
	Marketplace            Marketplace        `db:"marketplace_id"`
	Fulfillment            Fulfillment        `db:"fulfillment_od"`
	TaxCollectionModel     TaxCollectionModel `dn:"tax_collection_model_id"`
	ProductSales           float64
	ProductSalesTax        float64
	ShippingCredits        float64
	ShippingCreditsTax     float64
	GiftwrapCredits        float64
	GiftwrapCreditsTax     float64
	PromotionalRebates     float64
	PromotionalRebatesTax  float64
	MarketplaceWithheldTax float64
	SellingFees            float64
	FBAFees                float64
	OtherTransactionFees   float64
	Other                  float64
	Total                  float64
}

func SaveTransaction(txn Transaction, db *sqlx.DB) {
	query := `INSERT INTO transactions (
                idx,
				user_id,
				date_time,
				settlement_id,
				transaction_type_id,
				order_id,
				sku,
				quantity,
				marketplace_id,
				fulfillment_id,
				tax_collection_model_id,
				product_sales,
				product_sales_tax,
				shipping_credits,
				shipping_credits_tax,
				giftwrap_credits,
				giftwrap_credits_tax,
				promotional_rebates,
				promotional_rebates_tax,
				marketplace_withheld_tax,
				selling_fees,
				fba_fees,
				other_transaction_fees,
				other,
				total)
			VALUES (
			    :idx,
				:user_id,
				:date_time,
				:settlement_id,
				:transaction_type_id,
				:order_id,
				:sku,
				:quantity,
				:marketplace_id,
				:fulfillment_id,
				:tax_collection_model_id,
				:product_sales,
				:product_sales_tax,
				:shipping_credits,
				:shipping_credits_tax,
				:giftwrap_credits,
				:giftwrap_credits_tax,
				:promotional_rebates,
				:promotional_rebates_tax,
				:marketplace_withheld_tax,
				:selling_fees,
				:fba_fees,
				:other_transaction_fees,
				:other,
				:total)`

	_, err := db.NamedExec(query, map[string]interface{}{
		"idx":                      txn.Idx,
		"user_id":                  txn.User.Id,
		"date_time":                txn.DateTime,
		"settlement_id":            txn.SettlementId,
		"transaction_type_id":      txn.TransactionType.Id,
		"order_id":                 txn.OrderId,
		"sku":                      txn.Sku,
		"quantity":                 txn.Quantity,
		"marketplace_id":           txn.Marketplace.Id,
		"fulfillment_id":           txn.Fulfillment.Id,
		"tax_collection_model_id":  txn.TaxCollectionModel.Id,
		"product_sales":            txn.ProductSales,
		"product_sales_tax":        txn.ProductSalesTax,
		"shipping_credits":         txn.ShippingCredits,
		"shipping_credits_tax":     txn.ShippingCreditsTax,
		"giftwrap_credits":         txn.GiftwrapCredits,
		"giftwrap_credits_tax":     txn.GiftwrapCreditsTax,
		"promotional_rebates":      txn.PromotionalRebates,
		"promotional_rebates_tax":  txn.PromotionalRebatesTax,
		"marketplace_withheld_tax": txn.MarketplaceWithheldTax,
		"selling_fees":             txn.SellingFees,
		"fba_fees":                 txn.FBAFees,
		"other_transaction_fees":   txn.OtherTransactionFees,
		"other":                    txn.Other,
		"total":                    txn.Total,
	})

	if err != nil {
		fmt.Printf("%v\n", txn)
		panic(err)
	}
}
