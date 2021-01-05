CREATE TABLE amazon_marketplaces
(
    id                      BIGSERIAL,
    amazon_marketplace_id   UUID DEFAULT UUID_GENERATE_V4(),
    name                    VARCHAR NOT NULL,
    exchange_rate_id        BIGSERIAL references exchange_rates(id),
    created_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (name)
);

INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('UNKNOWN', (SELECT id FROM exchange_rates WHERE code = 'UNKNOWN'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.com', (SELECT id FROM exchange_rates WHERE code = 'USD'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.ca', (SELECT id FROM exchange_rates WHERE code = 'CAD'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.mx', (SELECT id FROM exchange_rates WHERE code = 'MXN'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.br', (SELECT id FROM exchange_rates WHERE code = 'BRL'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.ae', (SELECT id FROM exchange_rates WHERE code = 'AED'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.de', (SELECT id FROM exchange_rates WHERE code = 'EUR'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.es', (SELECT id FROM exchange_rates WHERE code = 'EUR'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.fr', (SELECT id FROM exchange_rates WHERE code = 'EUR'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.co.uk', (SELECT id FROM exchange_rates WHERE code = 'GBP'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.in', (SELECT id FROM exchange_rates WHERE code = 'INR'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.it', (SELECT id FROM exchange_rates WHERE code = 'EUR'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.nl', (SELECT id FROM exchange_rates WHERE code = 'EUR'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.sa', (SELECT id FROM exchange_rates WHERE code = 'SAR'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.com.tr', (SELECT id FROM exchange_rates WHERE code = 'TRY'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.sg', (SELECT id FROM exchange_rates WHERE code = 'SGD'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.com.au', (SELECT id FROM exchange_rates WHERE code = 'AUD'), NOW(), NOW());
INSERT INTO amazon_marketplaces (name, exchange_rate_id, created_at, updated_at) VALUES ('amazon.co.jp', (SELECT id FROM exchange_rates WHERE code = 'JPY'), NOW(), NOW());
