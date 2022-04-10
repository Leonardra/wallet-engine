package service

import (
	"github.com/shopspring/decimal"
	"walletEngine/data/models"
)

type TransactionService interface {
	DebitWallet(wallet *models.Wallet, amount decimal.Decimal) (*models.Wallet, error)
	CreditWallet(wallet *models.Wallet, amount decimal.Decimal) (*models.Wallet, error)
}
