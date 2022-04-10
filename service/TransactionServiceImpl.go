package service

import (
	"errors"
	"walletEngine/data/models"
	"walletEngine/data/repository"
	"walletEngine/util"
)

type TransactionServiceImpl struct {
	repository repository.Repository
}

func CreateTransactionService(repository repository.Repository)TransactionServiceImpl{
	return TransactionServiceImpl{repository}
}

func (transactionService *TransactionServiceImpl) DebitWallet(wallet *models.Wallet, amount float64) (*models.Wallet, error){
	if amount > wallet.Balance{
		return wallet, errors.New("debit amount cannot exceed balance")
	}
	if wallet.ActivationStatus == true {
		wallet.Balance = wallet.Balance - amount
	}else{
		return wallet, errors.New("wallet must be activated")
	}
	return wallet, nil


}
func  (transactionService *TransactionServiceImpl) CreditWallet(wallet *models.Wallet, amount float64) (*models.Wallet, error){
	if amount <= 0.0{
		return wallet, errors.New("credit amount cannot be negative number")
	}
	if wallet.ActivationStatus == true {
		wallet.Balance = wallet.Balance + amount
		util.ApplicationLog.Printf("New Balance%v\n", wallet.Balance)
	}else{
		return wallet, errors.New("wallet must be activated")
	}
	return wallet, nil
}

