package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type AmazonCustomTransactionModel struct {
	ID                         int64     `db:"id"`
	AmazonCustomTransactionID  string    `db:"amazon_custom_transaction_id"`
	UserID                     string    `db:"user_id"`
	DateTime                   time.Time `db:"date_time"`
	DateRange                  string    `db:"date_range"`
	SettlementID               int64     `db:"settlement_id"`
	AmazonTransactionType      string    `db:"amazon_transaction_type"`
	AmazonTransactionTypeID    int64     `db:"amazon_transaction_type_id"`
	OrderID                    string    `db:"order_id"`
	SKU                        string    `db:"sku"`
	Quantity                   int64     `db:"quantity"`
	AmazonMarketplace          string    `db:"amazon_marketplace"`
	AmazonMarketplaceID        int64     `db:"amazon_marketplace_id"`
	AmazonFulfillmentID        int64     `db:"amazon_fulfillment_id"`
	AmazonTaxCollectionModelID int64     `dn:"amazon_tax_collection_model_id"`
	ProductSales               float64   `db:"product_sales"`
	ProductSalesTax            float64   `db:"product_sales_tax"`
	ShippingCredits            float64   `db:"shipping_credits"`
	ShippingCreditsTax         float64   `db:"shipping_credits_tax"`
	GiftwrapCredits            float64   `db:"giftwrap_credits"`
	GiftwrapCreditsTax         float64   `db:"giftwrap_credits_tax"`
	PromotionalRebates         float64   `db:"promotional_rebates"`
	PromotionalRebatesTax      float64   `db:"promotional_rebates_tax"`
	MarketplaceWithheldTax     float64   `db:"marketplace_withheld_tax"`
	SellingFees                float64   `db:"selling_fees"`
	FBAFees                    float64   `db:"fba_fees"`
	OtherTransactionFees       float64   `db:"other_transaction_fees"`
	Other                      float64   `db:"other"`
	Total                      float64   `db:"total"`
	CreatedAt                  time.Time `db:"created_at"`
	UpdatedAt                  time.Time `db:"updated_at"`
}

func (ac *AmazonCustomTransactionModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	return nil
}
func (ac *AmazonCustomTransactionModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{} {
	return nil
}
func (ac *AmazonCustomTransactionModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	return nil
}
func (ac *AmazonCustomTransactionModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	return int64(0)
}

func (ac *AmazonCustomTransactionModel) Create(ctx context.Context, db *sqlx.DB) error {
	query := `INSERT INTO amazon_custom_transactions (
                                        user_id,
                                        date_time,
                                        settlement_id,
                                        amazon_transaction_type_id,
                                        order_id,
                                        sku,
                                        quantity,
                                        amazon_marketplace_id,
                                        amazon_fulfillment_id,
                                        amazon_tax_collection_model_id,
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
                                                :amazon_transaction_type_id,
                                                :order_id,
                                                :sku,
                                                :quantity,
                                                :amazon_marketplace_id,
                                                :amazon_fulfillment_id,
                                                :amazon_tax_collection_model_id,
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
                                                ON CONFLICT (user_id, date_time, settlement_id, amazon_transaction_type_id, order_id, sku)
                                                    DO UPDATE SET quantity = :quantity,
                                                                  amazon_marketplace_id = :amazon_marketplace_id,
                                                                  amazon_fulfillment_id = :amazon_fulfillment_id,
                                                                  amazon_tax_collection_model_id = :amazon_tax_collection_model_id,
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
	_, err := db.NamedExecContext(ctx, query, map[string]interface{}{
		"user_id":                        ac.UserID,
		"date_time":                      ac.DateTime,
		"settlement_id":                  ac.SettlementID,
		"amazon_transaction_type_id":     ac.AmazonTransactionTypeID,
		"order_id":                       ac.OrderID,
		"sku":                            ac.SKU,
		"quantity":                       ac.Quantity,
		"amazon_marketplace_id":          ac.AmazonMarketplaceID,
		"amazon_fulfillment_id":          ac.AmazonFulfillmentID,
		"amazon_tax_collection_model_id": ac.AmazonTaxCollectionModelID,
		"product_sales":                  ac.ProductSales,
		"product_sales_tax":              ac.ProductSalesTax,
		"shipping_credits":               ac.ShippingCredits,
		"shipping_credits_tax":           ac.ShippingCreditsTax,
		"giftwrap_credits":               ac.GiftwrapCredits,
		"giftwrap_credits_tax":           ac.GiftwrapCreditsTax,
		"promotional_rebates":            ac.PromotionalRebates,
		"promotional_rebates_tax":        ac.PromotionalRebatesTax,
		"marketplace_withheld_tax":       ac.MarketplaceWithheldTax,
		"selling_fees":                   ac.SellingFees,
		"fba_fees":                       ac.FBAFees,
		"other_transaction_fees":         ac.OtherTransactionFees,
		"other":                          ac.Other,
		"total":                          ac.Total,
	})

	if err != nil {
		return err
	}

	return nil
}

func (ac *AmazonCustomTransactionModel) Update(ctx context.Context, db *sqlx.DB) error {
	query := `UPDATE amazon_custom_transactions SET
                                      quantity = :quantity,
                                      amazon_marketplace_id = :amazon_marketplace_id,
                                      amazon_fulfillment_id = :amazon_fulfillment_id,
                                      amazon_tax_collection_model_id = :amazon_tax_collection_model_id,
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
                                      updated_at = NOW()
                                      WHERE id = :id
                                        AND amazon_custom_transaction_id = :amazon_custom_transaction_id
                                        AND user_id = :user_id`
	_, err := db.NamedExecContext(ctx, query, map[string]interface{}{
		"quantity":                       ac.Quantity,
		"amazon_marketplace_id":          ac.AmazonMarketplaceID,
		"amazon_fulfillment_id":          ac.AmazonFulfillmentID,
		"amazon_tax_collection_model_id": ac.AmazonTaxCollectionModelID,
		"product_sales":                  ac.ProductSales,
		"product_sales_tax":              ac.ProductSalesTax,
		"shipping_credits":               ac.ShippingCredits,
		"shipping_credits_tax":           ac.ShippingCreditsTax,
		"giftwrap_credits":               ac.GiftwrapCredits,
		"giftwrap_credits_tax":           ac.GiftwrapCreditsTax,
		"promotional_rebates":            ac.PromotionalRebates,
		"promotional_rebates_tax":        ac.PromotionalRebatesTax,
		"marketplace_withheld_tax":       ac.MarketplaceWithheldTax,
		"selling_fees":                   ac.SellingFees,
		"fba_fees":                       ac.FBAFees,
		"other_transaction_fees":         ac.OtherTransactionFees,
		"other":                          ac.Other,
		"total":                          ac.Total,
		"id":                             ac.ID,
		"amazon_custom_transaction_id":   ac.AmazonCustomTransactionID,
		"user_id":                        ac.UserID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (ac *AmazonCustomTransactionModel) Delete(ctx context.Context, db *sqlx.DB) error { return nil }
