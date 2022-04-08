package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
	"walletEngine/configs"
	"walletEngine/dto"
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
		wallet.DateCreated = time.Now().Local()
		wallet.AccountNumber = util.GenerateAccountNumber()
		wallet.Balance = 0.0
		wallet.ActivationStatus = true
		util.ApplicationLog.Printf("wallet before saving %v\n", wallet)

		savedWallet, err := walletCollection.InsertOne(ctx, wallet)

		if err != nil {
			util.ApplicationLog.Printf("Error Saving wallet %v\n", err)
			util.GenerateInternalServerErrorResponse(c, err.Error())
			return
		}


		var foundWallet models.Wallet
		filter := bson.D{{"_id", savedWallet.InsertedID}}
		err = walletCollection.FindOne(ctx, filter).Decode(&foundWallet)
		if err == mongo.ErrNoDocuments {
			util.GenerateJSONResponse(c, http.StatusNotFound, "Not Found", gin.H{})
			return
		} else if err != nil {
			util.GenerateInternalServerErrorResponse(c, err.Error())
		}

		util.GenerateJSONResponse(c, http.StatusCreated, "Success", gin.H{
			"wallet": foundWallet,
		})
	}
}

  func ChangeActivationStatus()gin.HandlerFunc{
	  return func(c *gin.Context){
		  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		  defer cancel()

		  walletId := c.Param("walletId")
		  util.ApplicationLog.Printf("walletId received %v\n",walletId)
		  objectId, _ := primitive.ObjectIDFromHex(walletId)


		  var activationRequest dto.ActivationRequest
		  if err := c.ShouldBindJSON(&activationRequest); err != nil {
			  util.ApplicationLog.Printf("Error binding Json Obj %v\n", err)
			  util.GenerateJSONResponse(c, http.StatusBadRequest, err.Error(), gin.H{})
			  return
		  }

		  var walletToUpdate models.Wallet

		  filter := bson.D{{"_id", objectId}}
		  err := walletCollection.FindOne(ctx, filter).Decode(&walletToUpdate)
		  if err == mongo.ErrNoDocuments {
			  util.GenerateJSONResponse(c, http.StatusNotFound, "Not Found", gin.H{})
			  return
		  } else if err != nil {
			  util.GenerateInternalServerErrorResponse(c, err.Error())
			  return
		  }

		  walletToUpdate.ActivationStatus = activationRequest.Active

		  singleResult := walletCollection.FindOneAndReplace(ctx, filter, walletToUpdate)
		  err = singleResult.Err()
		  if err == mongo.ErrNoDocuments {
			  util.GenerateBadRequestResponse(c, err.Error())
		  } else if err != nil {
			  util.GenerateInternalServerErrorResponse(c, err.Error())
		  }

		  util.GenerateJSONResponse(c, http.StatusOK, "Success", gin.H{
			  "wallet": singleResult,
		  })
	  }
  }