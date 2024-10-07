-- name: CreateSession :one
INSERT INTO sessions (user_id, session_token, session_expires_at)
VALUES (@user_id, @session_token, @session_expires_at)
    RETURNING id, user_id, session_token, session_expires_at;

-- name: GetSessionByToken :one
SELECT user_id, session_token, session_expires_at
FROM sessions
WHERE session_token = @session_token
  AND session_expires_at > NOW()
    LIMIT 1;

-- name: GetSessionByUserId :many
SELECT id, session_token, session_expires_at
FROM sessions
WHERE user_id = @user_id
  AND session_expires_at > NOW();

-- name: UpdateSessionExpiry :exec
UPDATE sessions
SET session_expires_at = @new_expiry
WHERE session_token = @session_token;

-- name: DeleteSessionByToken :exec
DELETE FROM sessions
WHERE session_token = @session_token;

-- name: DeleteAllSessionsForUser :exec
DELETE FROM sessions
WHERE user_id = @user_id;