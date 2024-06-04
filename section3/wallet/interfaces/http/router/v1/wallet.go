package v1

import (
	"section3/wallet/interfaces/http/controllers"

	"github.com/gin-gonic/gin"
)

// todo v1 routes
func WalletV1Routes(v1 *gin.RouterGroup) {
	walletV1 := v1.Group("/wallet")

	{
		walletV1.GET("", controllers.CreateWallet)
	}
}
