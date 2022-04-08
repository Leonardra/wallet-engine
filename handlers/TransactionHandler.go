package handlers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
	"walletEngine/models"
	"walletEngine/util"
)


func DebitWallet() gin.HandlerFunc{
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		walletId := c.Param("walletId")
		util.ApplicationLog.Printf("walletId received %v\n", walletId)
		objectId, _ := primitive.ObjectIDFromHex(walletId)

		transaction := new(models.Transaction)
		if err := c.ShouldBindJSON(&transaction); err != nil {
			util.ApplicationLog.Printf("Error binding Json Obj %v\n", err)
			util.GenerateJSONResponse(c, http.StatusBadRequest, err.Error(), gin.H{})
			return
		}


		filter := bson.D{{"_id", objectId}}
		foundResult, done := findWallet(c, ctx, filter)
		if done {
			return
		}

		wallet, err := debitFromWallet(foundResult, transaction)
		if err != nil {
			util.GenerateBadRequestResponse(c, err.Error())
			return
		}

		singleResult := walletCollection.FindOneAndReplace(ctx, filter, wallet)
		err = singleResult.Err()
		if err == mongo.ErrNoDocuments {
			util.GenerateBadRequestResponse(c, err.Error())
		} else if err != nil {
			util.GenerateInternalServerErrorResponse(c, err.Error())
		}

		foundWallet, done := findWallet(c, ctx, filter)
		if done {
			return
		}
		util.GenerateJSONResponse(c, http.StatusOK, "Success", gin.H{
			"wallet": foundWallet,
		})
	}
}

func findWallet(c *gin.Context, ctx context.Context, filter bson.D) (models.Wallet, bool) {
	var foundResult models.Wallet
	err := walletCollection.FindOne(ctx, filter).Decode(&foundResult)
	if err == mongo.ErrNoDocuments {
		util.GenerateJSONResponse(c, http.StatusNotFound, "Not Found", gin.H{})
		return models.Wallet{}, true
	} else if err != nil {
		util.GenerateInternalServerErrorResponse(c, err.Error())
		return models.Wallet{}, true
	}
	return foundResult, false
}

func CreditWallet()gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		walletId := c.Param("walletId")
		util.ApplicationLog.Printf("walletId received %v\n", walletId)
		objectId, _ := primitive.ObjectIDFromHex(walletId)

		transaction := new(models.Transaction)
		if err := c.ShouldBindJSON(&transaction); err != nil {
			util.ApplicationLog.Printf("Error binding Json Obj %v\n", err)
			util.GenerateJSONResponse(c, http.StatusBadRequest, err.Error(), gin.H{})
			return
		}


		filter := bson.D{{"_id", objectId}}
		walletToUpdate, done := findWallet(c, ctx, filter)
		if done {
			return
		}

		wallet, err := creditWallet(walletToUpdate, transaction)
		if err != nil {
			util.GenerateBadRequestResponse(c, err.Error())
			return
		}

		singleResult := walletCollection.FindOneAndReplace(ctx, filter, wallet)
		err = singleResult.Err()
		if err == mongo.ErrNoDocuments {
			util.GenerateBadRequestResponse(c, err.Error())
		} else if err != nil {
			util.GenerateInternalServerErrorResponse(c, err.Error())
		}


		foundWallet, done := findWallet(c, ctx, filter)
		if done {
			return
		}
		util.GenerateJSONResponse(c, http.StatusOK, "Success", gin.H{
			"wallet": foundWallet,
		})

	}
}

func creditWallet(wallet models.Wallet, transaction *models.Transaction)(*models.Wallet, error){
	if transaction.Amount <= 0.0{
		return &wallet, errors.New("credit amount cannot be negative number")
	}
	if wallet.ActivationStatus == true {
		wallet.Balance = wallet.Balance + transaction.Amount
		util.ApplicationLog.Printf("New Balance%v\n", wallet.Balance)
	}else{
		return &wallet, errors.New("wallet must be activated")
	}
	return &wallet, nil
}


func debitFromWallet(wallet models.Wallet, transaction *models.Transaction) (*models.Wallet, error){
	if transaction.Amount > wallet.Balance{
		return &wallet, errors.New("debit amount cannot exceed balance")
	}
	if wallet.ActivationStatus == true {
		wallet.Balance = wallet.Balance - transaction.Amount
	}else{
		return &wallet, errors.New("wallet must be activated")
	}
	return &wallet, nil
}
