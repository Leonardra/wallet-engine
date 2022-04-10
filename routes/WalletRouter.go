package routes

import (
	"github.com/gin-gonic/gin"
	"walletEngine/handlers"
)

type Router struct{
	walletHandler  handlers.Handler
	transactionHandler handlers.TransactionHandler
}

func CreateRouter (walletHandler  handlers.Handler, transactionHandler handlers.TransactionHandler) Router{
	return Router{walletHandler, transactionHandler}
}

func (appRouter *Router) WalletRouter (router *gin.Engine){
	walletRoutes := router.Group("api/v1/wallet")
	{
		walletRoutes.POST("/", appRouter.walletHandler.CreateWallet())
		walletRoutes.GET("/:walletId", appRouter.walletHandler.GetWallet())
		walletRoutes.PATCH("/:walletId/debit", appRouter.transactionHandler.DebitWallet())
		walletRoutes.PATCH("/:walletId/credit", appRouter.transactionHandler.CreditWallet())
		walletRoutes.PATCH("/:walletId/active/", appRouter.walletHandler.ChangeActivationStatus())
	}
}
