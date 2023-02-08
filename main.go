package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/teten-nugraha/books_api/config"
	"github.com/teten-nugraha/books_api/controller"
	"github.com/teten-nugraha/books_api/repository"
	"github.com/teten-nugraha/books_api/routes"
	"github.com/teten-nugraha/books_api/service"
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
