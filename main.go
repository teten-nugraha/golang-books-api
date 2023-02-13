package main

import (
	"books_api/config"
	"books_api/controller"
	"books_api/repository"
	"books_api/routes"
	"books_api/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	db = config.SetupDBConnection()

	userRepository = repository.NewUserRepository(db)
	bookRepository = repository.NewBookRepository(db)

	authService = service.NewAuthService(userRepository)
	bookService = service.NewBookService(bookRepository)
	userService = service.NewUserService(userRepository)

	authController = controller.NewAuthController(authService)
	bookController = controller.NewBookController(bookService, authService)
	userController = controller.NewUserController(userService, authService)
)

func main() {
	defer config.CloseDBConnection(db)

	loadEnv()
	loadRoutes()
}

func loadRoutes() {
	r := gin.Default()

	routes.AuthRoutes(r, authController)
	routes.TestRoutes(r)
	routes.BookRoutes(r, bookController)
	routes.UserRoutes(r, userController)

	//Add routes for check health and readiness
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "UP",
		})
	})
	r.GET("/readiness", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "READY",
		})
	})

	err := r.Run(os.Getenv("API_PORT"))
	if err != nil {
		return
	}
}

func loadEnv() {

	args := os.Args[1:]
	env := args[0]
	var activeEnv string = ".env"
	if "cloud" == env {
		activeEnv = ".env-cloud"
	}

	err := godotenv.Load(activeEnv)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Application Running using " + activeEnv)
}
