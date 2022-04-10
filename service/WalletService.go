package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"walletEngine/data/models"
	"walletEngine/dto"
)

type WalletService interface {
	GetWallet(id primitive.ObjectID) (*models.Wallet, error)
	CreateWallet(walletRequest dto.WalletRequest) (*models.Wallet, error)
	UpdateWallet(id primitive.ObjectID, wallet *models.Wallet)(*models.Wallet, error)
}
