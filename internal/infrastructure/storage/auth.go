package storage

import (
	"context"

	"github.com/nordew/go-errx"
	"github.com/nordew/warpbid/internal/models"
)

const (
	ErrFailedToSaveNonce = "failed to save nonce"
	ErrFailedToGetNonce  = "failed to get nonce"
)

func (s *Storage) SaveNonce(ctx context.Context, authChallenge models.AuthChallenge) error {
	err := s.dragonfly.Set(ctx, authChallenge.WalletAddress, authChallenge.Nonce, authChallenge.ExpiresAt).Err()
	if err != nil {
		return errx.NewInternal().WithDescriptionAndCause(ErrFailedToSaveNonce, err)
	}

	return nil
}

func (s *Storage) GetNonce(ctx context.Context, walletAddress string) (string, error) {
	nonce, err := s.dragonfly.Get(ctx, walletAddress).Result()
	if err != nil {
		return "", errx.NewInternal().WithDescriptionAndCause(ErrFailedToGetNonce, err)
	}

	return nonce, nil
}
