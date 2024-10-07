-- name: CreateUser :one
INSERT INTO auth (username, email, password)
VALUES (@username, @email, @password)
    RETURNING id, username, email, password;

-- name: GetUserByUsername :one
SELECT id, username, email, password
FROM auth
WHERE username = @username
LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, username, email, password
FROM auth
WHERE email = @email
    LIMIT 1;

-- name: GetUserById :one
SELECT id, username, email, password, last_login
FROM auth
WHERE id = @id
    LIMIT 1;

-- name: UpdateUserPassword :exec
UPDATE auth
SET password = @new_password, updated_at = NOW()
WHERE id = @id;

-- name: UpdateLastLogin :exec
UPDATE auth
SET last_login = NOW()
WHERE id = @id;

-- name: DeleteUser :exec
DELETE FROM auth
WHERE id = @id;