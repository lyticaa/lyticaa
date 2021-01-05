package models

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type AmazonSponsoredProductModel struct {
	ID                       int64     `db:"id"`
	AmazonSponsoredProductID string    `db:"amazon_sponsored_product_id"`
	UserID                   string    `db:"user_id"`
	DateTime                 time.Time `db:"date_time"`
	PortfolioName            string    `db:"portfolio_name"`
	AmazonMarketplaceID      int64     `db:"amazon_marketplace_id"`
	CampaignName             string    `db:"campaign_name"`
	AdGroupName              string    `db:"ad_group_name"`
	SKU                      string    `db:"sku"`
	ASIN                     string    `db:"asin"`
	Impressions              int64     `db:"impressions"`
	Clicks                   int64     `db:"clicks"`
	CTR                      float64   `db:"ctr"`
	CPC                      float64   `db:"cpc"`
	Spend                    float64   `db:"spend"`
	TotalSales               float64   `db:"total_sales"`
	ACoS                     float64   `db:"acos"`
	RoAS                     float64   `db:"roas"`
	TotalOrders              int64     `db:"total_orders"`
	TotalUnits               int64     `db:"total_units"`
	ConversionRate           float64   `db:"conversion_rate"`
	AdvertisedSKUUnits       int64     `db:"advertised_sku_units"`
	OtherSKUUnits            int64     `db:"other_sku_units"`
	AdvertisedSKUSales       float64   `db:"advertised_sku_sales"`
	OtherSKUSales            float64   `db:"other_sku_sales"`
	CreatedAt                time.Time `db:"created_at"`
	UpdatedAt                time.Time `db:"updated_at"`
}

func (as *AmazonSponsoredProductModel) FetchOne(ctx context.Context, db *sqlx.DB) interface{} {
	return nil
}
func (as *AmazonSponsoredProductModel) FetchBy(ctx context.Context, db *sqlx.DB) interface{} {
	return nil
}
func (as *AmazonSponsoredProductModel) FetchAll(ctx context.Context, data map[string]interface{}, filter *Filter, db *sqlx.DB) interface{} {
	return nil
}
func (as *AmazonSponsoredProductModel) Count(ctx context.Context, data map[string]interface{}, db *sqlx.DB) int64 {
	return int64(0)
}

func (as *AmazonSponsoredProductModel) Create(ctx context.Context, db *sqlx.DB) error {
	query := `INSERT INTO amazon_sponsored_products (
                                       user_id,
                                       date_time,
                                       portfolio_name,
                                       amazon_marketplace_id,
                                       campaign_name,
                                       ad_group_name,
                                       sku,
                                       asin,
                                       impressions,
                                       clicks,
                                       ctr,
                                       cpc,
                                       spend,
                                       total_sales,
                                       acos,
                                       roas,
                                       total_orders,
                                       total_units,
                                       conversion_rate,
                                       advertised_sku_units,
                                       other_sku_units,
                                       advertised_sku_sales,
                                       other_sku_sales)
                                       VALUES (
                                               :user_id,
                                               :date_time,
                                               :portfolio_name,
                                               :amazon_marketplace_id,
                                               :campaign_name,
                                               :ad_group_name,
                                               :sku,
                                               :asin,
                                               :impressions,
                                               :clicks,
                                               :ctr,
                                               :cpc,
                                               :spend,
                                               :total_sales,
                                               :acos,
                                               :roas,
                                               :total_orders,
                                               :total_units,
                                               :conversion_rate,
                                               :advertised_sku_units,
                                               :other_sku_units,
                                               :advertised_sku_sales,
                                               :other_sku_sales)
                                               ON CONFLICT (user_id, date_time, portfolio_name, campaign_name, ad_group_name, sku, asin)
                                                   DO UPDATE SET impressions = :impressions,
                                                                 clicks = :clicks,
                                                                 ctr = :ctr,
                                                                 cpc = :cpc,
                                                                 spend = :spend,
                                                                 total_sales = :total_sales,
                                                                 acos = :acos,
                                                                 roas = :roas,
                                                                 total_orders = :total_orders,
                                                                 total_units = :total_units,
                                                                 conversion_rate = :conversion_rate,
                                                                 advertised_sku_units = :advertised_sku_units,
                                                                 other_sku_units = :other_sku_units,
                                                                 advertised_sku_sales = :advertised_sku_sales,
                                                                 other_sku_sales = :other_sku_sales,
                                                                 updated_at = NOW()`
	_, err := db.NamedExecContext(ctx, query, map[string]interface{}{
		"user_id":               as.UserID,
		"date_time":             as.DateTime,
		"portfolio_name":        as.PortfolioName,
		"amazon_marketplace_id": as.AmazonMarketplaceID,
		"campaign_name":         as.CampaignName,
		"ad_group_name":         as.AdGroupName,
		"sku":                   as.SKU,
		"asin":                  as.ASIN,
		"impressions":           as.Impressions,
		"clicks":                as.Clicks,
		"ctr":                   as.CTR,
		"cpc":                   as.CPC,
		"spend":                 as.Spend,
		"total_sales":           as.Spend,
		"acos":                  as.ACoS,
		"roas":                  as.RoAS,
		"total_orders":          as.TotalOrders,
		"total_units":           as.TotalUnits,
		"conversion_rate":       as.ConversionRate,
		"advertised_sku_units":  as.AdvertisedSKUUnits,
		"other_sku_units":       as.OtherSKUUnits,
		"advertised_sku_sales":  as.AdvertisedSKUSales,
		"other_sku_sales":       as.OtherSKUSales,
	})

	if err != nil {
		return err
	}

	return nil
}

func (as *AmazonSponsoredProductModel) Update(ctx context.Context, db *sqlx.DB) error {
	query := `UPDATE amazon_sponsored_products SET
                                     impressions = :impressions,
                                     clicks = :clicks,
                                     ctr = :ctr,
                                     cpc = :cpc,
                                     spend = :spend,
                                     total_sales = :total_sales,
                                     acos = :acos,
                                     roas = :roas,
                                     total_orders = :total_orders,
                                     total_units = :total_units,
                                     conversion_rate = :conversion_rate,
                                     advertised_sku_units = :advertised_sku_units,
                                     other_sku_units = :other_sku_units,
                                     advertised_sku_sales = :advertised_sku_sales,
                                     other_sku_sales = :other_sku_sales,
                                     updated_at = NOW()
                                     WHERE id := :id
                                       AND amazon_sponsored_product_id = :amazon_sponsored_products
                                       AND user_id = :user_id`
	_, err := db.NamedExecContext(ctx, query, map[string]interface{}{
		"impressions":                 as.Impressions,
		"clicks":                      as.Clicks,
		"ctr":                         as.CTR,
		"cpc":                         as.CPC,
		"spend":                       as.Spend,
		"total_sales":                 as.Spend,
		"acos":                        as.ACoS,
		"roas":                        as.RoAS,
		"total_orders":                as.TotalOrders,
		"total_units":                 as.TotalUnits,
		"conversion_rate":             as.ConversionRate,
		"advertised_sku_units":        as.AdvertisedSKUUnits,
		"other_sku_units":             as.OtherSKUUnits,
		"advertised_sku_sales":        as.AdvertisedSKUSales,
		"other_sku_sales":             as.OtherSKUSales,
		"id":                          as.ID,
		"amazon_sponsored_product_id": as.AmazonSponsoredProductID,
		"user_id":                     as.UserID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (as *AmazonSponsoredProductModel) Delete(ctx context.Context, db *sqlx.DB) error { return nil }
