-- Total Sales

CREATE TABLE metrics_total_sales_today
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_yesterday
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_last_thirty_days
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_previous_thirty_days
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_this_month
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_last_month
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_month_before_last
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_last_three_months
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_previous_three_months
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_last_six_months
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_previous_six_months
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_this_year
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_last_year
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_total_sales_all_time
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_today
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_yesterday
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_last_thirty_days
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_previous_thirty_days
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_this_month
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_last_month
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_month_before_last
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_last_three_months
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_previous_three_months
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_last_six_months
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_previous_six_months
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_this_year
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_last_year
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_units_sold_all_time
(
    id          BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_today
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_yesterday
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_last_thirty_days
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_this_month
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_last_month
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_month_before_last
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_last_three_months
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_previous_three_months
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_last_six_months
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_previous_six_months
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_this_year
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_last_year
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_amazon_costs_all_time
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_product_costs_today
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_yesterday
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_last_thirty_days
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_previous_thirty_days
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_this_month
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_last_month
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_month_before_last
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_last_three_months
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_previous_three_months
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_last_six_months
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_previous_six_months
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_this_year
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_last_year
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_product_costs_all_time
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_today
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_yesterday
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_last_thirty_days
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_previous_thirty_days
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_this_month
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_last_month
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_month_before_last
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_last_three_months
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_previous_three_months
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_last_six_months
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_previous_six_months
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_this_year
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_last_year
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_advertising_spend_all_time
(
    id                              BIGSERIAL,
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

CREATE TABLE metrics_refunds_today
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_yesterday
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_last_thirty_days
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_previous_thirty_days
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_this_month
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_last_month
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_month_before_last
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_last_three_months
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_previous_three_months
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_last_six_months
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_previous_six_months
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_this_year
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_last_year
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_refunds_all_time
(
    id                  BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_today
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_yesterday
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_last_thirty_days
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_previous_thirty_days
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_this_month
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_last_month
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_month_before_last
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_last_three_months
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_previous_three_months
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_last_six_months
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_previous_six_months
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_this_year
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_last_year
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_shipping_credits_all_time
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_today
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_yesterday
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_last_thirty_days
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_previous_thirty_days
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_this_month
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_last_month
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_month_before_last
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_last_three_months
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_previous_three_months
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_last_six_months
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_previous_six_months
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_this_year
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_last_year
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_promotional_rebates_all_time
(
    id                      BIGSERIAL,
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

CREATE TABLE metrics_total_costs_today
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_yesterday
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_last_thirty_days
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_previous_thirty_days
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_this_month
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_last_month
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_month_before_last
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_last_three_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_previous_three_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_last_six_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_previous_six_months
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_this_year
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_last_year
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_total_costs_all_time
(
    id                  BIGSERIAL,
    user_id             VARCHAR NOT NULL,
    date_time           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    marketplace         VARCHAR NOT NULL,
    sku                 VARCHAR NOT NULL,
    description         VARCHAR NOT NULL,
    amazon_costs        REAL NOT NULL,
    product_costs       REAL NOT NULL,
    product_costs_unit  REAL NOT NULL,
    total_costs         REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Gross Margin

CREATE TABLE metrics_gross_margin_today
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_yesterday
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_last_thirty_days
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_previous_thirty_days
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_this_month
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_last_month
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_month_before_last
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_last_three_months
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_previous_three_months
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_last_six_months
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_previous_six_months
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_this_year
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_last_year
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

CREATE TABLE metrics_gross_margin_all_time
(
    id                  BIGSERIAL,
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
    sales_tax_collected REAL NOT NULL,
    total_collected     REAL NOT NULL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);

-- Net Margin

CREATE TABLE metrics_net_margin_today
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_yesterday
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_last_thirty_days
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_previous_thirty_days
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_this_month
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_last_month
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_month_before_last
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_last_three_months
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_previous_three_months
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_last_six_months
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_previous_six_months
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_this_year
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_last_year
(
    id              BIGSERIAL,
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

CREATE TABLE metrics_net_margin_all_time
(
    id              BIGSERIAL,
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
