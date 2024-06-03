package controllers

import (
	"fmt"
	"section3/wallet/usecase/card"

	"github.com/gin-gonic/gin"
)

func CreateVirtualAccount(ginContext *gin.Context) {

	todoService := card.NewCreateVirtualCardService()
	fmt.Println("me")
	todoService.CreateVirtualAccount(ginContext)
}
