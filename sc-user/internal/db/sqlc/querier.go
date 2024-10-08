// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateSubscription(ctx context.Context, arg CreateSubscriptionParams) error
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteSubscription(ctx context.Context, arg DeleteSubscriptionParams) error
	DeleteUser(ctx context.Context, id pgtype.UUID) error
	GetSubscriptionsByUserID(ctx context.Context, subscriberID pgtype.UUID) ([]pgtype.UUID, error)
	GetUserByID(ctx context.Context, id pgtype.UUID) (User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
