package routes

import (
	"books_api/controller"
	"books_api/middlewares"

	"github.com/gin-gonic/gin"
)

func BookRoutes(route *gin.Engine, bookHandler controller.BookController) {
	bookRoutes := route.Group("api/books", middlewares.JWTAuthMiddleware())
	{
		bookRoutes.GET("/", bookHandler.All)
		bookRoutes.POST("/", bookHandler.Insert)
		bookRoutes.GET("/:id", bookHandler.FindByID)
		bookRoutes.PUT("/:id", bookHandler.Update)
		bookRoutes.DELETE("/:id", bookHandler.Delete)
	}
}
