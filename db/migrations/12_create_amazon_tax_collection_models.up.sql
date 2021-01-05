CREATE TABLE amazon_tax_collection_models
(
    id                              BIGSERIAL,
    amazon_tax_collection_model_id  UUID DEFAULT UUID_GENERATE_V4(),
    name                            VARCHAR NOT NULL,
    created_at                      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at                      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (name)
);

INSERT INTO amazon_tax_collection_models (name, created_at, updated_at) VALUES ('UNKNOWN', NOW(), NOW());
INSERT INTO amazon_tax_collection_models (name, created_at, updated_at) VALUES ('MarketplaceFacilitator', NOW(), NOW());
