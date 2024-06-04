// internal/services/stellar.go
package pkg

import (
	"io/ioutil"
	"net/http"

	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/protocols/horizon"
)

type StellarService struct{}

func NewStellarService() *StellarService {
	return &StellarService{}
}

// name
func (s *StellarService) CreateWallet() ([]horizon.Balance, string, string, error) {
	pair, err := keypair.Random()
	if err != nil {
		return []horizon.Balance{}, "", "", err
	}

	address := pair.Address()
	secret := pair.Seed()

	resp, err := http.Get("https://friendbot.stellar.org/?addr=" + address)
	if err != nil {
		return []horizon.Balance{}, "", "", err
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return []horizon.Balance{}, "", "", err
	}
	request := horizonclient.AccountRequest{AccountID: address}
	account, err := horizonclient.DefaultTestNetClient.AccountDetail(request)
	if err != nil {
		return []horizon.Balance{}, "", "", err
	}

	return account.Balances, address, secret, nil
}
