CREATE TABLE currencies
(
    id         BIGSERIAL,
    name       VARCHAR NOT NULL,
    symbol     VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (name)
);

CREATE TABLE sponsored_products
(
    id                   BIGSERIAL,
    user_id              BIGSERIAL REFERENCES users(id),
    start_date           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_date             TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    portfolio_name       VARCHAR NOT NULL,
    currency_id          BIGSERIAL REFERENCES currencies(id),
    campaign_name        VARCHAR NOT NULL,
    ad_group_name        VARCHAR NOT NULL,
    sku                  VARCHAR NOT NULL,
    asin                 VARCHAR NOT NULL,
    impressions          BIGSERIAL,
    clicks               BIGSERIAL,
    ctr                  REAL,
    cpc                  REAL,
    spend                REAL,
    total_sales          REAL,
    acos                 REAL,
    roas                 REAL,
    total_orders         BIGSERIAL,
    total_units          BIGSERIAL,
    conversion_rate      REAL,
    advertised_sku_units BIGSERIAL,
    other_sku_units      BIGSERIAL,
    advertised_sku_sales REAL,
    other_sku_sales      REAL,
    created_at           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id, start_date, end_date, portfolio_name, campaign_name, ad_group_name, sku, asin)
);

INSERT INTO currencies (name, symbol, created_at, updated_at) VALUES ('UNKNOWN', 'N/A', NOW(), NOW());
INSERT INTO currencies (name, symbol, created_at, updated_at) VALUES ('USD', '$', NOW(), NOW());
INSERT INTO currencies (name, symbol, created_at, updated_at) VALUES ('GBP', '£', NOW(), NOW());
INSERT INTO currencies (name, symbol, created_at, updated_at) VALUES ('EUR', '€', NOW(), NOW());
