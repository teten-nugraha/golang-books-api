package controller

import (
	"books_api/dto/request"
	"books_api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(context *gin.Context)
	Login(context *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func (a authController) Login(context *gin.Context) {
	var input request.AuthenticationInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := a.authService.Login(input)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
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
