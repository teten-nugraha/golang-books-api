package routes

import (
	"books_api/controller"
	"books_api/middlewares"

	"github.com/gin-gonic/gin"
)

func TestRoutes(route *gin.Engine) {
	testRoutes := route.Group("/api/test", middlewares.JWTAuthMiddleware())
	{
		testRoutes.GET("/middleware", controller.TestAPI)
	}
}
