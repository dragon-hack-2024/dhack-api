-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
  name, email, weight, birth_date
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET name = $1, email = $2, weight = $3, birth_date = $4
WHERE id = $5
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- Challenges

-- name: GetChallenge :one
SELECT * FROM challenges
WHERE id = $1 LIMIT 1;

-- name: ListChallenges :many
SELECT * FROM challenges
ORDER BY name;

-- name: CreateChallenge :one
INSERT INTO challenges (
  name, steps, file_name, duration
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: UpdateChallenge :one
UPDATE challenges
SET name = $1, steps = $2, file_name = $3, duration = $4
WHERE id = $5
RETURNING *;

-- name: DeleteChallenge :exec
DELETE FROM challenges
WHERE id = $1;

-- Stats

-- name: GetStat :one
SELECT * FROM stats
WHERE id = $1 LIMIT 1;

-- name: ListStats :many
SELECT * FROM stats
ORDER BY score;

-- name: CreateStat :one
INSERT INTO stats (
  calories_burned, rpm, duration, score, challenge_id, user_id
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: UpdateStat :one
UPDATE stats
SET calories_burned = $1, rpm = $2, duration = $3, score = $4, challenge_id = $5, user_id = $6
WHERE id = $7
RETURNING *;

-- name: DeleteStat :exec
DELETE FROM stats
WHERE id = $1;

-- name: GetWeekyProgress :many
SELECT 
  CAST(created_at AS DATE) AS date,
  CAST(SUM(rpm * (duration / 60.0)) AS INTEGER) AS total_skips
FROM stats 
WHERE created_at >= now() - INTERVAL '7 days'
  AND user_id = $1
GROUP BY date
ORDER BY date;