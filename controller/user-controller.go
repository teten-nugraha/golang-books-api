package controller

import (
	"books_api/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	authService service.AuthService
}

func (service userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	userID, err := service.authService.GetUserIDByToken(authHeader)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process request"})
		return
	}
	id := fmt.Sprintf("%v", userID)
	user := service.userService.Profile(id)
	context.JSON(http.StatusOK, gin.H{"profile": user})
}

func NewUserController(userService service.UserService, authService service.AuthService) UserController {
	return &userController{
		userService: userService,
		authService: authService,
	}
}
