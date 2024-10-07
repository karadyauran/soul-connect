// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: post.query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (title, description)
VALUES ($1, $2)
RETURNING user_id, title, description, likes_count, created_at, updated_at
`

type CreatePostParams struct {
	Title       string      `json:"title"`
	Description pgtype.Text `json:"description"`
}

type CreatePostRow struct {
	UserID      pgtype.UUID      `json:"user_id"`
	Title       string           `json:"title"`
	Description pgtype.Text      `json:"description"`
	LikesCount  pgtype.Int4      `json:"likes_count"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (CreatePostRow, error) {
	row := q.db.QueryRow(ctx, createPost, arg.Title, arg.Description)
	var i CreatePostRow
	err := row.Scan(
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.LikesCount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deletePost, id)
	return err
}

const getPostByID = `-- name: GetPostByID :one
SELECT id, user_id, title, description, likes_count, created_at, updated_at
FROM posts
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetPostByID(ctx context.Context, id pgtype.UUID) (Post, error) {
	row := q.db.QueryRow(ctx, getPostByID, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.LikesCount,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPostsWithCommentsAndLikes = `-- name: GetPostsWithCommentsAndLikes :many
SELECT
    p.id AS post_id,
    p.title AS post_title,
    p.description AS post_description,
    p.likes_count AS post_likes,
    COUNT(DISTINCT c.id) AS total_comments,
    COUNT(DISTINCT l.id) AS total_likes
FROM posts p
         LEFT JOIN comments c ON p.id = c.post_id
         LEFT JOIN likes l ON p.id = l.post_id
GROUP BY p.id
ORDER BY p.created_at DESC
`

type GetPostsWithCommentsAndLikesRow struct {
	PostID          pgtype.UUID `json:"post_id"`
	PostTitle       string      `json:"post_title"`
	PostDescription pgtype.Text `json:"post_description"`
	PostLikes       pgtype.Int4 `json:"post_likes"`
	TotalComments   int64       `json:"total_comments"`
	TotalLikes      int64       `json:"total_likes"`
}

func (q *Queries) GetPostsWithCommentsAndLikes(ctx context.Context) ([]GetPostsWithCommentsAndLikesRow, error) {
	rows, err := q.db.Query(ctx, getPostsWithCommentsAndLikes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsWithCommentsAndLikesRow{}
	for rows.Next() {
		var i GetPostsWithCommentsAndLikesRow
		if err := rows.Scan(
			&i.PostID,
			&i.PostTitle,
			&i.PostDescription,
			&i.PostLikes,
			&i.TotalComments,
			&i.TotalLikes,
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

const updatePost = `-- name: UpdatePost :exec
UPDATE posts
SET title = COALESCE($1, title),
    description = COALESCE($2, description),
    updated_at = NOW()
WHERE id = $3
`

type UpdatePostParams struct {
	Title       string      `json:"title"`
	Description pgtype.Text `json:"description"`
	ID          pgtype.UUID `json:"id"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) error {
	_, err := q.db.Exec(ctx, updatePost, arg.Title, arg.Description, arg.ID)
	return err
}
