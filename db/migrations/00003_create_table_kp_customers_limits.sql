-- +goose Up
-- +goose StatementBegin
CREATE TABLE kp_customers_limits (
    customer_limit_id CHAR(42) PRIMARY KEY,
    customer_id CHAR(42) NOT NULL,
    limit_amount DECIMAL NOT NULL,
    tenor INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES kp_customers(customer_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE kp_customers_limits;
-- +goose StatementEnd
