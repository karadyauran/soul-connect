// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Notification struct {
	ID        pgtype.UUID      `json:"id"`
	UserID    pgtype.UUID      `json:"user_id"`
	Content   string           `json:"content"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}
