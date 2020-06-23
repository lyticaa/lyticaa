package models

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Summary struct {
	UserId      int64     `db:"user_id"`
	Total       float64   `db:"total"`
	Marketplace string    `db:"marketplace"`
	OrderDate   time.Time `db:"order_date"`
}

type Transaction struct {
	Id                     int64
	User                   User `db:"user_id"`
	DateTime               time.Time
	SettlementId           int64
	SettlementIdx          int64
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
	CreatedAt              time.Time `db:"created_at"`
	UpdatedAt              time.Time `db:"updated_at"`
}

func LoadSummary(userId int64, view, dateRange string, db *sqlx.DB) *[]Summary {
	var summary []Summary

	query := fmt.Sprintf(`SELECT user_id, total, order_date, marketplace FROM %v_%v WHERE user_id = $1`, view, dateRange)
	err := db.Select(
		&summary,
		query,
		userId,
	)

	if err != nil {
		logger().Error().Err(err).Msgf("failed to load the summary for the user %v", userId)
		return &[]Summary{}
	}

	return &summary
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
		"sku":                      t.Sku,
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
