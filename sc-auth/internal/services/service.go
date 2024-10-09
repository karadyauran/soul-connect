package services

import (
	"github.com/jackc/pgx/v5/pgxpool"
	db "soul-connect/sc-auth/internal/db/sqlc"
)

type Service struct {
	AuthService *AuthService
}

func NewService(pool *pgxpool.Pool) *Service {
	queries := db.New(pool)
	return &Service{
		AuthService: NewAuthService(queries),
	}
}
