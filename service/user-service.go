package service

import (
	"books_api/models"
	"books_api/repository"
)

type UserService interface {
	Profile(userID string) models.User
}

type userService struct {
	userRepository repository.UserRepository
}

func (service userService) Profile(userID string) models.User {
	return service.userRepository.ProfileUser(userID)
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}
