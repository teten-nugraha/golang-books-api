package main

import (
	"log"
	"os"

	"books_api/config"
	"books_api/controller"
	"books_api/repository"
	"books_api/routes"
	"books_api/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	db = config.SetupDBConnection()

	userRepository = repository.NewUserRepository(db)

	authService = service.NewAuthService(userRepository)

	authController = controller.NewAuthController(authService)
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

	err := r.Run(os.Getenv("API_PORT"))
	if err != nil {
		return
	}
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
