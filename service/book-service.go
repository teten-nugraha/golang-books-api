package service

import (
	"fmt"
	"log"

	"books_api/dto/request"
	"books_api/models"
	"books_api/repository"

	"github.com/mashingan/smapping"
)

type BookService interface {
	Insert(b request.BookCreateDto) models.Book
	Update(b request.BookUpdateDto) models.Book
	Delete(b models.Book)
	All() []models.Book
	FindByID(bookID uint64) models.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func (service bookService) Insert(b request.BookCreateDto) models.Book {
	book := models.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.bookRepository.InsertBook(book)
	return res
}

func (service bookService) Update(b request.BookUpdateDto) models.Book {
	book := models.Book{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.bookRepository.UpdateBook(book)
	return res
}

func (service bookService) Delete(b models.Book) {
	service.bookRepository.DeleteBook(b)
}

func (service bookService) All() []models.Book {
	return service.bookRepository.AllBook()
}

func (service bookService) FindByID(bookID uint64) models.Book {
	return service.bookRepository.FindBookById(bookID)
}

func (service bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	b := service.bookRepository.FindBookById(bookID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}
