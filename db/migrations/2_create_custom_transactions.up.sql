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

/* Total Sales */
CREATE MATERIALIZED VIEW total_sales_today AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('day', NOW())
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_yesterday AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('day', NOW()) - interval '1 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_last_thirty_days AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('day', NOW()) - interval '30 day'
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_previous_thirty_days AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('day', NOW()) - interval '60 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '30 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_this_month AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('month', NOW())
      AND t.date_time <= date_trunc('month', NOW()) + interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_last_month AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('month', NOW()) - interval '1 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_month_before_last AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('month', NOW()) - interval '2 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_last_three_months AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= NOW() - interval '3 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_previous_three_months AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW() - interval '3 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_last_six_months AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_previous_six_months AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= NOW() - interval '12 month'
      AND t.date_time <= NOW() - interval '6 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_this_year AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('year', NOW())
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_last_year AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('year', NOW()) - interval '1 year'
      AND t.date_time <= date_trunc('year', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_sales_all_time AS
    SELECT t.user_id, SUM(product_sales*e.rate) AS total, date_trunc('month', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
    GROUP BY t.user_id, date_trunc('month', t.date_time), m.name;

/* Units Sold */
CREATE MATERIALIZED VIEW units_sold_today AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('day', NOW())
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_yesterday AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('day', NOW()) - interval '1 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_last_thirty_days AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('day', NOW()) - interval '30 day'
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_previous_thirty_days AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('day', NOW()) - interval '60 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '30 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_this_month AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('month', NOW())
      AND t.date_time <= date_trunc('month', NOW()) + interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_last_month AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('month', NOW()) - interval '1 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_month_before_last AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('month', NOW()) - interval '2 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_last_three_months AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= NOW() - interval '3 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_previous_three_months AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW() - interval '3 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_last_six_months AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_previous_six_months AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= NOW() - interval '12 month'
      AND t.date_time <= NOW() - interval '6 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_this_year AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('year', NOW())
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_last_year AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('year', NOW()) - interval '1 year'
      AND t.date_time <= date_trunc('year', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW units_sold_all_time AS
    SELECT t.user_id, SUM(quantity) AS total, date_trunc('month', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
    WHERE tt.name = 'Order'
    GROUP BY t.user_id, date_trunc('month', t.date_time), m.name;

/* Amazon Costs */
CREATE MATERIALIZED VIEW amazon_costs_today AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= date_trunc('day', NOW())
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_yesterday AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= date_trunc('day', NOW()) - interval '1 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_last_thirty_days AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= date_trunc('day', NOW()) - interval '30 day'
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_previous_thirty_days AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= date_trunc('day', NOW()) - interval '60 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '30 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_this_month AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= date_trunc('month', NOW())
      AND t.date_time <= date_trunc('month', NOW()) + interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_last_month AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= date_trunc('month', NOW()) - interval '1 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_month_before_last AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= date_trunc('month', NOW()) - interval '2 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_last_three_months AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= NOW() - interval '3 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_previous_three_months AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW() - interval '3 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_last_six_months AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_previous_six_months AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= NOW() - interval '12 month'
      AND t.date_time <= NOW() - interval '6 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_this_year AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= date_trunc('year', NOW())
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_last_year AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
      AND t.date_time >= date_trunc('year', NOW()) - interval '1 year'
      AND t.date_time <= date_trunc('year', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW amazon_costs_all_time AS
    SELECT t.user_id, SUM((selling_fees+fba_fees+other)*e.rate) AS total, date_trunc('month', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order','Service Fee','FBA Inventory Fee')
    GROUP BY t.user_id, date_trunc('month', t.date_time), m.name;

/* Advertising Spend */
CREATE MATERIALIZED VIEW advertising_spend_today AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= date_trunc('day', NOW())
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_yesterday AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= date_trunc('day', NOW()) - interval '1 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_last_thirty_days AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= date_trunc('day', NOW()) - interval '30 day'
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_previous_thirty_days AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= date_trunc('day', NOW()) - interval '60 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '30 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_this_month AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= date_trunc('month', NOW())
      AND t.date_time <= date_trunc('month', NOW()) + interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_last_month AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= date_trunc('month', NOW()) - interval '1 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_month_before_last AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= date_trunc('month', NOW()) - interval '2 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_last_three_months AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
         LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
         LEFT JOIN marketplaces m on t.marketplace_id = m.id
         LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= NOW() - interval '3 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_previous_three_months AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW() - interval '3 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_last_six_months AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_previous_six_months AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= NOW() - interval '12 month'
      AND t.date_time <= NOW() - interval '6 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_this_year AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= date_trunc('year', NOW())
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_last_year AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
      AND t.date_time >= date_trunc('year', NOW()) - interval '1 year'
      AND t.date_time <= date_trunc('year', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW advertising_spend_all_time AS
    SELECT t.user_id, SUM(other_transaction_fees*e.rate) AS total, date_trunc('month', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Service Fee'
    GROUP BY t.user_id, date_trunc('month', t.date_time), m.name;

/* Refunds */
CREATE MATERIALIZED VIEW refunds_today AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= date_trunc('day', NOW())
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_yesterday AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= date_trunc('day', NOW()) - interval '1 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_last_thirty_days AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= date_trunc('day', NOW()) - interval '30 day'
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_previous_thirty_days AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= date_trunc('day', NOW()) - interval '60 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '30 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_this_month AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= date_trunc('month', NOW())
      AND t.date_time <= date_trunc('month', NOW()) + interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_last_month AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= date_trunc('month', NOW()) - interval '1 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_month_before_last AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= date_trunc('month', NOW()) - interval '2 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_last_three_months AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= NOW() - interval '3 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_previous_three_months AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW() - interval '3 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_last_six_months AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_previous_six_months AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= NOW() - interval '12 month'
      AND t.date_time <= NOW() - interval '6 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_this_year AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= date_trunc('year', NOW())
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_last_year AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
      AND t.date_time >= date_trunc('year', NOW()) - interval '1 year'
      AND t.date_time <= date_trunc('year', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW refunds_all_time AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('month', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Refund'
    GROUP BY t.user_id, date_trunc('month', t.date_time), m.name;

/* Refunds */
CREATE MATERIALIZED VIEW shipping_credits_today AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('day', NOW())
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_yesterday AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('day', NOW()) - interval '1 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_last_thirty_days AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('day', NOW()) - interval '30 day'
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_previous_thirty_days AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('day', NOW()) - interval '60 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '30 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_this_month AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('month', NOW())
      AND t.date_time <= date_trunc('month', NOW()) + interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_last_month AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('month', NOW()) - interval '1 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_month_before_last AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('month', NOW()) - interval '2 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_last_three_months AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= NOW() - interval '3 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_previous_three_months AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW() - interval '3 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_last_six_months AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_previous_six_months AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= NOW() - interval '12 month'
      AND t.date_time <= NOW() - interval '6 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_this_year AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('year', NOW())
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_last_year AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
      AND t.date_time >= date_trunc('year', NOW()) - interval '1 year'
      AND t.date_time <= date_trunc('year', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW shipping_credits_all_time AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax)*e.rate) AS total, date_trunc('month', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name = 'Order'
    GROUP BY t.user_id, date_trunc('month', t.date_time), m.name;

/* Promotional Rebates */
CREATE MATERIALIZED VIEW promotional_rebates_today AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= date_trunc('day', NOW())
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_yesterday AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= date_trunc('day', NOW()) - interval '1 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_last_thirty_days AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= date_trunc('day', NOW()) - interval '30 day'
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_previous_thirty_days AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= date_trunc('day', NOW()) - interval '60 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '30 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_this_month AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= date_trunc('month', NOW())
      AND t.date_time <= date_trunc('month', NOW()) + interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_last_month AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= date_trunc('month', NOW()) - interval '1 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_month_before_last AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= date_trunc('month', NOW()) - interval '2 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_last_three_months AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= NOW() - interval '3 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_previous_three_months AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW() - interval '3 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_last_six_months AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_previous_six_months AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= NOW() - interval '12 month'
      AND t.date_time <= NOW() - interval '6 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_this_year AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= date_trunc('year', NOW())
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_last_year AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
      AND t.date_time >= date_trunc('year', NOW()) - interval '1 year'
      AND t.date_time <= date_trunc('year', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW promotional_rebates_all_time AS
    SELECT t.user_id, SUM((promotional_rebates+promotional_rebates_tax)*e.rate) AS total, date_trunc('month', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund')
    GROUP BY t.user_id, date_trunc('month', t.date_time), m.name;

/* Total Costs */
CREATE MATERIALIZED VIEW total_costs_today AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('day', NOW())
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_yesterday AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('day', NOW()) - interval '1 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_last_thirty_days AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('day', NOW()) - interval '30 day'
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_previous_thirty_days AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('day', NOW()) - interval '60 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '30 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_this_month AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('month', NOW())
      AND t.date_time <= date_trunc('month', NOW()) + interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_last_month AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('month', NOW()) - interval '1 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_month_before_last AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('month', NOW()) - interval '2 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_last_three_months AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= NOW() - interval '3 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_previous_three_months AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW() - interval '3 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_last_six_months AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_previous_six_months AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= NOW() - interval '12 month'
      AND t.date_time <= NOW() - interval '6 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_this_year AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('year', NOW())
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_last_year AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('year', NOW()) - interval '1 year'
      AND t.date_time <= date_trunc('year', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW total_costs_all_time AS
    SELECT t.user_id, SUM((shipping_credits+shipping_credits_tax+promotional_rebates+promotional_rebates_tax+selling_fees+fba_fees+other_transaction_fees+other)*e.rate) AS total,
           date_trunc('month', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
    GROUP BY t.user_id, date_trunc('month', t.date_time), m.name;

/* Net Margin */
CREATE MATERIALIZED VIEW net_margin_today AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('day', NOW())
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_yesterday AS
    SELECT t.user_id, SUM(total*e.rate)  AS total, date_trunc('hour', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('day', NOW()) - interval '1 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('hour', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_last_thirty_days AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('day', NOW()) - interval '30 day'
      AND t.date_time <= date_trunc('day', NOW()) + interval '1 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_previous_thirty_days AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('day', NOW()) - interval '60 day'
      AND t.date_time <= date_trunc('day', NOW()) - interval '30 day' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_this_month AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('month', NOW())
      AND t.date_time <= date_trunc('month', NOW()) + interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_last_month AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('month', NOW()) - interval '1 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_month_before_last AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('day', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('month', NOW()) - interval '2 month'
      AND t.date_time <= date_trunc('month', NOW()) - interval '1 month' - interval '1 second'
    GROUP BY t.user_id, date_trunc('day', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_last_three_months AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= NOW() - interval '3 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_previous_three_months AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW() - interval '3 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_last_six_months AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= NOW() - interval '6 month'
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_previous_six_months AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= NOW() - interval '12 month'
      AND t.date_time <= NOW() - interval '6 month'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_this_year AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('year', NOW())
      AND t.date_time <= NOW()
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_last_year AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('week', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
      AND t.date_time >= date_trunc('year', NOW()) - interval '1 year'
      AND t.date_time <= date_trunc('year', NOW()) - interval '1 second'
    GROUP BY t.user_id, date_trunc('week', t.date_time), m.name;

CREATE MATERIALIZED VIEW net_margin_all_time AS
    SELECT t.user_id, SUM(total*e.rate) AS total, date_trunc('month', date_time) AS order_date, m.name AS marketplace
    FROM transactions t
        LEFT JOIN transaction_types tt ON t.transaction_type_id = tt.id
        LEFT JOIN marketplaces m on t.marketplace_id = m.id
        LEFT JOIN exchange_rates e ON m.id = e.marketplace_id
    WHERE tt.name IN ('Order', 'Refund', 'Service Fee')
    GROUP BY t.user_id, date_trunc('month', t.date_time), m.name;
