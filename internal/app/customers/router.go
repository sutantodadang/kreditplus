package customers

import (
	"kreditplus/internal/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterCustomerRoute(router *gin.Engine, handler ICustomerHandler) {
	customerRoute := router.Group("/api/v1/customers")
	customerRoute.POST("/register", middlewares.AuthMiddleware(), handler.CreateCustomer)

}
