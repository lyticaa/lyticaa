CREATE MATERIALIZED VIEW amazon_users AS
    SELECT DISTINCT(user_id) FROM amazon_custom_transactions;
