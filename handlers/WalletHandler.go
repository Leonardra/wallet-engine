package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
	"walletEngine/configs"
	"walletEngine/models"
	"walletEngine/util"
)
  var walletCollection = configs.GetCollection(configs.DbClient, "wallets")
  var validate = validator.New()

func CreateWallet() gin.HandlerFunc{

	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var wallet models.Wallet
		if err := c.ShouldBindJSON(&wallet); err != nil {
			util.ApplicationLog.Printf("Error binding Json Obj %v\n", err)
			util.GenerateJSONResponse(c, http.StatusBadRequest, err.Error(), gin.H{})
			return
		}
		if validationErr := validate.Struct(&wallet); validationErr != nil {
			util.ApplicationLog.Println("validation error")
			util.GenerateBadRequestResponse(c, validationErr.Error())
			return
		}

		wallet.Id = primitive.NewObjectID()
		wallet.DateCreated = time.Now()
		wallet.AccountNumber = util.GenerateAccountNumber()
		wallet.Balance = decimal.RequireFromString("0.0")
		util.ApplicationLog.Printf("wallet before saving %v\n", wallet)

		savedWallet, err := walletCollection.InsertOne(ctx, wallet)

		if err != nil {
			util.ApplicationLog.Printf("Error Saving wallet %v\n", err)
			util.GenerateInternalServerErrorResponse(c, err.Error())
			return
		}

		util.GenerateJSONResponse(c, http.StatusCreated, "Success", gin.H{
			"wallet": savedWallet,
		})
	}


}