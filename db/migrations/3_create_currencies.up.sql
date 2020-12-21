CREATE TABLE currencies
(
    id          BIGSERIAL,
    currency_id UUID DEFAULT UUID_GENERATE_V4(),
    code        VARCHAR NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (code)
);

INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'USD', NOW(), NOW());
INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'CAD', NOW(), NOW());
INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'MXN', NOW(), NOW());
INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'BRL', NOW(), NOW());
INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'AED', NOW(), NOW());
INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'EUR', NOW(), NOW());
INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'GBP', NOW(), NOW());
INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'INR', NOW(), NOW());
INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'SAR', NOW(), NOW());
INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'TRY', NOW(), NOW());
INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'SGD', NOW(), NOW());
INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'AUD', NOW(), NOW());
INSERT INTO currencies (currency_id, code, created_at, updated_at) VALUES (uuid_generate_v4()::uuid, 'JPY', NOW(), NOW());
