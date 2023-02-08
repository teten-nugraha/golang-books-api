package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teten-nugraha/books_api/controller"
	"github.com/teten-nugraha/books_api/middlewares"
)

func TestRoutes(route *gin.Engine) {
	testRoutes := route.Group("/api/test", middlewares.JWTAuthMiddleware())
	{
		testRoutes.GET("/middleware", controller.TestAPI)
	}
}
