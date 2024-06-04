package pkg

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"section3/wallet/config"
	"section3/wallet/infra/models"
)

const flutterwaveURL = "https://api.flutterwave.com/v3/virtual-account-numbers"

var FlutterwaveAPIKey = config.Get["flutterwave"].FlutterwaveApikey

func CreateVirtualAccount(request models.VirtualAccountRequest) (models.VirtualAccountResponse, error) {
	var response models.VirtualAccountResponse

	reqBody, err := json.Marshal(request)
	if err != nil {
		return response, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", flutterwaveURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return response, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+FlutterwaveAPIKey)

	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != http.StatusOK {
		return response, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	return response, nil
}
