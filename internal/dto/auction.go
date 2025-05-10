package dto

import "time"

type GetAuctionFilter struct {
	ID           string    `json:"id,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	StartPrice   int       `json:"startPrice,omitempty"`
	CurrentPrice int       `json:"currentPrice,omitempty"`
	Status       string    `json:"status,omitempty"`
}

type CreateAuctionDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	StartPrice  float64   `json:"startPrice"`
}

type UpdateAuctionDTO struct {
	ID          string    `json:"id"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	EndTime     time.Time `json:"endTime"`
}
