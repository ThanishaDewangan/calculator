-- name: CreateUser :one
INSERT INTO users (name, dob)
VALUES ($1, $2)
RETURNING id, name, dob, created_at, updated_at;

-- name: GetUser :one
SELECT id, name, dob, created_at, updated_at FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET name = $1, dob = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $3
RETURNING id, name, dob, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, dob, created_at, updated_at FROM users
ORDER BY id ASC
LIMIT $1 OFFSET $2;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;


