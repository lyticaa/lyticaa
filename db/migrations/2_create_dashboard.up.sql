CREATE TABLE dashboard_today
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);

CREATE TABLE dashboard_yesterday
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);

CREATE TABLE dashboard_last_thirty_days
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);

CREATE TABLE dashboard_this_month
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);

CREATE TABLE dashboard_last_month
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);

CREATE TABLE dashboard_month_before_last
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);

CREATE TABLE dashboard_last_three_months
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);

CREATE TABLE dashboard_previous_three_months
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);

CREATE TABLE dashboard_last_six_months
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);

CREATE TABLE dashboard_previous_six_months
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);

CREATE TABLE dashboard_this_year
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);

CREATE TABLE dashboard_last_year
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);

CREATE TABLE dashboard_all_time
(
    id                  BIGSERIAL,
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
    UNIQUE (user_id)
);
