package service

import (
	"github.com/teten-nugraha/books_api/dto/request"
	"github.com/teten-nugraha/books_api/models"
	"github.com/teten-nugraha/books_api/repository"
)

type AuthService interface {
	CreateUser(user request.AuthenticationInput) models.User
	FindByUsername(username string) models.User
}

type authService struct {
	userRepository repository.UserRepository
}

func (service authService) CreateUser(payload request.AuthenticationInput) models.User {
	user := models.User{
		Username: payload.Username,
		Password: payload.Password,
	}

	savedUser := service.userRepository.InsertUser(user)
	return savedUser
}

func (service authService) FindByUsername(username string) models.User {
	return service.userRepository.FindByUsername(username)
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}
