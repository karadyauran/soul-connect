// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: comment.query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (post_id, user_id, content)
VALUES ($1, $2, $3)
RETURNING id, post_id, user_id, content, likes_count, created_at, updated_at
`

type CreateCommentParams struct {
	PostID  pgtype.UUID `json:"post_id"`
	UserID  pgtype.UUID `json:"user_id"`
	Content string      `json:"content"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, createComment, arg.PostID, arg.UserID, arg.Content)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.UserID,
		&i.Content,
		&i.LikesCount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCommentsByPostID = `-- name: GetCommentsByPostID :many
SELECT id, post_id, user_id, content, likes_count, created_at, updated_at
FROM comments
WHERE post_id = $1
ORDER BY created_at
`

func (q *Queries) GetCommentsByPostID(ctx context.Context, postID pgtype.UUID) ([]Comment, error) {
	rows, err := q.db.Query(ctx, getCommentsByPostID, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.UserID,
			&i.Content,
			&i.LikesCount,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
