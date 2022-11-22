-- name: CreateUser :one
INSERT INTO users (
  first_name,
  last_name,
  age,
  emal,
  password,
  role
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;