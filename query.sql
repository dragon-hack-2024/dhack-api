-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY email;

-- name: CreateUser :one
INSERT INTO users (
  email, weight, height, birth_date
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET email = $1, weight = $2, height = $3, birth_date = $4
WHERE id = $5
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;