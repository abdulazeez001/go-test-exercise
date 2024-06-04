package handler

import (
	"net/http"

	"section3/wallet/infra/models"
	pkg "section3/wallet/pkg/stellar"

	"section3/wallet/interfaces/http/manager"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	StellarService *pkg.StellarService
}

func NewWalletHandler() *WalletHandler {
	return &WalletHandler{
		StellarService: pkg.NewStellarService(),
	}
}

func (h *WalletHandler) CreateWallet(ginContext *gin.Context) {
	balance, address, secret, err := h.StellarService.CreateWallet()
	if err != nil {
		http.Error(ginContext.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.WalletCreationResponse{Address: address, Secret: secret, Balance: balance}

	responseMap := map[string]interface{}{
		"address": response.Address,
		"secret":  response.Secret,
		"balance": response.Balance,
	}

	manager.GetSuccessResponse(
		responseMap,
		"Wallet created successfully",
		http.StatusOK,
		ginContext,
	)
}
