package models

import (
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	configs "walletEngine/configs/util"
)

type Wallet struct{
	Id					primitive.ObjectID 	`json:"id" bson:"id"`
	FirstName          string             	`json:"firstName" validate:"required" bson:"firstName"`
	LastName          string             	`json:"lastName" validate:"required" bson:"lastName"`
	DateCreated   	time.Time          `json:"dateCreated" bson:"dateCreated"`
	Balance			decimal.Decimal		`json:"balance" bson:"balance" `
	AccountNumber	string				`json:"accountNumber" bson:"accountNumber" `
}


func  createWalletInstance(firstName string, lastName string) *Wallet{
	 wallet := new(Wallet)
	 wallet.FirstName = firstName
	 wallet.LastName = lastName
	 wallet.AccountNumber = configs.GenerateAccountNumber()
	 wallet.Balance = decimal.RequireFromString("0.0")
	 return wallet
}