package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
	"wtf-credential/configs"
)

func CreateToken(ctx context.Context, id uuid.UUID) (token string, err error) {
	cfg := configs.Config()
	expirationTime := time.Now().Add(24 * 2 * time.Hour)
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Subject:   id.String(),
	}
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := newToken.SignedString([]byte(cfg.JwtSecret))
	if err != nil {
		return ``, err
	}
	return tokenString, nil
}
