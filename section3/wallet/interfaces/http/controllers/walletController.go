package controllers

import (
	"fmt"
	wallet "section3/wallet/usecase/wallet"

	"github.com/gin-gonic/gin"
)

func CreateWallet(ginContext *gin.Context) {

	walletService := wallet.NewWalletHandler()
	walletService.CreateWallet(ginContext)
}
