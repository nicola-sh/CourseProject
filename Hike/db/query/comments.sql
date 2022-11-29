-- name: GetComment :one
SELECT * FROM comments
WHERE id = $1 LIMIT 1;

-- name: ListComments :many
SELECT * FROM comments
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateComment :one
INSERT INTO comments (
  post_id,
  user_id,
  comment_text
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateComment :one
UPDATE comments
SET comment_text = $2
WHERE id = $1
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;