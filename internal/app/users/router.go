package users

import "github.com/gin-gonic/gin"

func RegisterUserRoute(router *gin.Engine, handler IUserHandler) {
	userRoute := router.Group("/api/v1/users")
	userRoute.POST("/register", handler.RegisterUser)
	userRoute.POST("/login", handler.LoginUser)

}
