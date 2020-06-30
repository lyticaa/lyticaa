-- High Margin

CREATE TABLE cohorts_high_margin_today
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_yesterday
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_last_thirty_days
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_previous_thirty_days
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_this_month
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_last_month
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_month_before_last
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_last_three_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_previous_three_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_last_six_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_previous_six_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_this_year
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_last_year
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_high_margin_all_time
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Low Margin

CREATE TABLE cohorts_low_margin_today
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_yesterday
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_last_thirty_days
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_previous_thirty_days
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_this_month
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_last_month
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_month_before_last
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_last_three_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_previous_three_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_last_six_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_previous_six_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_this_year
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_last_year
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_low_margin_all_time
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Negative Margin

CREATE TABLE cohorts_negative_margin_today
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_yesterday
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_last_thirty_days
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_previous_thirty_days
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_this_month
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_last_month
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_month_before_last
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_last_three_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_previous_three_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_last_six_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_previous_six_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_this_year
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_last_year
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE cohorts_negative_margin_all_time
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    quantity            BIGINT NOT NULL,
    total_sales         REAL NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    advertising_spend   REAL NOT NULL,
    refunds             REAL NOT NULL,
    shipping_credits    REAL NOT NULL,
    promotional_rebates REAL NOT NULL,
    total_costs         REAL NOT NULL,
    net_margin          REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);
