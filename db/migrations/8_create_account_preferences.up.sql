CREATE TABLE account_preferences
(
    id                      BIGSERIAL NOT NULL,
    account_preference_id   UUID DEFAULT UUID_GENERATE_V4(),
    user_id                 BIGSERIAL REFERENCES users(id) ON DELETE CASCADE,
    setup_completed         BOOLEAN NOT NULL,
    mailing_list            BOOLEAN NULL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);
