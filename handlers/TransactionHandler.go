package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"walletEngine/dto"
	"walletEngine/service"
	"walletEngine/util"
)


type TransactionHandler struct {
	transactionService 		service.TransactionService
	walletService 		service.WalletService
	validate			validator.Validate
}



func CreateTransactionHandler(transactionService service.TransactionService, walletService 	service.WalletService) TransactionHandler {
	return TransactionHandler{transactionService,walletService, *validator.New()}
}

func (transactionHandler *TransactionHandler) DebitWallet() gin.HandlerFunc {
	return func(c *gin.Context){
		walletId := c.Param("walletId")
		util.ApplicationLog.Printf("walletId received %v\n", walletId)
		objectId, _ := primitive.ObjectIDFromHex(walletId)

		transaction := new(dto.Transaction)
		if err := c.ShouldBindJSON(&transaction); err != nil {
			util.ApplicationLog.Printf("Error binding Json Obj %v\n", err)
			util.GenerateJSONResponse(c, http.StatusBadRequest, err.Error(), gin.H{})
			return
		}
		wallet, err := transactionHandler.walletService.GetWallet(objectId)
		if err != nil {
			util.GenerateJSONResponse(c, http.StatusNotFound, "Not Found", gin.H{})
			return
		}



		debitWallet, err := transactionHandler.transactionService.DebitWallet(wallet, transaction.Amount)
		if err != nil {
			util.GenerateBadRequestResponse(c, err.Error())
			return
		}

		updateWallet, err := transactionHandler.walletService.UpdateWallet(debitWallet.Id, debitWallet)
		if err != nil {
			util.GenerateBadRequestResponse(c, err.Error())
			return
		}


		util.GenerateJSONResponse(c, http.StatusFound, "Found", gin.H{
			"wallet": updateWallet,
		})


	}
}

func (transactionHandler *TransactionHandler) CreditWallet() gin.HandlerFunc {
	return func(c *gin.Context){
		walletId := c.Param("walletId")
		util.ApplicationLog.Printf("walletId received %v\n", walletId)
		objectId, _ := primitive.ObjectIDFromHex(walletId)

		transaction := new(dto.Transaction)
		if err := c.ShouldBindJSON(&transaction); err != nil {
			util.ApplicationLog.Printf("Error binding Json Obj %v\n", err)
			util.GenerateJSONResponse(c, http.StatusBadRequest, err.Error(), gin.H{})
			return
		}
		wallet, err := transactionHandler.walletService.GetWallet(objectId)
		if err != nil {
			util.GenerateJSONResponse(c, http.StatusNotFound, "Not Found", gin.H{})
			return
		}


		creditWallet, err := transactionHandler.transactionService.CreditWallet(wallet, transaction.Amount)
		if err != nil {
			util.GenerateBadRequestResponse(c, err.Error())
			return
		}

		updateWallet, err := transactionHandler.walletService.UpdateWallet((*creditWallet).Id, creditWallet)
		if err != nil {
			util.GenerateBadRequestResponse(c, err.Error())
			return
		}


		util.GenerateJSONResponse(c, http.StatusFound, "Found", gin.H{
			"wallet": updateWallet,
		})


	}
}