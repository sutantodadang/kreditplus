-- name: CreateUser :execresult
INSERT INTO kp_users (
  user_id, email, password
) VALUES (
  ?, ?, ?
);

-- name: GetUserByEmail :one
SELECT user_id, email, password FROM kp_users WHERE email = ?;