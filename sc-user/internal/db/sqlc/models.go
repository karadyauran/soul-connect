// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Subscription struct {
	ID           pgtype.UUID      `json:"id"`
	SubscriberID pgtype.UUID      `json:"subscriber_id"`
	AuthorID     pgtype.UUID      `json:"author_id"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
}

type User struct {
	ID        pgtype.UUID      `json:"id"`
	AuthID    pgtype.UUID      `json:"auth_id"`
	FullName  string           `json:"full_name"`
	Bio       pgtype.Text      `json:"bio"`
	PhotoLink pgtype.Text      `json:"photo_link"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
