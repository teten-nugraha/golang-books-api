package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/teten-nugraha/books_api/dto/request"
	"github.com/teten-nugraha/books_api/models"
	"github.com/teten-nugraha/books_api/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"strings"
	"time"
)

type AuthService interface {
	CreateUser(user request.AuthenticationInput) (models.User, error)
	ValidateJWT(context *gin.Context) error
	Login(user request.AuthenticationInput) (string, error)
}

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

type authService struct {
	userRepository repository.UserRepository
}

func (service authService) Login(user request.AuthenticationInput) (string, error) {
	userExist := service.userRepository.FindByUsername(user.Username)
	if (userExist == models.User{}) {
		return "", errors.New("User not found")
	}

	err := validatePassword(userExist.Password, user.Password)
	if err != nil {
		return "", err
	}

	jwt, err := generateJwt(userExist)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func generateJwt(user models.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func (service authService) ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

func validatePassword(existPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(existPassword), []byte(password))
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

func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""

}
