CREATE TABLE users
(
    id                      BIGSERIAL NOT NULL,
    user_id                 VARCHAR NOT NULL,
    email                   VARCHAR NOT NULL,
    nickname                VARCHAR NULL,
    avatar_url              VARCHAR NULL,
    admin                   BOOLEAN NOT NULL DEFAULT false,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id, email)
);
