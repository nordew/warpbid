package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nordew/go-errx"
)

var (
	ErrInvalidAuctionID = errx.NewBadRequest().WithDescription("invalid auction id")
	ErrInvalidUserID    = errx.NewBadRequest().WithDescription("invalid user id")
	ErrInvalidAmount    = errx.NewBadRequest().WithDescription("invalid amount")
)

type Bid struct {
	ID        string    `json:"id" db:"id"`
	AuctionID string    `json:"auctionId" db:"auction_id"`
	UserID    string    `json:"userId" db:"user_id"`
	Amount    int       `json:"amount" db:"amount"`
	Time      time.Time `json:"time" db:"time"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

func NewBid(
	id,
	auctionID,
	userID string,
	amount int,
) (*Bid, error) {
	now := time.Now()

	bid := &Bid{
		ID:        id,
		AuctionID: auctionID,
		UserID:    userID,
		Amount:    amount,
		Time:      now,
	}

	if err := bid.Validate(); err != nil {
		return nil, err
	}

	return bid, nil
}

func (b *Bid) Validate() error {
	if _, err := uuid.Parse(b.ID); err != nil {
		return ErrInvalidID
	}
	if b.AuctionID == "" {
		return ErrInvalidAuctionID
	}
	if b.UserID == "" {
		return ErrInvalidUserID
	}
	if b.Amount <= 0 {
		return ErrInvalidAmount
	}

	return nil
}
