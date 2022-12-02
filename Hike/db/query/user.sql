-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
  username,
  age,
  email,
  hashed_password,
  role
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET email = $2,
    hashed_password = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;