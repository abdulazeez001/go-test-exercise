package handler

import (
	"log"
	"net/http"

	pkg "section3/wallet/pkg/stellar"

	"github.com/gin-gonic/gin"
)


type Handler struct {
	StellarService     *pkg.StellarService
}

type WalletHandler interface {
	CreateWallet(ctx *gin.Context)
}

type walletHandler struct {
	walletService wallet.WalletService
}

func NewWalletHandler(ws wallet.WalletService) *walletHandler {
	return &walletHandler{
		walletService: ws,
	}
}

func (vh *walletHandler) CreateWallet(ctx *gin.Context) {
	walletResponse, err := pkg.NewStellarService.
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	ctx.JSON(http.StatusOK, walletResponse)
}
