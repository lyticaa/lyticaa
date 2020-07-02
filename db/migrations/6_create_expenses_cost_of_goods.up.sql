CREATE TABLE expenses_cost_of_goods
(
    id          BIGSERIAL NOT NULL,
    user_id     VARCHAR NOT NULL,
    marketplace VARCHAR NOT NULL,
    sku         VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    cost        REAL,
    from_date   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id, marketplace, sku, from_date)
);
