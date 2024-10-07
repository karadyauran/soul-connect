-- name: LogSuccessfulLogin :exec
INSERT INTO login_attempts (username, success)
VALUES (@username, TRUE);

-- name: LogFailedLogin :exec
INSERT INTO login_attempts (username, success)
VALUES (@username, FALSE);

-- name: GetLoginAttemptsByUsername :many
SELECT id, username, success, attempt_time
FROM login_attempts
WHERE username = @username
ORDER BY attempt_time DESC
LIMIT @request_limit;

-- name: GetFailedLoginAttemptsByUsername :many
SELECT id, username, success, attempt_time
FROM login_attempts
WHERE username = @username
  AND success = FALSE
ORDER BY attempt_time DESC
LIMIT @request_limit;

-- name: DeleteOldLoginAttempts :exec
DELETE FROM login_attempts
WHERE attempt_time < NOW() - INTERVAL '30 days';