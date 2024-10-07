-- name: CreateNotification :one
INSERT INTO notifications (user_id, content)
VALUES (@user_id, @content)
RETURNING id, user_id, content, created_at;

-- name: GetNotificationsByUser :many
SELECT id, user_id, content, created_at
FROM notifications
WHERE user_id = @user_id
ORDER BY created_at DESC;