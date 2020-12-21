CREATE TABLE account_notifications
(
    id                      BIGSERIAL,
    account_notification_id UUID DEFAULT UUID_GENERATE_V4(),
    user_id                 BIGSERIAL REFERENCES users(id) ON DELETE CASCADE,
    notification            VARCHAR NOT NULL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
