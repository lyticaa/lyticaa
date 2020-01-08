CREATE TABLE types
(
    id         BIGSERIAL NOT NULL,
    type       VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id),
    UNIQUE (type)
);

CREATE TABLE marketplaces
(
    id          BIGSERIAL NOT NULL,
    marketplace VARCHAR NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMPTZ,
    PRIMARY KEY (id),
    UNIQUE (marketplace)
);

CREATE TABLE fulfillments
(
    id          BIGSERIAL NOT NULL,
    fulfillment VARCHAR NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMPTZ,
    PRIMARY KEY (id),
    UNIQUE (fulfillment)
);

CREATE TABLE tax_collection_models
(
    id                   BIGSERIAL NOT NULL,
    tax_collection_model VARCHAR NOT NULL,
    created_at           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at           TIMESTAMPTZ,
    PRIMARY KEY (id),
    UNIQUE (tax_collection_model)
);

CREATE TABLE custom_transactions
(
    id                       BIGSERIAL NOT NULL,
    date_time                TIMESTAMPTZ NOT NULL,
    settlement_id            BIGSERIAL NOT NULL,
    type_id                  BIGSERIAL REFERENCES types(id),
    order_id                 VARCHAR NOT NULL,
    sku                      VARCHAR NOT NULL,
    quantity                 INT,
    marketplace_id           BIGSERIAL REFERENCES marketplaces(id),
    fulfillment_id           BIGSERIAL REFERENCES fulfillments(id),
    tax_collection_model_id  BIGSERIAL REFERENCES tax_collection_models(id),
    product_sales            REAL,
    product_sales_tax        REAL,
    shipping_credits         REAL,
    shipping_credits_tax     REAL,
    gift_wrap_credits        REAL,
    gift_wrap_credits_tax    REAL,
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
    deleted_at               TIMESTAMPTZ,
    PRIMARY KEY (id),
    UNIQUE (order_id, sku)
);
