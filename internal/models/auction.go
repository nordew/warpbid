package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nordew/go-errx"
)

var (
	ErrInvalidID           = errx.NewBadRequest().WithDescription("invalid id")
	ErrInvalidTitle        = errx.NewBadRequest().WithDescription("invalid title")
	ErrInvalidDescription  = errx.NewBadRequest().WithDescription("invalid description")
	ErrInvalidStartTime    = errx.NewBadRequest().WithDescription("invalid start time")
	ErrInvalidEndTime      = errx.NewBadRequest().WithDescription("invalid end time")
	ErrInvalidStartPrice   = errx.NewBadRequest().WithDescription("invalid start price")
	ErrInvalidCurrentPrice = errx.NewBadRequest().WithDescription("invalid current price")
)

type Auction struct {
	ID           string        `json:"id" db:"id"`
	Title        string        `json:"title" db:"title"`
	Description  string        `json:"description" db:"description"`
	StartTime    time.Time     `json:"startTime" db:"start_time"`
	EndTime      time.Time     `json:"endTime" db:"end_time"`
	StartPrice   float64       `json:"startPrice" db:"start_price"`
	CurrentPrice float64       `json:"currentPrice" db:"current_price"`
	Status       AuctionStatus `json:"status" db:"status"`
	CreatedAt    time.Time     `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time     `json:"updatedAt" db:"updated_at"`
}

type AuctionStatus string

const (
	AuctionStatusPending   AuctionStatus = "PENDING"
	AuctionStatusRunning   AuctionStatus = "RUNNING"
	AuctionStatusCompleted AuctionStatus = "COMPLETED"
	AuctionStatusCanceled  AuctionStatus = "CANCELED"
)

func NewAuction(
	id,
	title,
	description string,
	startTime,
	endTime time.Time,
	startPrice,
	currentPrice float64,
	status AuctionStatus,
) (*Auction, error) {
	now := time.Now()

	auction := &Auction{
		ID:           id,
		Title:        title,
		Description:  description,
		StartTime:    startTime,
		EndTime:      endTime,
		StartPrice:   startPrice,
		CurrentPrice: currentPrice,
		Status:       status,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := auction.Validate(); err != nil {
		return nil, err
	}

	return auction, nil
}

func (a *Auction) Validate() error {
	if _, err := uuid.Parse(a.ID); err != nil {
		return ErrInvalidID
	}
	if a.Title == "" {
		return ErrInvalidTitle
	}
	if a.Description == "" {
		return ErrInvalidDescription
	}
	if a.StartTime.IsZero() {
		return ErrInvalidStartTime
	}
	if a.EndTime.IsZero() {
		return ErrInvalidEndTime
	}
	if a.StartPrice <= 0 {
		return ErrInvalidStartPrice
	}
	if a.CurrentPrice <= 0 {
		return ErrInvalidCurrentPrice
	}

	return nil
}
