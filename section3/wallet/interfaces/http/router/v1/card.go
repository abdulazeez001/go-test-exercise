package v1

import (
	"section3/wallet/interfaces/http/controllers"

	"github.com/gin-gonic/gin"
)

// todo v1 routes
func CardV1Routes(v1 *gin.RouterGroup) {
	cardV1 := v1.Group("/create/dva")

	{
		cardV1.GET("/", controllers.CreateVirtualAccount)
	}
}
