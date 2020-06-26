package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type SponsoredProduct struct {
	Id                 int64
	User               `db:"user_id"`
	StartDate          time.Time    `db:"start_date"`
	EndDate            time.Time    `db:"end_date"`
	PortfolioName      string       `db:"portfolio_name"`
	ExchangeRate       ExchangeRate `db:"exchange_rate_id"`
	CampaignName       string       `db:"campaign_name"`
	AdGroupName        string       `db:"ad_group_name"`
	SKU                string       `db:"sku"`
	ASIN               string       `db:"asin"`
	Impressions        int64        `db:"impressions"`
	Clicks             int64        `db:"clicks"`
	CTR                float64      `db:"ctr"`
	CPC                float64      `db:"cpc"`
	Spend              float64      `db:"spend"`
	TotalSales         float64      `db:"total_sales"`
	ACoS               float64      `db:"acos"`
	RoAS               float64      `db:"roas"`
	TotalOrders        int64        `db:"total_orders"`
	TotalUnits         int64        `db:"total_units"`
	ConversionRate     float64      `db:"conversion_rate"`
	AdvertisedSKUUnits int64        `db:"advertised_sku_units"`
	OtherSKUUnits      int64        `db:"other_sku_units"`
	AdvertisedSKUSales float64      `db:"advertised_sku_sales"`
	OtherSKUSales      float64      `db:"other_sku_sales"`
	CreatedAt          time.Time    `db:"created_at"`
	UpdatedAt          time.Time    `db:"updated_at"`
}

func (s *SponsoredProduct) Save(db *sqlx.DB) error {
	query := `INSERT INTO sponsored_products (
                                user_id,
                                start_date,
                                end_date,
                                portfolio_name,
                                exchange_rate_id,
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
                                    :start_date,
                                    :end_date,
                                    :portfolio_name,
                                    :exchange_rate_id,
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
                            ON CONFLICT (user_id, start_date, end_date, portfolio_name, campaign_name, ad_group_name, sku, asin)
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

	_, err := db.NamedExec(query, map[string]interface{}{
		"user_id":              s.User.Id,
		"start_date":           s.StartDate,
		"end_date":             s.EndDate,
		"portfolio_name":       s.PortfolioName,
		"exchange_rate_id":     s.ExchangeRate.Id,
		"campaign_name":        s.CampaignName,
		"ad_group_name":        s.AdGroupName,
		"sku":                  s.SKU,
		"asin":                 s.ASIN,
		"impressions":          s.Impressions,
		"clicks":               s.Clicks,
		"ctr":                  s.CTR,
		"cpc":                  s.CPC,
		"spend":                s.Spend,
		"total_sales":          s.Spend,
		"acos":                 s.ACoS,
		"roas":                 s.RoAS,
		"total_orders":         s.TotalOrders,
		"total_units":          s.TotalUnits,
		"conversion_rate":      s.ConversionRate,
		"advertised_sku_units": s.AdvertisedSKUUnits,
		"other_sku_units":      s.OtherSKUUnits,
		"advertised_sku_sales": s.AdvertisedSKUSales,
		"other_sku_sales":      s.OtherSKUSales,
	})

	if err != nil {
		logger().Error().Err(err)
	}

	return err
}
