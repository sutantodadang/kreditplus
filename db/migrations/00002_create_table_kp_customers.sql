-- +goose Up
-- +goose StatementBegin
CREATE TABLE kp_customers (
    customer_id CHAR(42) PRIMARY KEY,
    identification_number VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    legal_name VARCHAR(255) NOT NULL,
    place_of_birth VARCHAR(255) NOT NULL,
    date_of_birth DATE NOT NULL,
    salary DECIMAL NOT NULL,
    photo_ktp VARCHAR(255) NOT NULL,
    photo_selfie VARCHAR(255) NOT NULL,
    user_id CHAR(42) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES kp_users(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE kp_customers;
-- +goose StatementEnd
