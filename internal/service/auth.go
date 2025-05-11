package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"github.com/nordew/go-errx"
	"github.com/nordew/warpbid/internal/dto"
)

const (
	ErrInvalidWalletAddress = "invalid wallet address"
	ErrInvalidSignature     = "invalid signature"
)

func (s *Service) GetNonce(ctx context.Context, walletAddress string) (string, error) {
	if !common.IsHexAddress(walletAddress) {
		return "", errx.NewBadRequest().WithDescription(ErrInvalidWalletAddress)
	}

	nonce := uuid.NewString()
	expiration := 15 * time.Minute

	if err := s.storage.SaveNonce(ctx, walletAddress, nonce, expiration); err != nil {
		return "", err
	}

	return nonce, nil
}

func (s *Service) Verify(ctx context.Context, req dto.VerifyAuthRequest) (string, string, error) {
	nonce, err := s.storage.GetNonce(ctx, req.WalletAddress)
	if err != nil {
		return "", "", err
	}

	message := "Sign this message: " + nonce

	recoveredAddr, err := verifyEthSignature(message, req.Signature)
	if err != nil || !strings.EqualFold(recoveredAddr, req.WalletAddress) {
		return "", "", errx.NewUnauthorized().WithDescription(ErrInvalidSignature)
	}

	user, err := s.storage.GetUserByFilter(ctx, &dto.GetUserFilter{
		WalletAddress: req.WalletAddress,
	})
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := s.jwtManager.GenerateTokens(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func verifyEthSignature(message, signatureHex string) (string, error) {
	signatureHex = strings.TrimPrefix(signatureHex, "0x")

	signature := common.FromHex(signatureHex)

	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	messageHash := crypto.Keccak256Hash([]byte(prefixedMessage))

	if len(signature) == 65 && (signature[64] == 27 || signature[64] == 28) {
		signature[64] -= 27
	}

	pubKey, err := crypto.SigToPub(messageHash.Bytes(), signature)
	if err != nil {
		return "", fmt.Errorf("failed to recover public key: %w", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey).Hex()

	return strings.ToLower(recoveredAddr), nil
}
