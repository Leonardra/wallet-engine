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
type Handler struct {
	walletService 		service.WalletService
	validate			validator.Validate
}



func CreateWalletHandler(walletService service.WalletService) Handler {
	return Handler{walletService, *validator.New()}
}


func (handler *Handler) CreateWallet() gin.HandlerFunc{
	return func(c *gin.Context){

		var walletRequest dto.WalletRequest
		if err := c.ShouldBindJSON(&walletRequest); err != nil {
			util.ApplicationLog.Printf("Error binding Json Obj %v\n", err)
			util.GenerateJSONResponse(c, http.StatusBadRequest, err.Error(), gin.H{})
			return
		}
		if validationErr := handler.validate.Struct(&walletRequest); validationErr != nil {
			util.ApplicationLog.Println("validation error")
			util.GenerateBadRequestResponse(c, validationErr.Error())
			return
		}

		wallet, err := handler.walletService.CreateWallet(walletRequest)
		if err != nil {
			util.GenerateInternalServerErrorResponse(c, err.Error())
			return
		}

		util.GenerateJSONResponse(c, http.StatusCreated, "Success", gin.H{
			"wallet": wallet,
		})
	}
}


func (handler *Handler) GetWallet() gin.HandlerFunc{
	return func(c *gin.Context){
		walletId := c.Param("walletId")
		util.ApplicationLog.Printf("walletId received %v\n",walletId)
		objectId, _ := primitive.ObjectIDFromHex(walletId)

		wallet, err := handler.walletService.GetWallet(objectId)
		if err != nil {
			util.GenerateJSONResponse(c, http.StatusNotFound, "Not Found", gin.H{})
			return
		}

		util.GenerateJSONResponse(c, http.StatusFound, "Found", gin.H{
			"wallet": wallet,
		})
	}
}

func (handler *Handler) ChangeActivationStatus()gin.HandlerFunc{
	return func(c *gin.Context){

		walletId := c.Param("walletId")
		util.ApplicationLog.Printf("walletId received %v\n",walletId)
		objectId, _ := primitive.ObjectIDFromHex(walletId)


		var activationRequest dto.ActivationRequest
		if err := c.ShouldBindJSON(&activationRequest); err != nil {
			util.ApplicationLog.Printf("Error binding Json Obj %v\n", err)
			util.GenerateJSONResponse(c, http.StatusBadRequest, err.Error(), gin.H{})
			return
		}

		walletToUpdate, err :=  handler.walletService.GetWallet(objectId)
		if err != nil {
			util.GenerateJSONResponse(c, http.StatusNotFound, "Not Found", gin.H{})
			return
		}

		walletToUpdate.ActivationStatus = activationRequest.Active

		updateWallet, err := handler.walletService.UpdateWallet(walletToUpdate.Id, walletToUpdate)
		if err != nil {
			util.GenerateBadRequestResponse(c, err.Error())
			return
		}


		util.GenerateJSONResponse(c, http.StatusFound, "Found", gin.H{
			"wallet": updateWallet,
		})

	}
}

