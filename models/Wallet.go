package models

import (
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"walletEngine/util"
)

type Wallet struct{
	Id					primitive.ObjectID 	`json:"id" bson:"id"`
	FirstName          string             	`json:"firstName" validate:"required" bson:"firstName"`
	LastName          string             	`json:"lastName" validate:"required" bson:"lastName"`
	DateCreated   	time.Time          `json:"dateCreated" bson:"dateCreated"`
	Balance			decimal.Decimal		`json:"balance" bson:"balance" `
	AccountNumber	string				`json:"accountNumber" bson:"accountNumber" `
	ActivationStatus  bool				`json:"activationStatus" bson:"activationStatus"`
}


func  CreateWalletInstance(firstName string, lastName string) *Wallet{
	 wallet := new(Wallet)
	 wallet.Id = primitive.NewObjectID()
	 wallet.FirstName = firstName
	 wallet.LastName = lastName
	 wallet.AccountNumber = util.GenerateAccountNumber()
	 wallet.Balance = decimal.RequireFromString("0.0")
	 wallet.ActivationStatus = true
	 return wallet
}