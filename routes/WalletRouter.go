package routes

import (
	"github.com/gin-gonic/gin"
	"walletEngine/handlers"
)

func WalletRouter(router *gin.Engine){
	walletRoutes := router.Group("api/v1/wallet")
	{
		walletRoutes.POST("/", handlers.CreateWallet())
		walletRoutes.PATCH("/debit", handlers.DebitWallet())
	}
}
