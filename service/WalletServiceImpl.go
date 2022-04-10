package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"walletEngine/data/models"
	"walletEngine/data/repository"
	"walletEngine/dto"
	"walletEngine/util"
)

type WalletServiceImpl struct {
	repository repository.Repository
}

func CreateWalletService(repository repository.Repository) WalletServiceImpl{
	return WalletServiceImpl{repository}
}

func (walletService WalletServiceImpl)CreateWallet(walletRequest dto.WalletRequest) (*models.Wallet, error){
	wallet := models.Wallet{
		Id : primitive.NewObjectID(),
		FirstName: walletRequest.FirstName,
		LastName: walletRequest.LastName,
		DateCreated : time.Now().Local(),
		AccountNumber :util.GenerateAccountNumber(),
		Balance : 0.0,
		ActivationStatus : true,
	}
	util.ApplicationLog.Printf("wallet before saving %v\n", wallet)
	err := walletService.repository.CreateWallet(wallet)
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (walletService WalletServiceImpl) GetWallet(id primitive.ObjectID) (*models.Wallet, error){
	foundWallet, err := walletService.repository.GetWallet(id)
	if err != nil {
		return nil, err
	}
	return foundWallet, nil
}

func (walletService WalletServiceImpl) UpdateWallet(id primitive.ObjectID, wallet *models.Wallet)(*models.Wallet, error){
	err := walletService.repository.UpdateWallet(id, wallet)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

