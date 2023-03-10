package repository

import (
	"books_api/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user models.User) models.User
	FindByUsername(username string) models.User
	ProfileUser(userId string) models.User
}

type userConnection struct {
	connection *gorm.DB
}

func (db userConnection) ProfileUser(userId string) models.User {
	var user models.User
	db.connection.Find(&user, userId)
	return user
}

func (db userConnection) InsertUser(user models.User) models.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db userConnection) FindByUsername(username string) models.User {
	var user models.User
	db.connection.Where("username=?", username).Take(&user)
	return user
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}

	return string(hash)
}

// Dependency Injection
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}
