CREATE TABLE amazon_custom_transactions
(
    id                              BIGSERIAL NOT NULL,
    amazon_custom_transaction_id    UUID DEFAULT UUID_GENERATE_V4(),
    user_id                         VARCHAR NOT NULL,
    date_time                       TIMESTAMPTZ NOT NULL,
    settlement_id                   BIGSERIAL NOT NULL,
    amazon_transaction_type_id      BIGSERIAL REFERENCES amazon_transaction_types(id),
    order_id                        VARCHAR NOT NULL,
    sku                             VARCHAR NOT NULL,
    quantity                        BIGINT,
    amazon_marketplace_id           BIGSERIAL REFERENCES amazon_marketplaces(id),
    amazon_fulfillment_id           BIGSERIAL REFERENCES amazon_fulfillments(id),
    amazon_tax_collection_model_id  BIGSERIAL REFERENCES amazon_tax_collection_models(id),
    product_sales                   REAL,
    product_sales_tax               REAL,
    shipping_credits                REAL,
    shipping_credits_tax            REAL,
    giftwrap_credits                REAL,
    giftwrap_credits_tax            REAL,
    promotional_rebates             REAL,
    promotional_rebates_tax         REAL,
    marketplace_withheld_tax        REAL,
    selling_fees                    REAL,
    fba_fees                        REAL,
    other_transaction_fees          REAL,
    other                           REAL,
    total                           REAL,
    created_at                      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at                      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id, date_time, settlement_id, amazon_transaction_type_id, order_id, sku)
);

CREATE MATERIALIZED VIEW amazon_custom_transactions_today AS
SELECT user_id,
       date_trunc('hour', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= date_trunc('day', NOW())
  AND date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
GROUP BY user_id, date_trunc('hour', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('hour', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_yesterday AS
SELECT user_id,
       date_trunc('hour', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= date_trunc('day', NOW()) - interval '1 day'
  AND date_time <= date_trunc('day', NOW()) - interval '1 second'
GROUP BY user_id, date_trunc('hour', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('hour', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_last_thirty_days AS
SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= date_trunc('day', NOW()) - interval '30 day'
  AND date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
GROUP BY user_id, date_trunc('day', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('day', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_previous_thirty_days AS
SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= date_trunc('day', NOW()) - interval '60 day'
  AND date_time <= date_trunc('day', NOW()) - interval '30 day' - interval '1 second'
GROUP BY user_id, date_trunc('day', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('day', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_this_month AS
SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= date_trunc('month', NOW())
  AND date_time <= date_trunc('month', NOW()) + interval '1 month' - interval '1 second'
GROUP BY user_id, date_trunc('day', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('day', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_last_month AS
SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= date_trunc('month', NOW()) - interval '1 month'
  AND date_time <= date_trunc('month', NOW()) - interval '1 second'
GROUP BY user_id, date_trunc('day', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('day', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_month_before_last AS
SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= date_trunc('month', NOW()) - interval '2 month'
  AND date_time <= date_trunc('month', NOW()) - interval '1 month' - interval '1 second'
GROUP BY user_id, date_trunc('day', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('day', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_last_three_months AS
SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= NOW() - interval '3 month'
  AND date_time <= NOW()
GROUP BY user_id, date_trunc('week', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('week', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_previous_three_months AS
SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= NOW() - interval '6 month'
  AND date_time <= NOW() - interval '3 month'
GROUP BY user_id, date_trunc('week', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('week', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_last_six_months AS
SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= NOW() - interval '6 month'
  AND date_time <= NOW()
GROUP BY user_id, date_trunc('week', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('week', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_previous_six_months AS
SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= NOW() - interval '12 month'
  AND date_time <= NOW() - interval '6 month'
GROUP BY user_id, date_trunc('week', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('week', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_this_year AS
SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= date_trunc('year', NOW())
  AND date_time <= NOW()
GROUP BY user_id, date_trunc('week', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('week', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_last_year AS
SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
WHERE date_time >= date_trunc('year', NOW()) - interval '1 year'
  AND date_time <= date_trunc('year', NOW()) - interval '1 second'
GROUP BY user_id, date_trunc('week', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('week', date_time), amazon_marketplace_id ASC;

CREATE MATERIALIZED VIEW amazon_custom_transactions_all_time AS
SELECT user_id,
       date_trunc('month', date_time) AS date_time,
       amazon_transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       amazon_marketplace_id,
       SUM(product_sales) AS product_sales,
       SUM(product_sales_tax) AS product_sales_tax,
       SUM(shipping_credits) AS shipping_credits,
       SUM(shipping_credits_tax) AS shipping_credits_tax,
       SUM(giftwrap_credits) AS giftwrap_credits,
       SUM(giftwrap_credits_tax) AS giftwrap_credits_tax,
       SUM(promotional_rebates) AS promotional_rebates,
       SUM(promotional_rebates_tax) AS promotional_rebates_tax,
       SUM(marketplace_withheld_tax) AS marketplace_withheld_tax,
       SUM(selling_fees) AS selling_fees,
       SUM(fba_fees) AS fba_fees,
       SUM(other_transaction_fees) AS other_transaction_fees,
       SUM(other) AS other,
       SUM(total) AS total
FROM amazon_custom_transactions
GROUP BY user_id, date_trunc('month', date_time), amazon_transaction_type_id, sku, amazon_marketplace_id
ORDER BY date_trunc('month', date_time), amazon_marketplace_id ASC;
