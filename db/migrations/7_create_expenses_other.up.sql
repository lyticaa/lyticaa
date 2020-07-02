CREATE TABLE expenses_other
(
    id          BIGSERIAL NOT NULL,
    user_id     VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    cost        REAL,
    date_time   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id, date_time)
);
