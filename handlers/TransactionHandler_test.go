package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"walletEngine/data/models"
	"walletEngine/dto"
	"walletEngine/util"
)


func Test_debitWallet(t *testing.T){
	wallet := models.CreateWalletInstance("John", "Doe")
	wallet.Balance = 200000.10
	transaction := dto.Transaction{
		Amount: 100000.00,

	}

	debitedWallet , _ := debitFromWallet(*wallet, &transaction)
	util.ApplicationLog.Printf("Debited Wallet %v\n", debitedWallet)

	assert.Equal(t, debitedWallet.Balance, 100000.10)
}

func Test_walletWillNotBeDebitedIfDeactivated(t *testing.T){
	wallet := models.CreateWalletInstance("John", "Doe")
	wallet.ActivationStatus = false
	wallet.Balance = 200000.10
	transaction := dto.Transaction{
		Amount: 100000.00,
	}

	debitedWallet , _ := debitFromWallet(*wallet, &transaction)
	util.ApplicationLog.Printf("Debited Wallet %v\n", debitedWallet)

	assert.Equal(t, debitedWallet.Balance, 200000.10)
}

func Test_debitAmountDoesNotExceedBalance(t *testing.T){
	wallet := models.CreateWalletInstance("John", "Doe")
	wallet.Balance = 100000.00
	transaction := dto.Transaction{
		Amount: 200000.00,
	}

	debitedWallet, err  := debitFromWallet(*wallet, &transaction)
	util.ApplicationLog.Printf("Debited Wallet %v\n", debitedWallet)

	assert.Error(t, err, "debit amount cannot exceed balance")
	assert.Equal(t, debitedWallet.Balance, 100000.00)
}

func Test_creditWallet(t *testing.T){
	wallet := models.CreateWalletInstance("John", "Doe")
	transaction := dto.Transaction{
		Amount: 100000.10,
	}

	creditedWallet , _ := creditWallet(*wallet, &transaction)
	util.ApplicationLog.Printf("Credited Wallet %v\n", creditedWallet)

	assert.Equal(t, creditedWallet.Balance, 100000.10)
}

func Test_creditAmountIsNotNegative(t *testing.T){
	wallet := models.CreateWalletInstance("John", "Doe")
	transaction := dto.Transaction{
		Amount: -200000.00,
	}

	creditedWallet , err := creditWallet(*wallet, &transaction)
	util.ApplicationLog.Printf("Credited Wallet %v\n", creditedWallet)

	assert.Error(t, err, "credit amount cannot be negative number")
	assert.Equal(t, creditedWallet.Balance, 0.0)
}

func Test_walletWillNotBeCreditedIfDeactivated(t *testing.T){
	wallet := models.CreateWalletInstance("John", "Doe")
	wallet.ActivationStatus = false
	transaction := dto.Transaction{
		Amount: 100000.00,
	}

	creditedWallet , _ := creditWallet(*wallet, &transaction)
	util.ApplicationLog.Printf("Credited Wallet %v\n", creditedWallet)

	assert.Equal(t, creditedWallet.Balance, 0.0)
}