package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/teten-nugraha/books_api/config"
	"github.com/teten-nugraha/books_api/controller"
	"github.com/teten-nugraha/books_api/repository"
	"github.com/teten-nugraha/books_api/routes"
	"github.com/teten-nugraha/books_api/service"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db *gorm.DB = config.SetupDBConnection()

	userRepository repository.UserRepository = repository.NewUserRepository(db)

	authService service.AuthService = service.NewAuthService(userRepository)

	authController controller.AuthController = controller.NewAuthController(authService)
)

func main() {
	defer config.CloseDBConnection(db)

	loadEnv()
	loadRoutes()
}

func loadRoutes() {
	r := gin.Default()

	routes.AuthRoutes(r, authController)

	r.Run(os.Getenv("API_PORT"))
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
