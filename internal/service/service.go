package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/nordew/warpbid/internal/infrastructure/storage"
	"github.com/nordew/warpbid/internal/models"
	"github.com/nordew/warpbid/pkg/auth"
)

type AuthJWT interface {
	GenerateAccessToken(user *models.User) (string, error)
	GenerateRefreshToken(user *models.User) (string, error)
	GenerateTokens(user *models.User) (string, string, error)
	VerifyToken(tokenStr string) (*jwt.RegisteredClaims, error)
}

type Service struct {
	storage    storage.Storage
	jwtManager auth.JWTManager
}

func New(storage storage.Storage, jwtManager auth.JWTManager) *Service {
	return &Service{
		storage:    storage,
		jwtManager: jwtManager,
	}
}
