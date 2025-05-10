package models

import (
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/nordew/go-errx"
)

const (
	_SolanaWalletRegex   = `^0x[a-fA-F0-9]{40}$`
	_EthereumWalletRegex = `^0x[a-fA-F0-9]{40}$`
)

var (
	solanaWalletRe   = regexp.MustCompile(_SolanaWalletRegex)
	ethereumWalletRe = regexp.MustCompile(_EthereumWalletRegex)
)

var (
	ErrInvalidUsername      = errx.NewBadRequest().WithDescription("invalid username")
	ErrInvalidWalletAddress = errx.NewBadRequest().WithDescription("invalid wallet address")
)

type User struct {
	ID            string    `json:"id" db:"id"`
	Username      string    `json:"username" db:"username"`
	WalletAddress string    `json:"walletAddress" db:"wallet_address"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt     time.Time `json:"updatedAt" db:"updated_at"`
}

func NewUser(
	id,
	username,
	walletAddress string,
) (*User, error) {
	now := time.Now()

	user := &User{
		ID:            id,
		Username:      username,
		WalletAddress: walletAddress,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if _, err := uuid.Parse(u.ID); err != nil {
		return ErrInvalidID
	}
	if u.Username == "" {
		return ErrInvalidUsername
	}
	if !isValidWalletAddress(u.WalletAddress) {
		return ErrInvalidWalletAddress
	}

	return nil
}

func isValidWalletAddress(address string) bool {
	return solanaWalletRe.MatchString(address) ||
		ethereumWalletRe.MatchString(address)
}
