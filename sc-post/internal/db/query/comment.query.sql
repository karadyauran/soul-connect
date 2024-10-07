-- name: CreateComment :one
INSERT INTO comments (post_id, user_id, content)
VALUES (@post_id, @user_id, @content)
RETURNING id, post_id, user_id, content, likes_count, created_at, updated_at;

-- name: GetCommentsByPostID :many
SELECT id, post_id, user_id, content, likes_count, created_at, updated_at
FROM comments
WHERE post_id = @post_id
ORDER BY created_at;