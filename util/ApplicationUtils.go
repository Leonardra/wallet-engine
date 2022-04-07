package util

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
	"walletEngine/dto"
)

var ApplicationLog = log.New(os.Stdout, "[wallet-service] ", log.LstdFlags)

func GenerateJSONResponse(c *gin.Context, statusCode int, message string, data map[string]interface{}) {
	c.JSON(statusCode, dto.APIResponse{
		Status:    statusCode,
		Message:   message,
		Timestamp: time.Now(),
		Data:      data,
	})
}

func GenerateInternalServerErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, dto.APIResponse{
		Status:    http.StatusInternalServerError,
		Message:   message,
		Timestamp: time.Now(),
		Data:      gin.H{},
	})
}

func GenerateBadRequestResponse(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, dto.APIResponse{
		Status:    http.StatusBadRequest,
		Message:   message,
		Timestamp: time.Now(),
		Data:      gin.H{},
	})
}
