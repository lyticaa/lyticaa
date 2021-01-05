CREATE TABLE amazon_fulfillments
(
    id                      BIGSERIAL,
    amazon_fulfillment_id   UUID DEFAULT UUID_GENERATE_V4(),
    name                    VARCHAR NOT NULL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (name)
);

INSERT INTO amazon_fulfillments (name, created_at, updated_at) VALUES ('UNKNOWN', NOW(), NOW());
INSERT INTO amazon_fulfillments (name, created_at, updated_at) VALUES ('Amazon', NOW(), NOW());
