package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"walletEngine/configs"
	"walletEngine/dto"
	"walletEngine/util"
)

func MockJsonRequestBody(c *gin.Context, content interface{}, methodName string) {
	c.Request.Method = methodName
	c.Request.Header.Set("Content-Type", "application/json")

	requestBody, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
}

func CleanUpDbOps(ctx *gin.Context) {
	err := configs.GetCollection(configs.DbClient, "wallets").Database().Drop(ctx)
	if err != nil {
		return
	}
	err = configs.DbClient.Disconnect(ctx)
	if err != nil {
		return
	}
}


func Test_createWallet(t *testing.T){
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	requestBody := gin.H{
		"firstName": "John",
		"lastName": "Doe",
	}

	MockJsonRequestBody(ctx, requestBody, "POST")
	createWalletHandler := CreateWallet()
	createWalletHandler(ctx)

	var response dto.APIResponse
	responseString := w.Body.String()
	err := json.Unmarshal([]byte(responseString), &response)
	if err != nil {
		util.ApplicationLog.Printf("ERROR UNMARSHALLING RESPONSE %v\n", err)
	}
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, http.StatusCreated, response.Status)
	assert.Equal(t, "Success", response.Message)
	assert.NotNil(t, response.Timestamp)
	assert.NotNil(t, response.Data)
	t.Cleanup(func() {
		CleanUpDbOps(ctx)
	})
}


