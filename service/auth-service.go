package service

import (
	"errors"
	"github.com/teten-nugraha/books_api/dto/request"
	"github.com/teten-nugraha/books_api/models"
	"github.com/teten-nugraha/books_api/repository"
)

type AuthService interface {
	CreateUser(user request.AuthenticationInput) (models.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func (service authService) CreateUser(payload request.AuthenticationInput) (models.User, error) {

	// check user exist or not
	existUserByUsername := service.userRepository.FindByUsername(payload.Username)
	if !(existUserByUsername == models.User{}) {
		return models.User{}, errors.New("username already exist")
	}

	user := models.User{
		Username: payload.Username,
		Password: payload.Password,
	}

	savedUser := service.userRepository.InsertUser(user)
	return savedUser, nil
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}
