CREATE TABLE transaction_types
(
    id         BIGSERIAL,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (name)
);

CREATE TABLE marketplaces
(
    id         BIGSERIAL,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (name)
);

CREATE TABLE exchange_rates
(
    id              BIGSERIAL,
    marketplace_id  BIGSERIAL REFERENCES marketplaces(id) ON DELETE CASCADE,
    code            VARCHAR NOT NULL,
    symbol          VARCHAR NOT NULL,
    rate            REAL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE cost_of_goods
(
    id                  BIGSERIAL NOT NULL,
    user_id             BIGSERIAL REFERENCES users(id) ON DELETE CASCADE,
    marketplace_id      BIGSERIAL REFERENCES marketplaces(id) ON DELETE CASCADE,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    cost                REAL,
    start_at            TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id, marketplace_id, sku, start_at, end_at)
);

CREATE TABLE fulfillments
(
    id         BIGSERIAL,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (name)
);

CREATE TABLE tax_collection_models
(
    id         BIGSERIAL,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (name)
);

CREATE TABLE transactions
(
    id                       BIGSERIAL NOT NULL,
    user_id                  BIGSERIAL REFERENCES users(id) ON DELETE CASCADE,
    date_time                TIMESTAMPTZ NOT NULL,
    settlement_id            BIGSERIAL NOT NULL,
    settlement_idx           BIGSERIAL NOT NULL,
    transaction_type_id      BIGSERIAL REFERENCES transaction_types(id),
    order_id                 VARCHAR NOT NULL,
    sku                      VARCHAR NOT NULL,
    description              VARCHAR NOT NULL,
    quantity                 BIGSERIAL,
    marketplace_id           BIGSERIAL REFERENCES marketplaces(id),
    fulfillment_id           BIGSERIAL REFERENCES fulfillments(id),
    tax_collection_model_id  BIGSERIAL REFERENCES tax_collection_models(id),
    product_sales            REAL,
    product_sales_tax        REAL,
    shipping_credits         REAL,
    shipping_credits_tax     REAL,
    giftwrap_credits         REAL,
    giftwrap_credits_tax     REAL,
    promotional_rebates      REAL,
    promotional_rebates_tax  REAL,
    marketplace_withheld_tax REAL,
    selling_fees             REAL,
    fba_fees                 REAL,
    other_transaction_fees   REAL,
    other                    REAL,
    total                    REAL,
    created_at               TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at               TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id, date_time, settlement_id, settlement_idx, transaction_type_id, order_id, sku)
);

INSERT INTO transaction_types (name, created_at, updated_at) VALUES ('UNKNOWN', NOW(), NOW());
INSERT INTO transaction_types (name, created_at, updated_at) VALUES ('Order', NOW(), NOW());
INSERT INTO transaction_types (name, created_at, updated_at) VALUES ('Refund', NOW(), NOW());
INSERT INTO transaction_types (name, created_at, updated_at) VALUES ('Service Fee', NOW(), NOW());
INSERT INTO transaction_types (name, created_at, updated_at) VALUES ('Adjustment', NOW(), NOW());
INSERT INTO transaction_types (name, created_at, updated_at) VALUES ('Transfer', NOW(), NOW());
INSERT INTO transaction_types (name, created_at, updated_at) VALUES ('FBA Inventory Fee', NOW(), NOW());

INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('UNKNOWN', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.com', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.ca', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.mx', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.br', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.ae', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.de', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.es', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.fr', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.co.uk', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.in', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.it', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.nl', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.sa', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.com.tr', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.sg', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.com.au', NOW(), NOW());
INSERT INTO marketplaces (name, created_at, updated_at) VALUES ('amazon.co.jp', NOW(), NOW());

INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.com'), 'UNKNOWN', '$', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.com'), 'USD', '$', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.ca'), 'CAD', '$', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.mx'), 'MXN', '$', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.br'), 'BRL', 'R$', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.ae'), 'AED', 'Dhs', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.de'), 'EUR', '€', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.es'), 'EUR', '€', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.fr'), 'EUR', '€', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.co.uk'), 'GBP', '£', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.in'), 'INR', '₹', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.it'), 'EUR', '€', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.nl'), 'EUR', '€', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.sa'), 'SAR', 'SAR', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.com.tr'), 'TRY', '₺', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.sg'), 'SGD', '$', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.com.au'), 'AUD', '$', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (marketplace_id, code, symbol, rate, created_at, updated_at) VALUES ((SELECT id FROM marketplaces WHERE name = 'amazon.co.jp'), 'JPY', '¥', 1.00, NOW(), NOW());

INSERT INTO fulfillments (name, created_at, updated_at) VALUES ('UNKNOWN', NOW(), NOW());
INSERT INTO fulfillments (name, created_at, updated_at) VALUES ('Amazon', NOW(), NOW());

INSERT INTO tax_collection_models (name, created_at, updated_at) VALUES ('UNKNOWN', NOW(), NOW());
INSERT INTO tax_collection_models (name, created_at, updated_at) VALUES ('MarketplaceFacilitator', NOW(), NOW());

/* Products */
CREATE MATERIALIZED VIEW products AS
    SELECT DISTINCT(t.sku) AS sku, t.user_id, t.description, m.name AS marketplace
    FROM transactions t
        LEFT JOIN marketplaces m ON t.marketplace_id = m.id
    GROUP BY t.sku, t.user_id, t.description, m.name;

/* Transactions Views */
CREATE MATERIALIZED VIEW transactions_today AS
    SELECT user_id,
       date_trunc('hour', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= date_trunc('day', NOW())
      AND date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY user_id, date_trunc('hour', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('hour', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_yesterday AS
    SELECT user_id,
       date_trunc('hour', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= date_trunc('day', NOW()) - interval '1 day'
      AND date_time <= date_trunc('day', NOW()) - interval '1 second'
    GROUP BY user_id, date_trunc('hour', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('hour', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_last_thirty_days AS
    SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= date_trunc('day', NOW()) - interval '30 day'
      AND date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY user_id, date_trunc('day', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('day', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_previous_thirty_days AS
    SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= date_trunc('day', NOW()) - interval '60 day'
      AND date_time <= date_trunc('day', NOW()) - interval '30 day' - interval '1 second'
    GROUP BY user_id, date_trunc('day', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('day', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_this_month AS
    SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= date_trunc('month', NOW())
      AND date_time <= date_trunc('month', NOW()) + interval '1 month' - interval '1 second'
    GROUP BY user_id, date_trunc('day', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('day', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_last_month AS
    SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= date_trunc('month', NOW()) - interval '1 month'
      AND date_time <= date_trunc('month', NOW()) - interval '1 second'
    GROUP BY user_id, date_trunc('day', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('day', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_month_before_last AS
    SELECT user_id,
       date_trunc('day', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= date_trunc('month', NOW()) - interval '2 month'
      AND date_time <= date_trunc('month', NOW()) - interval '1 month' - interval '1 second'
    GROUP BY user_id, date_trunc('day', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('day', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_last_three_months AS
    SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= NOW() - interval '3 month'
      AND date_time <= NOW()
    GROUP BY user_id, date_trunc('week', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('week', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_previous_three_months AS
    SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= NOW() - interval '6 month'
      AND date_time <= NOW() - interval '3 month'
    GROUP BY user_id, date_trunc('week', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('week', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_last_six_months AS
    SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= NOW() - interval '6 month'
      AND date_time <= NOW()
    GROUP BY user_id, date_trunc('week', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('week', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_previous_six_months AS
    SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= NOW() - interval '12 month'
      AND date_time <= NOW() - interval '6 month'
    GROUP BY user_id, date_trunc('week', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('week', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_this_year AS
    SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= date_trunc('year', NOW())
      AND date_time <= NOW()
    GROUP BY user_id, date_trunc('week', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('week', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_last_year AS
    SELECT user_id,
       date_trunc('week', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    WHERE date_time >= date_trunc('year', NOW()) - interval '1 year'
      AND date_time <= date_trunc('year', NOW()) - interval '1 second'
    GROUP BY user_id, date_trunc('week', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('week', date_time), marketplace_id ASC;

CREATE MATERIALIZED VIEW transactions_all_time AS
    SELECT user_id,
       date_trunc('month', date_time) AS date_time,
       transaction_type_id,
       sku,
       SUM(quantity) AS quantity,
       marketplace_id,
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
    FROM transactions
    GROUP BY user_id, date_trunc('month', date_time), transaction_type_id, sku, marketplace_id
    ORDER BY date_trunc('month', date_time), marketplace_id ASC;
