package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Transaction struct {
	Id                     int64              `db:"id"`
	User                   User               `db:"user_id"`
	DateTime               time.Time          `db:"date_time"`
	SettlementId           int64              `db:"settlement_id"`
	SettlementIdx          int64              `db:"settlement_idx"`
	TransactionType        TransactionType    `db:"transaction_type_id"`
	OrderId                string             `db:"order_id"`
	SKU                    string             `db:"sku"`
	Description            string             `db:"description"`
	Quantity               int64              `db:"quantity"`
	Marketplace            Marketplace        `db:"marketplace_id"`
	Fulfillment            Fulfillment        `db:"fulfillment_od"`
	TaxCollectionModel     TaxCollectionModel `dn:"tax_collection_model_id"`
	ProductSales           float64            `db:"product_sales"`
	ProductSalesTax        float64            `db:"product_sales_tax"`
	ShippingCredits        float64            `db:"shipping_credits"`
	ShippingCreditsTax     float64            `db:"shipping_credits_tax"`
	GiftwrapCredits        float64            `db:"giftwrap_credits"`
	GiftwrapCreditsTax     float64            `db:"giftwrap_credits_tax"`
	PromotionalRebates     float64            `db:"promotional_rebates"`
	PromotionalRebatesTax  float64            `db:"promotional_rebates_tax"`
	MarketplaceWithheldTax float64            `db:"marketplace_withheld_tax"`
	SellingFees            float64            `db:"selling_fees"`
	FBAFees                float64            `db:"fba_fees"`
	OtherTransactionFees   float64            `db:"other_transaction_fees"`
	Other                  float64            `db:"other"`
	Total                  float64            `db:"total"`
	CreatedAt              time.Time          `db:"created_at"`
	UpdatedAt              time.Time          `db:"updated_at"`
}

func LoadTransaction(userId int64, dateRange string, db *sqlx.DB) *[]Transaction {
	var transactions []Transaction

	query := `SELECT t.* FROM transactions_%v t WHERE t.user_id = $1`
	_ = db.Select(&transactions, fmt.Sprintf(query, dateRange), userId)

	return &transactions
}

func TotalTransactions(userId int64, dateRange string, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(*) FROM transactions_%v WHERE user_id = $1`
	_ = db.QueryRow(fmt.Sprintf(query, dateRange), userId).Scan(&count)

	return count
}

func TotalRefundTransactions(userId int64, dateRange string, db *sqlx.DB) int64 {
	var count int64

	query := `SELECT COUNT(t.*) FROM transactions_%v AS t LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id WHERE t.user_id = $1 AND tt.name = 'Refund'`
	_ = db.QueryRow(fmt.Sprintf(query, dateRange), userId).Scan(&count)

	return count
}

func (t *Transaction) Save(db *sqlx.DB) error {
	query := `INSERT INTO transactions (
                          user_id,
                          date_time,
                          settlement_id,
                          settlement_idx,
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
                                :user_id,
                                :date_time,
                                :settlement_id,
                                :settlement_idx,
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
                                :total)
						ON CONFLICT (user_id, date_time, settlement_id, settlement_idx, transaction_type_id, order_id, sku)
							DO UPDATE SET quantity = :quantity,
							              marketplace_id = :marketplace_id,
							              fulfillment_id = :fulfillment_id,
							              tax_collection_model_id = :tax_collection_model_id,
							              product_sales = :product_sales,
							              product_sales_tax = :product_sales_tax,
							              shipping_credits = :shipping_credits,
							              shipping_credits_tax = :shipping_credits_tax,
							              giftwrap_credits = :giftwrap_credits,
							              giftwrap_credits_tax = :giftwrap_credits_tax,
							              promotional_rebates = :promotional_rebates,
							              promotional_rebates_tax = :promotional_rebates_tax,
							              marketplace_withheld_tax = :marketplace_withheld_tax,
							              selling_fees = :selling_fees,
							              fba_fees = :fba_fees,
							              other_transaction_fees = :other_transaction_fees,
							              other = :other,
							              total = :total,
							              updated_at = NOW()`
	_, err := db.NamedExec(query, map[string]interface{}{
		"user_id":                  t.User.Id,
		"date_time":                t.DateTime,
		"settlement_id":            t.SettlementId,
		"settlement_idx":           t.SettlementIdx,
		"transaction_type_id":      t.TransactionType.Id,
		"order_id":                 t.OrderId,
		"sku":                      t.SKU,
		"quantity":                 t.Quantity,
		"marketplace_id":           t.Marketplace.Id,
		"fulfillment_id":           t.Fulfillment.Id,
		"tax_collection_model_id":  t.TaxCollectionModel.Id,
		"product_sales":            t.ProductSales,
		"product_sales_tax":        t.ProductSalesTax,
		"shipping_credits":         t.ShippingCredits,
		"shipping_credits_tax":     t.ShippingCreditsTax,
		"giftwrap_credits":         t.GiftwrapCredits,
		"giftwrap_credits_tax":     t.GiftwrapCreditsTax,
		"promotional_rebates":      t.PromotionalRebates,
		"promotional_rebates_tax":  t.PromotionalRebatesTax,
		"marketplace_withheld_tax": t.MarketplaceWithheldTax,
		"selling_fees":             t.SellingFees,
		"fba_fees":                 t.FBAFees,
		"other_transaction_fees":   t.OtherTransactionFees,
		"other":                    t.Other,
		"total":                    t.Total,
	})

	if err != nil {
		return err
	}

	return nil
}
