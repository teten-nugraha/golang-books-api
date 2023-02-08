package routes

import (
	"books_api/controller"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine, controller controller.AuthController) {
	authRoutes := route.Group("api/auth")
	{
		authRoutes.POST("/register", controller.Register)
		authRoutes.POST("/login", controller.Login)
	}
}
