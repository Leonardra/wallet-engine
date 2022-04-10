package service

import (
	"walletEngine/data/models"
)

type TransactionService interface {
	DebitWallet(wallet *models.Wallet, amount float64) (*models.Wallet, error)
	CreditWallet(wallet *models.Wallet, amount float64) (*models.Wallet, error)
}
