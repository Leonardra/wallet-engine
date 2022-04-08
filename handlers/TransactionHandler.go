package handlers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

		var wallet  models.Wallet
		transaction := new(models.Transaction)
		if err := c.ShouldBindJSON(&transaction); err != nil {
			util.ApplicationLog.Printf("Error binding Json Obj %v\n", err)
			util.GenerateJSONResponse(c, http.StatusBadRequest, err.Error(), gin.H{})
			return
		}

		filter := bson.D{{"accountNumber", transaction.AccountNumber}}

		err := walletCollection.FindOne(ctx, filter).Decode(&wallet)
		if err == mongo.ErrNoDocuments {
			util.GenerateJSONResponse(c, http.StatusNotFound, "Not Found", gin.H{})
			return
		} else if err != nil {
			util.GenerateInternalServerErrorResponse(c, err.Error())
		}

		debitedWallet, err := debitFromWallet(wallet, transaction)
		if err != nil {
			util.GenerateBadRequestResponse(c, err.Error())
			return
		}



		idFilter := bson.D{{"id", debitedWallet.Id}}
		util.ApplicationLog.Printf("Wallet after debit %v\n",debitedWallet)
		updateResult := walletCollection.FindOneAndReplace(ctx, idFilter,debitedWallet)
		err = updateResult.Err()
		if err != nil {
			util.ApplicationLog.Printf("Error updating wallet: %v\n", err)
			util.GenerateInternalServerErrorResponse(c, err.Error())
			return
		}

		util.GenerateJSONResponse(c, http.StatusOK, "Success", gin.H{
			"wallet": updateResult,
		})

	}
}

func CreditWallet()gin.HandlerFunc{
	return func(c *gin.Context){
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var wallet  models.Wallet
		transaction := new(models.Transaction)
		if err := c.ShouldBindJSON(&transaction); err != nil {
			util.ApplicationLog.Printf("Error binding Json Obj %v\n", err)
			util.GenerateJSONResponse(c, http.StatusBadRequest, err.Error(), gin.H{})
			return
		}

		filter := bson.D{{"accountNumber", transaction.AccountNumber}}
		err := walletCollection.FindOne(ctx, filter).Decode(&wallet)
		if err == mongo.ErrNoDocuments {
			util.GenerateJSONResponse(c, http.StatusNotFound, "Not Found", gin.H{})
			return
		} else if err != nil {
			util.GenerateInternalServerErrorResponse(c, err.Error())
		}

		creditedWallet, err := creditWallet(wallet, transaction)
		if err != nil {
			util.GenerateBadRequestResponse(c, err.Error())
			return
		}


		idFilter := bson.D{{"id", creditedWallet.Id}}
		util.ApplicationLog.Printf("Wallet after credit %v\n",creditedWallet)
		updateResult := walletCollection.FindOneAndReplace(ctx, idFilter, creditedWallet)
		err = updateResult.Err()
		if err != nil {
			util.ApplicationLog.Printf("Error updating wallet: %v\n", err)
			util.GenerateInternalServerErrorResponse(c, err.Error())
			return
		}

		util.GenerateJSONResponse(c, http.StatusOK, "Success", gin.H{
			"wallet": updateResult,
		})

	}
}

func creditWallet(wallet models.Wallet, transaction *models.Transaction)(*models.Wallet, error){
	if transaction.Amount.IsNegative(){
		return &wallet, errors.New("credit amount cannot be negative number")
	}
	if wallet.ActivationStatus == true {
		wallet.Balance = wallet.Balance.Add(transaction.Amount)
	}
	return &wallet, nil
}


func debitFromWallet(wallet models.Wallet, transaction *models.Transaction) (*models.Wallet, error){
	if transaction.Amount.GreaterThan(wallet.Balance){
		return &wallet, errors.New("debit amount cannot exceed balance")
	}
	if wallet.ActivationStatus == true {
		wallet.Balance = wallet.Balance.Sub(transaction.Amount)
	}
	return &wallet, nil
}
