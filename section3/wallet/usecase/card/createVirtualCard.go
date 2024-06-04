package card

import (
	"net/http"
	"section3/wallet/infra/models"
	"section3/wallet/interfaces/http/manager"
	pkg "section3/wallet/pkg/flutterwave"

	"github.com/gin-gonic/gin"
)

type VirtualCardService struct{}

func NewCreateVirtualCardService() *VirtualCardService {
	return &VirtualCardService{}
}

func (virtualCardService *VirtualCardService) CreateVirtualAccount(ginContext *gin.Context) {
	var request models.VirtualAccountRequest

	if err := ginContext.ShouldBindJSON(&request); err != nil {
		http.Error(ginContext.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := pkg.CreateVirtualAccount(request)
	if err != nil {
		http.Error(ginContext.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	response = models.VirtualAccountResponse{AccountNumber: response.AccountNumber, BankName: response.BankName, AccountName: response.AccountName}
	responseMap := map[string]interface{}{
		"accountNumber": response.AccountName,
		"bankname":      response.BankName,
		"accountName":   response.AccountName,
	}

	manager.GetSuccessResponse(responseMap, "Wallet created successfully", http.StatusOK, ginContext)
}
