-- name: GetCustomerTransactionByLimitIdAndCustomerId :many
SELECT a.transaction_id, a.customer_id, a.customer_limit_id, a.contract_number, a.otr_amount, a.admin_fee, 
a.installment_amount, a.interest_amount, a.asset_name
FROM kp_transactions a
WHERE a.created_at >= DATE_FORMAT(NOW(), '%Y-%m-01')
   AND a.created_at < DATE_FORMAT(DATE_ADD(NOW(), INTERVAL 1 MONTH), '%Y-%m-01');

-- name: GetCustomerTransactionOtr :one
SELECT COALESCE(SUM(a.otr_amount),0) as history_transaction_limit
FROM kp_transactions a
JOIN kp_customers_limits b ON a.customer_id = b.customer_id
WHERE a.customer_id = ? AND a.customer_limit_id = ?;


-- name: CreateTransaction :exec
INSERT INTO kp_transactions (
    transaction_id,
    customer_id,
    customer_limit_id,
    contract_number,
    otr_amount,
    admin_fee,
    installment_amount,
    interest_amount,
    asset_name
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);