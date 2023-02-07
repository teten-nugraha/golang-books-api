package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teten-nugraha/books_api/dto/request"
	"github.com/teten-nugraha/books_api/service"
	"net/http"
)

type AuthController interface {
	Register(context *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func (a authController) Register(context *gin.Context) {
	var payload request.AuthenticationInput
	if err := context.ShouldBindJSON(&payload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedUser, err := a.authService.CreateUser(payload)
	if err != nil {
		context.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}
