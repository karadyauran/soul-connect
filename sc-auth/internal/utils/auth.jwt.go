package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}

const (
	MaxAge               = 15 * time.Minute   // 10 days
	RefreshTokenDuration = time.Hour * 24 * 7 // 7 days
)

type Token struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

func GenerateToken(userID string, secretSessionKey string) (*Token, error) {
	accessExpirationTime := time.Now().Add(MaxAge)

	accessTokenClaims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "soul-connect",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessToken, err := token.SignedString([]byte(secretSessionKey))
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshExpirationTime := time.Now().Add(RefreshTokenDuration)

	refreshTokenClaims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "soul-connect",
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenStr, err := refreshToken.SignedString([]byte(secretSessionKey))
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &Token{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenStr,
		ExpiresAt:    accessExpirationTime,
	}, nil
}

func ValidateJWT(tokenString string, secretSessionKey string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretSessionKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
