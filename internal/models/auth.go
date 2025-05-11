package models

import "time"

type AuthChallenge struct {
	WalletAddress string
	Nonce         string
	ExpiresAt     time.Duration
}

func NewAuthChallenge(walletAddress string, nonce string, expiresAt time.Duration) AuthChallenge {
	return AuthChallenge{
		WalletAddress: walletAddress,
		Nonce:         nonce,
		ExpiresAt:     expiresAt,
	}
}
