package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/nordew/warpbid/internal/models"
)

type Storage interface {
	// ...
}

type AuthJWT interface {
	GenerateAccessToken(user *models.User) (string, error)
	GenerateRefreshToken(user *models.User) (string, error)
	GenerateTokens(user *models.User) (string, string, error)
	VerifyToken(tokenStr string) (*jwt.RegisteredClaims, error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}
