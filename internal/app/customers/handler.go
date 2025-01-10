package customers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ICustomerHandler interface {
	CreateCustomer(ctx *gin.Context)
}

type CustomerHandler struct {
	service *CustomerService
}

func NewCustomerHandler(service *CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (h *CustomerHandler) CreateCustomer(ctx *gin.Context) {

	var req CreateCustomerRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, ok := ctx.Get("user_id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user_id in context"})
		return
	}

	req.UserID = id.(string)

	err = h.service.CreateCustomer(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}
