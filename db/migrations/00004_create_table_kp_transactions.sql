-- +goose Up
-- +goose StatementBegin
CREATE TABLE kp_transactions (
    transaction_id CHAR(42) PRIMARY KEY,
    customer_id CHAR(42) NOT NULL,
    customer_limit_id CHAR(42) NOT NULL,
    contract_number VARCHAR(255) NOT NULL,
    otr_amount DECIMAL NOT NULL,
    admin_fee DECIMAL NOT NULL,
    installment_amount DECIMAL NOT NULL,
    interest_amount DECIMAL NOT NULL,
    asset_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES kp_customers(customer_id),
    FOREIGN KEY (customer_limit_id) REFERENCES kp_customers_limits(customer_limit_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE kp_transactions;
-- +goose StatementEnd
