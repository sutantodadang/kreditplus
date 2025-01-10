-- name: CreateCustomers :exec
INSERT INTO kp_customers (
  customer_id, identification_number, full_name, legal_name, place_of_birth, date_of_birth, salary, photo_ktp, photo_selfie, user_id
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

