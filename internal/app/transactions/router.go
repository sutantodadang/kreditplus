package transactions

import (
	"kreditplus/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterTransactionRoute(router *gin.Engine, handler ITransactionHandler) {
	transactionRoute := router.Group("/api/v1/transactions")
	transactionRoute.POST("/create", middlewares.AuthMiddleware(), handler.CreateTransaction)

}
