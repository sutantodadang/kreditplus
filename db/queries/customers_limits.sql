-- name: CreateCustomersLimits :copyfrom
INSERT INTO kp_customers_limits (
    customer_limit_id,
    customer_id,
    limit_amount,
    tenor
) VALUES (?, ?, ?, ?);

-- name: GetCustomerLimitById :many
SELECT customer_limit_id, customer_id, limit_amount, tenor FROM kp_customers_limits WHERE customer_id = ?;