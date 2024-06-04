package controllers

import (
	"fmt"
	wallet "section3/wallet/usecase/wallet"

	"github.com/gin-gonic/gin"
)

func CreateWallet(ginContext *gin.Context) {

	walletService := wallet.NewWalletHandler()
	fmt.Println("me")
	walletService.CreateWallet(ginContext)
}
