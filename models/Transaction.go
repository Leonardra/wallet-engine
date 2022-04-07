package models

import (
	"github.com/shopspring/decimal"
)

type Transaction struct{
	Amount				decimal.Decimal		`json:"amount" bson:"amount" `
	AccountNumber   	string					`json:"accountNumber" bson:"accountNumber"`
}


func CreateTransactionInstance(amount string, accountNumber string) *Transaction{
	transaction := new(Transaction)
	transaction.Amount = decimal.RequireFromString(amount)
	transaction.AccountNumber = accountNumber
	return transaction
}
