package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nordew/go-errx"
	"github.com/nordew/warpbid/internal/models"
)

var (
	ErrFailedToGenerateAccessToken  = "failed to generate access token"
	ErrFailedToGenerateRefreshToken = "failed to generate refresh token"
	ErrFailedToVerifyToken          = "failed to verify token"
	ErrUnexpectedSigningMethod      = "unexpected signing method"
)

type JWTManager struct {
	secretKey       string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewJWTManager(
	secretKey string,
	accessTokenTTL time.Duration,
	refreshTokenTTL time.Duration,
) *JWTManager {
	return &JWTManager{
		secretKey:       secretKey,
		accessTokenTTL:  accessTokenTTL,
		refreshTokenTTL: refreshTokenTTL,
	}
}

func (m *JWTManager) GenerateAccessToken(user *models.User) (string, error) {
	now := time.Now()

	claims := jwt.RegisteredClaims{
		Subject:   user.ID,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(m.accessTokenTTL)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(m.secretKey))
	if err != nil {
		return "", errx.NewInternal().WithDescriptionAndCause(ErrFailedToGenerateAccessToken, err)
	}

	return signed, nil
}

func (m *JWTManager) GenerateRefreshToken(user *models.User) (string, error) {
	now := time.Now()

	claims := jwt.RegisteredClaims{
		Subject:   user.ID,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(m.refreshTokenTTL)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(m.secretKey))
	if err != nil {
		return "", errx.NewInternal().WithDescriptionAndCause(ErrFailedToGenerateRefreshToken, err)
	}

	return signed, nil
}

func (m *JWTManager) GenerateTokens(user *models.User) (string, string, error) {
	access, err := m.GenerateAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refresh, err := m.GenerateRefreshToken(user)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (m *JWTManager) VerifyToken(tokenStr string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&jwt.RegisteredClaims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errx.NewInternal().WithDescription(ErrUnexpectedSigningMethod)
			}

			return []byte(m.secretKey), nil
		},
	)

	if err != nil {
		return nil, errx.NewInternal().WithDescriptionAndCause(ErrFailedToVerifyToken, err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, errx.NewInternal().WithDescription(ErrFailedToVerifyToken)
	}

	return claims, nil
}
