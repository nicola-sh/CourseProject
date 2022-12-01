-- name: GetRoute :one
SELECT * FROM routes
WHERE id = $1 LIMIT 1;

-- name: ListRoutes :many
SELECT * FROM routes
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateRoute :one
INSERT INTO routes (
  admin_id,
  title,
  description,
  location,
  destination,
  height,
  level
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: UpdateRoute :one
UPDATE routes
SET admin_id = $2,
    title = $3,
    description = $4,
    location = $5,
    destination = $6,
    height = $7,
    level = $8
WHERE id = $1
RETURNING *;

-- name: DeleteRoute :exec
DELETE FROM routes
WHERE id = $1;