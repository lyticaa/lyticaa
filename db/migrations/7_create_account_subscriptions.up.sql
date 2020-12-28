CREATE TABLE account_subscriptions
(
    id                      BIGSERIAL NOT NULL,
    account_subscription_id UUID DEFAULT UUID_GENERATE_V4(),
    user_id                 BIGSERIAL REFERENCES users(id) ON DELETE CASCADE,
    stripe_subscription_id  VARCHAR NULL,
    stripe_plan_id          VARCHAR NULL,
    created_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (user_id)
);
