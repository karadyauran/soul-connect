-- name: CreateLikeForPost :exec
INSERT INTO likes (post_id, user_id)
VALUES (@post_id, @user_id)
RETURNING id, post_id, user_id, created_at;

-- name: DeleteLikeForPost :exec
DELETE FROM likes
WHERE post_id = @post_id AND user_id = @user_id;

-- name: CreateLikeForComment :exec
INSERT INTO likes (comment_id, user_id)
VALUES (@comment_id, @user_id)
RETURNING id, comment_id, user_id, created_at;

-- name: DeleteLikeForComment :exec
DELETE FROM likes
WHERE comment_id = @comment_id AND user_id = @user_id;

-- name: GetLikesCountForPost :one
SELECT likes_count
FROM posts
WHERE id = @post_id;

-- name: GetLikesCountForComment :one
SELECT likes_count
FROM comments
WHERE id = @comment_id;
