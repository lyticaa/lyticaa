CREATE TABLE users
(
    id          BIGSERIAL   NOT NULL,
    user_id     VARCHAR     NOT NULL,
    email       VARCHAR     NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL default CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL default CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMPTZ,
    PRIMARY KEY (id),
    UNIQUE (email, user_id)
);
