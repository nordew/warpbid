package storage

import (
	"context"
	"time"

	"github.com/nordew/go-errx"
)

const (
	ErrFailedToSaveNonce = "failed to save nonce"
	ErrFailedToGetNonce  = "failed to get nonce"
)

func (s *Storage) SaveNonce(ctx context.Context, walletAddress, nonce string, expire time.Duration) error {
	err := s.dragonfly.Set(ctx, walletAddress, nonce, expire).Err()
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
