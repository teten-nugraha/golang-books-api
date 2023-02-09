package routes

import (
	"books_api/controller"
	"books_api/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine, userController controller.UserController) {
	userRoutes := route.Group("api/user", middlewares.JWTAuthMiddleware())
	{
		userRoutes.GET("/profile", userController.Profile)
	}
}
