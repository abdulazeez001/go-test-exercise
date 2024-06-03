package router

import (
	"net/http"
	"os"

	v1 "section3/wallet/interfaces/http/router/v1"

	"github.com/gin-gonic/gin"
)

// Entry function for generating version1 routes
func Version1RouteEntry(router *gin.Engine) {
	version1 := router.Group("/api/v1")
	version1.GET("", func(ginContext *gin.Context) {
		ginContext.JSON(http.StatusOK, gin.H{
			"message":     "API v1  is running",
			"env":         os.Getenv("APP_ENV"),
			"serviceName": os.Getenv("SERVICE_NAME"),
		})
	})
	v1.CardV1Routes(version1)
}
