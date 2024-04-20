// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package model

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createChallenge = `-- name: CreateChallenge :one
INSERT INTO challenges (
  name, steps, file_name, duration
) VALUES (
  $1, $2, $3, $4
) RETURNING id, name, steps, file_name, duration
`

type CreateChallengeParams struct {
	Name     string `json:"name"`
	Steps    []byte `json:"steps"`
	FileName string `json:"file_name"`
	Duration int32  `json:"duration"`
}

func (q *Queries) CreateChallenge(ctx context.Context, arg CreateChallengeParams) (Challenge, error) {
	row := q.db.QueryRow(ctx, createChallenge,
		arg.Name,
		arg.Steps,
		arg.FileName,
		arg.Duration,
	)
	var i Challenge
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Steps,
		&i.FileName,
		&i.Duration,
	)
	return i, err
}

const createStat = `-- name: CreateStat :one
INSERT INTO stats (
  calories_burned, rpm, duration, score, challenge_id, user_id
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, calories_burned, rpm, duration, score, created_at, challenge_id, user_id
`

type CreateStatParams struct {
	CaloriesBurned int32   `json:"calories_burned"`
	Rpm            float32 `json:"rpm"`
	Duration       int32   `json:"duration"`
	Score          float32 `json:"score"`
	ChallengeID    int32   `json:"challenge_id"`
	UserID         int32   `json:"user_id"`
}

func (q *Queries) CreateStat(ctx context.Context, arg CreateStatParams) (Stat, error) {
	row := q.db.QueryRow(ctx, createStat,
		arg.CaloriesBurned,
		arg.Rpm,
		arg.Duration,
		arg.Score,
		arg.ChallengeID,
		arg.UserID,
	)
	var i Stat
	err := row.Scan(
		&i.ID,
		&i.CaloriesBurned,
		&i.Rpm,
		&i.Duration,
		&i.Score,
		&i.CreatedAt,
		&i.ChallengeID,
		&i.UserID,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email, weight, birth_date
) VALUES (
  $1, $2, $3
) RETURNING id, email, weight, birth_date, created_at
`

type CreateUserParams struct {
	Email     string      `json:"email"`
	Weight    int16       `json:"weight"`
	BirthDate pgtype.Date `json:"birth_date"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.Weight, arg.BirthDate)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Weight,
		&i.BirthDate,
		&i.CreatedAt,
	)
	return i, err
}

const deleteChallenge = `-- name: DeleteChallenge :exec
DELETE FROM challenges
WHERE id = $1
`

func (q *Queries) DeleteChallenge(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteChallenge, id)
	return err
}

const deleteStat = `-- name: DeleteStat :exec
DELETE FROM stats
WHERE id = $1
`

func (q *Queries) DeleteStat(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteStat, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getChallenge = `-- name: GetChallenge :one

SELECT id, name, steps, file_name, duration FROM challenges
WHERE id = $1 LIMIT 1
`

// Challenges
func (q *Queries) GetChallenge(ctx context.Context, id int32) (Challenge, error) {
	row := q.db.QueryRow(ctx, getChallenge, id)
	var i Challenge
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Steps,
		&i.FileName,
		&i.Duration,
	)
	return i, err
}

const getStat = `-- name: GetStat :one

SELECT id, calories_burned, rpm, duration, score, created_at, challenge_id, user_id FROM stats
WHERE id = $1 LIMIT 1
`

// Stats
func (q *Queries) GetStat(ctx context.Context, id int32) (Stat, error) {
	row := q.db.QueryRow(ctx, getStat, id)
	var i Stat
	err := row.Scan(
		&i.ID,
		&i.CaloriesBurned,
		&i.Rpm,
		&i.Duration,
		&i.Score,
		&i.CreatedAt,
		&i.ChallengeID,
		&i.UserID,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, email, weight, birth_date, created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Weight,
		&i.BirthDate,
		&i.CreatedAt,
	)
	return i, err
}

const listChallenges = `-- name: ListChallenges :many
SELECT id, name, steps, file_name, duration FROM challenges
ORDER BY name
`

func (q *Queries) ListChallenges(ctx context.Context) ([]Challenge, error) {
	rows, err := q.db.Query(ctx, listChallenges)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Challenge
	for rows.Next() {
		var i Challenge
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Steps,
			&i.FileName,
			&i.Duration,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listStats = `-- name: ListStats :many
SELECT id, calories_burned, rpm, duration, score, created_at, challenge_id, user_id FROM stats
ORDER BY score
`

func (q *Queries) ListStats(ctx context.Context) ([]Stat, error) {
	rows, err := q.db.Query(ctx, listStats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Stat
	for rows.Next() {
		var i Stat
		if err := rows.Scan(
			&i.ID,
			&i.CaloriesBurned,
			&i.Rpm,
			&i.Duration,
			&i.Score,
			&i.CreatedAt,
			&i.ChallengeID,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, weight, birth_date, created_at FROM users
ORDER BY email
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Weight,
			&i.BirthDate,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateChallenge = `-- name: UpdateChallenge :one
UPDATE challenges
SET name = $1, steps = $2, file_name = $3, duration = $4
WHERE id = $5
RETURNING id, name, steps, file_name, duration
`

type UpdateChallengeParams struct {
	Name     string `json:"name"`
	Steps    []byte `json:"steps"`
	FileName string `json:"file_name"`
	Duration int32  `json:"duration"`
	ID       int32  `json:"id"`
}

func (q *Queries) UpdateChallenge(ctx context.Context, arg UpdateChallengeParams) (Challenge, error) {
	row := q.db.QueryRow(ctx, updateChallenge,
		arg.Name,
		arg.Steps,
		arg.FileName,
		arg.Duration,
		arg.ID,
	)
	var i Challenge
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Steps,
		&i.FileName,
		&i.Duration,
	)
	return i, err
}

const updateStat = `-- name: UpdateStat :one
UPDATE stats
SET calories_burned = $1, rpm = $2, duration = $3, score = $4, challenge_id = $5, user_id = $6
WHERE id = $7
RETURNING id, calories_burned, rpm, duration, score, created_at, challenge_id, user_id
`

type UpdateStatParams struct {
	CaloriesBurned int32   `json:"calories_burned"`
	Rpm            float32 `json:"rpm"`
	Duration       int32   `json:"duration"`
	Score          float32 `json:"score"`
	ChallengeID    int32   `json:"challenge_id"`
	UserID         int32   `json:"user_id"`
	ID             int32   `json:"id"`
}

func (q *Queries) UpdateStat(ctx context.Context, arg UpdateStatParams) (Stat, error) {
	row := q.db.QueryRow(ctx, updateStat,
		arg.CaloriesBurned,
		arg.Rpm,
		arg.Duration,
		arg.Score,
		arg.ChallengeID,
		arg.UserID,
		arg.ID,
	)
	var i Stat
	err := row.Scan(
		&i.ID,
		&i.CaloriesBurned,
		&i.Rpm,
		&i.Duration,
		&i.Score,
		&i.CreatedAt,
		&i.ChallengeID,
		&i.UserID,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET email = $1, weight = $2, birth_date = $3
WHERE id = $4
RETURNING id, email, weight, birth_date, created_at
`

type UpdateUserParams struct {
	Email     string      `json:"email"`
	Weight    int16       `json:"weight"`
	BirthDate pgtype.Date `json:"birth_date"`
	ID        int32       `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Email,
		arg.Weight,
		arg.BirthDate,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Weight,
		&i.BirthDate,
		&i.CreatedAt,
	)
	return i, err
}
