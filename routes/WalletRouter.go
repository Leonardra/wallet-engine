package routes

import (
	"github.com/gin-gonic/gin"
	"walletEngine/handlers"
)

func WalletRouter(router *gin.Engine){
	walletRoutes := router.Group("api/v1/wallet")
	{
		walletRoutes.POST("/", handlers.CreateWallet())
		walletRoutes.PATCH("/:walletId/debit", handlers.DebitWallet())
		walletRoutes.PATCH("/:walletId/credit", handlers.CreditWallet())
		walletRoutes.PATCH("/:walletId/active/", handlers.ChangeActivationStatus())
	}
}
