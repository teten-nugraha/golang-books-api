package repository

import (
	"books_api/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(book models.Book) models.Book
	UpdateBook(book models.Book) models.Book
	DeleteBook(book models.Book)
	AllBook() []models.Book
	FindBookById(bookId uint64) models.Book
}

type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{connection: dbConn}
}

func (db bookConnection) InsertBook(book models.Book) models.Book {
	db.connection.Save(&book)
	db.connection.Preload("User").Find(&book)
	return book
}

func (db bookConnection) UpdateBook(book models.Book) models.Book {
	db.connection.Save(&book)
	db.connection.Preload("User").Find(&book)
	return book
}

func (db bookConnection) DeleteBook(book models.Book) {
	db.connection.Delete(&book)
}

func (db bookConnection) AllBook() []models.Book {
	var books []models.Book
	db.connection.Preload("User").Find(&books)
	return books
}

func (db bookConnection) FindBookById(bookId uint64) models.Book {
	var book models.Book
	db.connection.Preload("User").Find(&book, bookId)
	return book
}
