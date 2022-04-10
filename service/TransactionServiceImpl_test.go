package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"walletEngine/data/models"
	"walletEngine/dto"
	"walletEngine/util"
)

var transactionService = CreateTransactionService(walletRepository)

func Test_CreditWallet(t *testing.T){
	var wallet  models.Wallet
	wallet.ActivationStatus = true
	transaction := dto.Transaction{
		Amount: 200000.00,
	}
	creditedWallet, _ := transactionService.CreditWallet(&wallet, transaction.Amount)
	util.ApplicationLog.Printf("Credited Wallet %v\n", creditedWallet.Balance)

	assert.Equal(t, creditedWallet.Balance, 200000.00)
}


func Test_CreditWalletIsNegative(t *testing.T){
	var wallet  models.Wallet
	wallet.ActivationStatus = true
	transaction := dto.Transaction{
		Amount: -200000.00,
	}

	creditedWallet, err := transactionService.CreditWallet(&wallet, transaction.Amount)
	util.ApplicationLog.Printf("Credited Wallet %v\n", creditedWallet.Balance)

	assert.Equal(t, creditedWallet.Balance,  0.0)
	assert.NotNil(t, err)
}

func Test_DebitWallet(t *testing.T){
	var wallet  models.Wallet
	wallet.ActivationStatus = true
	wallet.Balance = 200000.00
	transaction := dto.Transaction{
		Amount: 100000.00,
	}
	debitedWallet, _ := transactionService.DebitWallet(&wallet, transaction.Amount)
	util.ApplicationLog.Printf("Credited Wallet %v\n", debitedWallet.Balance)

	assert.Equal(t, (*debitedWallet).Balance,100000.00)
}

func Test_debitAmountDoesNotExceedBalance(t *testing.T){
	var wallet  models.Wallet
	wallet.Balance = 100000.00
	transaction := dto.Transaction{
		Amount: 200000.00,
	}

	debitedWallet, err := transactionService.DebitWallet(&wallet, transaction.Amount)
	util.ApplicationLog.Printf("Debited Wallet %v\n", (*debitedWallet).Balance)

	assert.Error(t, err, "debit amount cannot exceed balance")
	assert.Equal(t, debitedWallet.Balance,100000.00)
}
