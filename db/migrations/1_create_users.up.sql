CREATE TABLE users
(
    id                      BIGSERIAL NOT NULL,
    user_id                 VARCHAR NOT NULL,
    email                   VARCHAR NOT NULL,
    stripe_user_id          VARCHAR NULL,
    stripe_subscription_id  VARCHAR NULL,
    stripe_plan_id          VARCHAR NULL,
    nickname                VARCHAR NULL,
    avatar_url              VARCHAR NULL,
    setup_completed         BOOLEAN NOT NULL,
    mailing_list            BOOLEAN NOT NULL,
    admin                   BOOLEAN NOT NULL DEFAULT false,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id, email)
);
