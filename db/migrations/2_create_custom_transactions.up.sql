CREATE TABLE transaction_types
(
    id         BIGSERIAL,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id),
    UNIQUE (NAME)
);

CREATE TABLE marketplaces
(
    id         BIGSERIAL,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id),
    UNIQUE (NAME)
);

CREATE TABLE fulfillments
(
    id         BIGSERIAL,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id),
    UNIQUE (NAME)
);

CREATE TABLE tax_collection_models
(
    id         BIGSERIAL,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id),
    UNIQUE (NAME)
);

CREATE TABLE transactions
(
    id                       BIGSERIAL NOT NULL,
    user_id                  BIGSERIAL REFERENCES users(id),
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
    deleted_at               TIMESTAMPTZ,
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

INSERT INTO fulfillments (name, created_at, updated_at) VALUES ('UNKNOWN', NOW(), NOW());
INSERT INTO fulfillments (name, created_at, updated_at) VALUES ('Amazon', NOW(), NOW());

INSERT INTO tax_collection_models (name, created_at, updated_at) VALUES ('UNKNOWN', NOW(), NOW());
INSERT INTO tax_collection_models (name, created_at, updated_at) VALUES ('MarketplaceFacilitator', NOW(), NOW());
