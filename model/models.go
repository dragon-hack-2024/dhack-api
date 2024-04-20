// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package model

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Challenge struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Steps    []byte `json:"steps"`
	FileName string `json:"file_name"`
	Duration int32  `json:"duration"`
}

type Stat struct {
	ID             int32            `json:"id"`
	CaloriesBurned int32            `json:"calories_burned"`
	Rpm            float32          `json:"rpm"`
	Duration       int32            `json:"duration"`
	Score          float32          `json:"score"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	ChallengeID    int32            `json:"challenge_id"`
	UserID         int32            `json:"user_id"`
}

type User struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Weight    int16            `json:"weight"`
	BirthDate pgtype.Date      `json:"birth_date"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}
