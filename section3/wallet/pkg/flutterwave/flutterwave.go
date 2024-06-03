// internal/services/flutterwave.go
package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"section5/wallet/infra/models"
)

type FlutterwaveService struct {
	APIKey string
}

func NewFlutterwaveService() *FlutterwaveService {
	return &FlutterwaveService{
		APIKey: os.Getenv("FLUTTERWAVE_API_KEY"),
	}
}

func (s *FlutterwaveService) HandlePayment(request models.PaymentRequest) (string, error) {
	url := "https://api.flutterwave.com/v3/payments"

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to process payment, status: %s", resp.Status)
	}

	var responseBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseBody)

	return responseBody["status"].(string), nil
}
