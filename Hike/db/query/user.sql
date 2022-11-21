-- name: CreateUser :one
INSERT INTO users (
  firstName,
  lastname,
  age,
  emal,
  password,
  role
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;