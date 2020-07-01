-- Total Sales

CREATE TABLE metrics_total_sales
(
    id          BIGSERIAL,
    date_range  VARCHAR NOT NULL,
    user_id     VARCHAR NOT NULL,
    date_time   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace VARCHAR NOT NULL,
    sku         VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    total_sales REAL NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Units Sold

CREATE TABLE metrics_units_sold
(
    id          BIGSERIAL,
    date_range  VARCHAR NOT NULL,
    user_id     VARCHAR NOT NULL,
    date_time   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace VARCHAR NOT NULL,
    sku         VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    quantity    BIGINT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Amazon Costs

CREATE TABLE metrics_amazon_costs
(
    id              BIGSERIAL,
    date_range      VARCHAR NOT NULL,
    user_id         VARCHAR NOT NULL,
    date_time       TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace     VARCHAR NOT NULL,
    sku             VARCHAR NOT NULL,
    description     VARCHAR NOT NULL,
    amazon_costs    REAL NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Product Costs

CREATE TABLE metrics_product_costs
(
    id                  BIGSERIAL,
    date_range          VARCHAR NOT NULL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    cost_of_goods       REAL NOT NULL,
    quantity            BIGINT NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Advertising Spend

CREATE TABLE metrics_advertising_spend
(
    id                              BIGSERIAL,
    date_range                      VARCHAR NOT NULL,
    user_id                         VARCHAR NOT NULL,
    date_time                       TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace                     VARCHAR NOT NULL,
    sku                             VARCHAR NOT NULL,
    description                     VARCHAR NOT NULL,
    advertising_spend               REAL NOT NULL,
    advertising_spend_percentage    REAL NOT NULL,
    created_at                      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at                      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Refunds

CREATE TABLE metrics_refunds
(
    id                  BIGSERIAL,
    date_range          VARCHAR NOT NULL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    refunds             REAL NOT NULL,
    refunds_percentage  REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Shipping Credits

CREATE TABLE metrics_shipping_credits
(
    id                      BIGSERIAL,
    date_range              VARCHAR NOT NULL,
    user_id                 VARCHAR NOT NULL,
    date_time               TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace             VARCHAR NOT NULL,
    sku                     VARCHAR NOT NULL,
    description             VARCHAR NOT NULL,
    shipping_credits        REAL NOT NULL,
    shipping_credits_tax    REAL NOT NULL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Promotional Rebates

CREATE TABLE metrics_promotional_rebates
(
    id                      BIGSERIAL,
    date_range              VARCHAR NOT NULL,
    user_id                 VARCHAR NOT NULL,
    date_time               TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace             VARCHAR NOT NULL,
    sku                     VARCHAR NOT NULL,
    description             VARCHAR NOT NULL,
    promotional_rebates     REAL NOT NULL,
    promotional_rebates_tax REAL NOT NULL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Total Costs

CREATE TABLE metrics_total_costs
(
    id                      BIGSERIAL,
    date_range              VARCHAR NOT NULL,
    user_id                 VARCHAR NOT NULL,
    date_time               TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace             VARCHAR NOT NULL,
    sku                     VARCHAR NOT NULL,
    description             VARCHAR NOT NULL,
    amazon_costs            REAL NOT NULL,
    product_costs           REAL NOT NULL,
    product_costs_unit      REAL NOT NULL,
    total_costs             REAL NOT NULL,
    total_costs_percentage  REAL NOT NULL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Gross Margin

CREATE TABLE metrics_gross_margin
(
    id                  BIGSERIAL,
    date_range          VARCHAR NOT NULL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    product_costs       REAL NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    gross_margin        REAL NOT NULL,
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Net Margin

CREATE TABLE metrics_net_margin
(
    id              BIGSERIAL,
    date_range      VARCHAR NOT NULL,
    user_id         VARCHAR NOT NULL,
    date_time       TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace     VARCHAR NOT NULL,
    sku             VARCHAR NOT NULL,
    description     VARCHAR NOT NULL,
    quantity        BIGINT NOT NULL,
    gross_margin    REAL NOT NULL,
    total_costs     REAL NOT NULL,
    net_margin      REAL NOT NULL,
    net_margin_unit REAL NOT NULL,
    roi             REAL NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);
