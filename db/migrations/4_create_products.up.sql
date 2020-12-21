CREATE TABLE products
(
    id          BIGSERIAL,
    product_id  UUID DEFAULT UUID_GENERATE_V4(),
    user_id     BIGSERIAL REFERENCES users(id) ON DELETE CASCADE,
    sku         VARCHAR NOT NULL,
    marketplace VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id, sku, marketplace)
);
