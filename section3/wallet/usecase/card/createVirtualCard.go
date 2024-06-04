package card

import (
	"net/http"
	"section3/wallet/interfaces/http/manager"

	"github.com/gin-gonic/gin"
)

type VirtualCardService struct{}

func NewCreateVirtualCardService() *VirtualCardService {
	return &VirtualCardService{}
}

func (virtualCardService *VirtualCardService) CreateVirtualAccount(ginContext *gin.Context) {
	manager.GetSuccessResponse(
		map[string]interface{}{"message": "pong"},
		"Fetched response successfully",
		http.StatusOK,
		ginContext,
	)
}
