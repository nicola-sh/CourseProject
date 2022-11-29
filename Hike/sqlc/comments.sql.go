// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: comments.sql

package hikedb

import (
	"context"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (
  post_id,
  user_id,
  comment_text
) VALUES (
  $1, $2, $3
)
RETURNING id, post_id, user_id, comment_text, "createdAt", "updatedAt"
`

type CreateCommentParams struct {
	PostID      int32  `json:"post_id"`
	UserID      int32  `json:"user_id"`
	CommentText string `json:"comment_text"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment, arg.PostID, arg.UserID, arg.CommentText)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.CommentText,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteComment, id)
	return err
}

const getComment = `-- name: GetComment :one
SELECT id, post_id, user_id, comment_text, "createdAt", "updatedAt" FROM comments
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetComment(ctx context.Context, id int32) (Comment, error) {
	row := q.db.QueryRowContext(ctx, getComment, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.CommentText,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listComments = `-- name: ListComments :many
SELECT id, post_id, user_id, comment_text, "createdAt", "updatedAt" FROM comments
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCommentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListComments(ctx context.Context, arg ListCommentsParams) ([]Comment, error) {
	rows, err := q.db.QueryContext(ctx, listComments, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Comment
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.CommentText,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateComment = `-- name: UpdateComment :one
UPDATE comments
SET comment_text = $2
WHERE id = $1
RETURNING id, post_id, user_id, comment_text, "createdAt", "updatedAt"
`

type UpdateCommentParams struct {
	ID          int32  `json:"id"`
	CommentText string `json:"comment_text"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, updateComment, arg.ID, arg.CommentText)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.CommentText,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}