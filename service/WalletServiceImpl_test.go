package service

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"walletEngine/configs"
	"walletEngine/data/repository"
	"walletEngine/dto"
)

var walletRepository = repository.CreateMongoRepository()
var service = CreateWalletService(walletRepository)

func CleanUpDbOps(ctx *gin.Context) {
	err := walletRepository.Collection.Database().Drop(ctx)
	if err != nil {
		return
	}
	err = configs.DbClient.Disconnect(ctx)
	if err != nil {
		return
	}
}

func Test_CreateWalletService(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	walletRequest := dto.WalletRequest{
		FirstName: "John",
		LastName:  "Doe",
	}

	wallet, err := service.CreateWallet(walletRequest)
	if err != nil {
		return
	}
	assert.Nil(t, err)
	assert.NotNil(t, wallet)
	t.Cleanup(func() {
		CleanUpDbOps(ctx)
	})
}


func Test_GetWallet(t *testing.T){
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	walletRequest := dto.WalletRequest{
		FirstName: "John",
		LastName:  "Doe",
	}

	wallet, err := service.CreateWallet(walletRequest)
	if err != nil {
		return
	}

	foundWallet, err := service.GetWallet(wallet.Id)

	assert.Nil(t, err)
	assert.NotNil(t,foundWallet)
	t.Cleanup(func() {
		CleanUpDbOps(ctx)
	})
}
