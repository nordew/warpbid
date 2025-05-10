package dto

type GetUserFilter struct {
	ID            string `json:"id,omitempty"`
	Username      string `json:"username,omitempty"`
	WalletAddress string `json:"walletAddress,omitempty"`
}

type CreateUserDTO struct {
	Username      string `json:"username"`
	WalletAddress string `json:"walletAddress"`
}

type UpdateUserDTO struct {
	ID            string `json:"id"`
	Username      string `json:"username,omitempty"`
	WalletAddress string `json:"walletAddress,omitempty"`
}
