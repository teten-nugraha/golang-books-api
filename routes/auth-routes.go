package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teten-nugraha/books_api/controller"
)

func AuthRoutes(route *gin.Engine, controller controller.AuthController) {
	authRoutes := route.Group("api/auth")
	{
		authRoutes.POST("/register", controller.Register)
	}
}
