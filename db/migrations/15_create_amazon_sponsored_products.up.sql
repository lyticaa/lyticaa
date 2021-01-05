CREATE TABLE amazon_sponsored_products
(
    id                      BIGSERIAL,
    user_id                 VARCHAR NOT NULL,
    date_time               TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    portfolio_name          VARCHAR NOT NULL,
    amazon_marketplace_id   BIGSERIAL REFERENCES amazon_marketplaces(id),
    campaign_name           VARCHAR NOT NULL,
    ad_group_name           VARCHAR NOT NULL,
    sku                     VARCHAR NOT NULL,
    asin                    VARCHAR NOT NULL,
    impressions             BIGSERIAL,
    clicks                  BIGSERIAL,
    ctr                     REAL,
    cpc                     REAL,
    spend                   REAL,
    total_sales             REAL,
    acos                    REAL,
    roas                    REAL,
    total_orders            BIGSERIAL,
    total_units             BIGSERIAL,
    conversion_rate         REAL,
    advertised_sku_units    BIGSERIAL,
    other_sku_units         BIGSERIAL,
    advertised_sku_sales    REAL,
    other_sku_sales         REAL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id, date_time, portfolio_name, campaign_name, ad_group_name, sku, asin)
);

CREATE MATERIALIZED VIEW amazon_sponsored_products_today AS
SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= date_trunc('day', NOW())
  AND date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
GROUP BY user_id, date_trunc('day', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('day', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_yesterday AS
SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= date_trunc('day', NOW()) - interval '1 day'
  AND date_time <= date_trunc('day', NOW()) - interval '1 second'
GROUP BY user_id, date_trunc('day', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('day', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_last_thirty_days AS
SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= date_trunc('day', NOW()) - interval '30 day'
  AND date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
GROUP BY user_id, date_trunc('day', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('day', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_previous_thirty_days AS
SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= date_trunc('day', NOW()) - interval '60 day'
  AND date_time <= date_trunc('day', NOW()) - interval '30 day' - interval '1 second'
GROUP BY user_id, date_trunc('day', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('day', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_this_month AS
SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= date_trunc('month', NOW())
  AND date_time <= date_trunc('month', NOW()) + interval '1 month' - interval '1 second'
GROUP BY user_id, date_trunc('day', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('day', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_last_month AS
SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= date_trunc('month', NOW()) - interval '1 month'
  AND date_time <= date_trunc('month', NOW()) - interval '1 second'
GROUP BY user_id, date_trunc('day', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('day', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_month_before_last AS
SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= date_trunc('month', NOW()) - interval '2 month'
  AND date_time <= date_trunc('month', NOW()) - interval '1 month' - interval '1 second'
GROUP BY user_id, date_trunc('day', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('day', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_last_three_months AS
SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= NOW() - interval '3 month'
  AND date_time <= NOW()
GROUP BY user_id, date_trunc('week', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('week', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_previous_three_months AS
SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= NOW() - interval '6 month'
  AND date_time <= NOW() - interval '3 month'
GROUP BY user_id, date_trunc('week', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('week', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_last_six_months AS
SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= NOW() - interval '6 month'
  AND date_time <= NOW()
GROUP BY user_id, date_trunc('week', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('week', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_previous_six_months AS
SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= NOW() - interval '12 month'
  AND date_time <= NOW() - interval '6 month'
GROUP BY user_id, date_trunc('week', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('week', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_this_year AS
SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= date_trunc('year', NOW())
  AND date_time <= NOW()
GROUP BY user_id, date_trunc('week', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('week', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_last_year AS
SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
WHERE date_time >= date_trunc('year', NOW()) - interval '1 year'
  AND date_time <= date_trunc('year', NOW()) - interval '1 second'
GROUP BY user_id, date_trunc('week', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('week', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_sponsored_products_all_time AS
SELECT user_id,
       date_trunc('month', date_time) AS date_time,
       portfolio_name,
       amazon_marketplace_id,
       campaign_name,
       ad_group_name,
       sku,
       asin,
       SUM(impressions) AS impressions,
       SUM(clicks) AS clicks,
       SUM(ctr) AS ctr,
       SUM(cpc) AS cpc,
       SUM(spend) AS spend,
       SUM(total_sales) AS total_sales,
       SUM(acos) AS acos,
       SUM(roas) AS roas,
       SUM(total_orders) AS total_orders,
       SUM(total_units) AS total_units,
       AVG(conversion_rate) AS conversion_rate,
       SUM(advertised_sku_units) AS advertised_sku_units,
       SUM(other_sku_units) AS other_sku_units,
       SUM(advertised_sku_sales) AS advertised_sku_sales,
       SUM(other_sku_sales) AS other_sku_sales
FROM amazon_sponsored_products
GROUP BY user_id, date_trunc('month', date_time), portfolio_name, amazon_marketplace_id, campaign_name, ad_group_name, sku, asin
ORDER BY date_trunc('month', date_time), amazon_marketplace_id ASC;
