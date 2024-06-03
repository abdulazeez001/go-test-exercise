// internal/services/stellar.go
package pkg

import (
	"github.com/stellar/go/keypair"
)

type StellarService struct{}

func NewStellarService() *StellarService {
	return &StellarService{}
}

// name
func (s *StellarService) CreateWallet() (string, string, error) {
	pair, err := keypair.Random()
	if err != nil {
		return "", "", err
	}

	address := pair.Address()
	secret := pair.Seed()

	return address, secret, nil
}
