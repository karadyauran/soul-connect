-- name: CreateUser :one
INSERT INTO users (auth_id, full_name, bio, photo_link)
VALUES (@auth_id, @full_name, @bio, @photo_link)
RETURNING id, auth_id, full_name, bio, photo_link, created_at, updated_at;

-- name: GetUserByID :one
SELECT id, auth_id, full_name, bio, photo_link, created_at, updated_at
FROM users
WHERE id = @id
LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
SET full_name = COALESCE(@full_name, full_name),
    bio = COALESCE(@bio, bio),
    photo_link = COALESCE(@photo_link, photo_link),
    updated_at = NOW()
WHERE id = @id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = @id;