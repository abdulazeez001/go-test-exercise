package models

import (
	"github.com/stellar/go/protocols/horizon"
)

type WalletCreationResponse struct {
	Address string            `json:"address"`
	Secret  string            `json:"secret"`
	Balance []horizon.Balance `json:"balance"`
}
