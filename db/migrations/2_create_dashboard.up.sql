CREATE TABLE dashboard
(
    id                  BIGSERIAL,
    date_range          VARCHAR NOT NULL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    total_sales         REAL NOT NULL,
    units_sold          BIGINT NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    gross_margin        REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (date_range,user_id,date_time,marketplace)
);
