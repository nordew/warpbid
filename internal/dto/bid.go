package dto

type GetBidFilter struct {
	ID        string `json:"id,omitempty"`
	AuctionID string `json:"auctionId,omitempty"`
	UserID    string `json:"userId,omitempty"`
	Amount    int    `json:"amount,omitempty"`
}

type CreateBidDTO struct {
	AuctionID string `json:"auctionId"`
	UserID    string `json:"userId"`
	Amount    int    `json:"amount"`
}
