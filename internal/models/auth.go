package models

import "time"

type AuthChallenge struct {
	WalletAddress string
	Nonce         string
	ExpiresAt     time.Time
}
