package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"walletEngine/configs"
	"walletEngine/data/repository"
	"walletEngine/handlers"
	"walletEngine/routes"
	"walletEngine/service"
	"walletEngine/util"
)

var(
	appRouter routes.Router
)
func main() {

	router := gin.Default()
	configs.ConnectDB()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	appRouter.WalletRouter(router)

	err := router.Run()

	if err != nil {
		return
	}

	server := &http.Server{
		Addr:    ":" + configs.EnvHTTPPort(),
		Handler: router,
	}

	go func() {
		util.ApplicationLog.Println("Starting server on port " + configs.EnvHTTPPort())
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			util.ApplicationLog.Printf("error listen: %v\n", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)


	sig := <-c
	util.ApplicationLog.Println("Received signal:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		util.ApplicationLog.Fatal("Server forced to shutdown:", err)
	}

	util.ApplicationLog.Println("Server exiting....")
}

func init() {
	appRepository :=repository.CreateMongoRepository()
	walletService := service.CreateWalletService(appRepository)
	transactionService := service.CreateTransactionService(appRepository)
	walletHandler := handlers.CreateWalletHandler(walletService)
	transactionHandler := handlers.CreateTransactionHandler(&transactionService, walletService)
	appRouter = routes.CreateRouter(walletHandler,transactionHandler)
}

