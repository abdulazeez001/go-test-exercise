package controllers

import (
	"fmt"
	"section3/wallet/usecase/card"

	"github.com/gin-gonic/gin"
)

func CreateVirtualAccount(ginContext *gin.Context) {

	cardService := card.NewCreateVirtualCardService()
	fmt.Println("me")
	cardService.CreateVirtualAccount(ginContext)
}
