CREATE TABLE exchange_rates
(
    id                  BIGSERIAL,
    exchange_rate_id    UUID DEFAULT UUID_GENERATE_V4(),
    code                VARCHAR NOT NULL,
    rate                REAL,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('UNKNOWN', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('USD', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('CAD', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('MXN', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('BRL', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('AED', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('EUR', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('GBP', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('INR', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('SAR', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('TRY', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('SGD', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('AUD', 1.00, NOW(), NOW());
INSERT INTO exchange_rates (code, rate, created_at, updated_at) VALUES ('JPY', 1.00, NOW(), NOW());
