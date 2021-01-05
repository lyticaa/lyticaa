CREATE TABLE amazon_transaction_types
(
    id         BIGSERIAL,
    name       VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (name)
);

INSERT INTO amazon_transaction_types (name, created_at, updated_at) VALUES ('UNKNOWN', NOW(), NOW());
INSERT INTO amazon_transaction_types (name, created_at, updated_at) VALUES ('Order', NOW(), NOW());
INSERT INTO amazon_transaction_types (name, created_at, updated_at) VALUES ('Refund', NOW(), NOW());
INSERT INTO amazon_transaction_types (name, created_at, updated_at) VALUES ('Service Fee', NOW(), NOW());
INSERT INTO amazon_transaction_types (name, created_at, updated_at) VALUES ('Adjustment', NOW(), NOW());
INSERT INTO amazon_transaction_types (name, created_at, updated_at) VALUES ('Transfer', NOW(), NOW());
INSERT INTO amazon_transaction_types (name, created_at, updated_at) VALUES ('FBA Inventory Fee', NOW(), NOW());
