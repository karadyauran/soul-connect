package services

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
	db "soul-connect/sc-auth/internal/db/sqlc"
	"soul-connect/sc-auth/internal/models"
	"soul-connect/sc-auth/internal/utils"
	"time"
)

type IAuthRepository interface {
	CreateUser(ctx context.Context, params db.CreateUserParams) (db.CreateUserRow, error)
	GetUserByEmail(ctx context.Context, email string) (db.GetUserByEmailRow, error)
	GetUserByUsername(ctx context.Context, username string) (db.GetUserByUsernameRow, error)
	DeleteUser(ctx context.Context, id pgtype.UUID) error
	UpdateUserPassword(ctx context.Context, params db.UpdateUserPasswordParams) error
	UpdateLastLogin(ctx context.Context, id pgtype.UUID) error
	LogFailedLogin(ctx context.Context, username string) error
	LogSuccessfulLogin(ctx context.Context, username string) error
	GetSessionByUserId(ctx context.Context, userID pgtype.UUID) ([]db.GetSessionByUserIdRow, error)
	CreateSession(ctx context.Context, arg db.CreateSessionParams) (db.Session, error)
	UpdateSessionExpiry(ctx context.Context, arg db.UpdateSessionExpiryParams) error
	DeleteSessionByToken(ctx context.Context, sessionToken string) error
	DeleteAllSessionsForUser(ctx context.Context, userID pgtype.UUID) error
}

type AuthService struct {
	authRepo IAuthRepository
}

func NewAuthService(authRepository IAuthRepository) *AuthService {
	return &AuthService{
		authRepo: authRepository,
	}
}

const (
	secret = "SECRET_JWT_KEY"
)

func (s *AuthService) Register(ctx context.Context, params models.CreateUserRequest) (*models.CreateUserResponse, error) {
	_, err := s.authRepo.GetUserByEmail(ctx, params.Email)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}

	_, err = s.authRepo.GetUserByUsername(ctx, params.Username)
	if err == nil {
		return nil, errors.New("user with this username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	newUser, err := s.authRepo.CreateUser(ctx, db.CreateUserParams{
		Username: params.Username,
		Email:    params.Email,
		Password: string(hashedPassword),
	})

	if err != nil {
		return nil, err
	}

	userID := uuid.UUID(newUser.ID.Bytes[:])
	return &models.CreateUserResponse{
		ID:       userID.String(),
		Username: newUser.Username,
		Email:    newUser.Email,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, loginCredentials models.UserLoginRequest) (*models.UserLoginResponseDTO, error) {
	user, err := s.authRepo.GetUserByUsername(ctx, loginCredentials.Username)
	if err != nil {
		return nil, err
	}

	// Comparing passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginCredentials.Password)); err != nil {
		if err := s.authRepo.LogFailedLogin(ctx, loginCredentials.Username); err != nil {
			return nil, err
		}
		return nil, errors.New("invalid credentials")
	}

	// Log successful login
	if err := s.authRepo.LogSuccessfulLogin(ctx, loginCredentials.Username); err != nil {
		return nil, err
	}

	// Update last date of login
	if err := s.authRepo.UpdateLastLogin(ctx, user.ID); err != nil {
		return nil, err
	}

	// Clean up sessions
	sessions, err := s.authRepo.GetSessionByUserId(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	activeSessions, err := s.cleanUpSessions(ctx, sessions)
	if err != nil {
		return nil, err
	}

	// Generation token
	userID := uuid.UUID(user.ID.Bytes[:])
	tokens, err := utils.GenerateToken(userID.String(), secret)
	if err != nil {
		return nil, errors.New("failed to generate tokens")
	}

	var refreshToken string
	if len(activeSessions) > 0 {
		refreshToken = activeSessions[0].SessionToken
	} else {
		_, err = s.authRepo.CreateSession(ctx, db.CreateSessionParams{
			UserID:           user.ID,
			SessionToken:     tokens.RefreshToken,
			SessionExpiresAt: pgtype.Timestamp{Time: tokens.ExpiresAt, Valid: true},
		})
		if err != nil {
			return nil, errors.New("failed to generate session")
		}
		refreshToken = tokens.RefreshToken
	}

	return &models.UserLoginResponseDTO{
		ID:           userID.String(),
		Username:     user.Username,
		Email:        user.Email,
		AccessToken:  tokens.AccessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) cleanUpSessions(ctx context.Context, sessions []db.GetSessionByUserIdRow) ([]db.GetSessionByUserIdRow, error) {
	activeSessions := make([]db.GetSessionByUserIdRow, 0)
	for _, session := range sessions {
		if session.SessionExpiresAt.Time.Before(time.Now()) {
			if err := s.authRepo.DeleteSessionByToken(ctx, session.SessionToken); err != nil {
				return nil, err
			}
		} else {
			activeSessions = append(activeSessions, session)
		}
	}
	return activeSessions, nil
}

func (s *AuthService) Logout(ctx context.Context, token string) error {
	if err := s.authRepo.DeleteSessionByToken(ctx, token); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) LogoutFromAllDevices(ctx context.Context, userID pgtype.UUID) error {
	if err := s.authRepo.DeleteAllSessionsForUser(ctx, userID); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) DeleteUser(ctx context.Context, userID pgtype.UUID) error {
	if err := s.authRepo.DeleteUser(ctx, userID); err != nil {
		return err
	}
	return nil
}
