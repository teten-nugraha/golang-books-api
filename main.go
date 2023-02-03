package main

import (
	"github.com/joho/godotenv"
	db "github.com/teten-nugraha/books_api/config"
	"github.com/teten-nugraha/books_api/models"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
}

func loadDatabase() {
	db.ConnectDB()
	db.Database.AutoMigrate(&models.User{})
	db.Database.AutoMigrate(&models.Book{})
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
