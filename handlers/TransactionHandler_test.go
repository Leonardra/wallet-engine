package handlers

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
	"walletEngine/models"
	"walletEngine/util"
)


func Test_debitWallet(t *testing.T){
	wallet := models.CreateWalletInstance("John", "Doe")
	wallet.Balance = decimal.RequireFromString("200000.10")
	transaction := models.Transaction{
		Amount: decimal.RequireFromString("100000.00"),
		AccountNumber: wallet.AccountNumber,
	}

	debitedWallet , _ := debitFromWallet(*wallet, transaction)
	util.ApplicationLog.Printf("Debited Wallet %v\n", debitedWallet)

	assert.Equal(t, debitedWallet.Balance, decimal.RequireFromString("100000.10"))
}

func Test_debitAmountDoesNotExceedBalance(t *testing.T){
	wallet := models.CreateWalletInstance("John", "Doe")
	wallet.Balance = decimal.RequireFromString("100000.00")
	transaction := models.Transaction{
		Amount: decimal.RequireFromString("200000.00"),
		AccountNumber: wallet.AccountNumber,
	}

	debitedWallet, err  := debitFromWallet(*wallet, transaction)
	util.ApplicationLog.Printf("Debited Wallet %v\n", debitedWallet)

	assert.Error(t, err, "debit amount cannot exceed balance")
	assert.Equal(t, debitedWallet.Balance, decimal.RequireFromString("100000.00"))
}