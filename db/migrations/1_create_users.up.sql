CREATE TABLE users
(
    id                      BIGSERIAL NOT NULL,
    user_id                 VARCHAR NOT NULL,
    stripe_user_id          VARCHAR NULL,
    stripe_subscription_id  VARCHAR NULL,
    stripe_plan_id          VARCHAR NULL,
    email                   VARCHAR NOT NULL,
    first_name              VARCHAR NULL,
    company_name            VARCHAR NULL,
    setup_completed         BOOLEAN NOT NULL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (email, user_id)
);
