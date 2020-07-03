CREATE TABLE expenses_others
(
    id          BIGSERIAL NOT NULL,
    expense_id  UUID DEFAULT UUID_GENERATE_V4(),
    user_id     VARCHAR NOT NULL,
    currency_id BIGSERIAL REFERENCES currencies(id),
    description VARCHAR NOT NULL,
    amount      REAL,
    date_time   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id, date_time)
);
