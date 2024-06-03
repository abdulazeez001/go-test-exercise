package models

type WalletCreationResponse struct {
	Address string `json:"address"`
	Secret  string `json:"secret"`
}
