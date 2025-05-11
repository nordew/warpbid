package dto

type VerifyAuthRequest struct {
	WalletAddress string `json:"wallet_address"`
	Signature     string `json:"signature"`
}
