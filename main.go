package main

import (
	"github.com/gin-gonic/gin"
	"walletEngine/configs"
)

func main() {

	router := gin.Default()
	configs.ConnectDB()

	err := router.Run()

	if err != nil {
		return
	}
}
