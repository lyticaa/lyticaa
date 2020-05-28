CREATE TABLE notifications
(
    id              BIGSERIAL,
    user_id         BIGSERIAL REFERENCES users(id),
    notification    VARCHAR NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
