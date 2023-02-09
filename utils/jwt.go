package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func ValidateJWT(context *gin.Context) error {
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
	return splitToken[0]

}
