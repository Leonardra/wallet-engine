package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"walletEngine/data/models"
)

type Repository interface {
	CreateWallet(wallet models.Wallet) error
	GetWallet(id primitive.ObjectID) (*models.Wallet, error)
	UpdateWallet(id primitive.ObjectID, wallet *models.Wallet)error
}
