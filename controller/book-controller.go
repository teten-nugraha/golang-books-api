package controller

import (
	"books_api/dto/request"
	"books_api/models"
	"books_api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type bookController struct {
	bookService service.BookService
	authService service.AuthService
}

func (b bookController) All(context *gin.Context) {
	var books = b.bookService.All()
	context.JSON(http.StatusOK, gin.H{"books": books})
}

func (b bookController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "No param id was found"})
		return
	}

	var book = b.bookService.FindByID(id)
	if (book == models.Book{}) {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Data not found\", \"No data with given id"})
	} else {
		context.JSON(http.StatusOK, gin.H{"book": book})
	}
}

func (b bookController) Insert(context *gin.Context) {
	var bookCreateDTO request.BookCreateDto
	errDTO := context.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process request"})
	} else {
		authHeader := context.GetHeader("Authorization")
		userID, err := b.authService.GetUserIDByToken(authHeader)
		if err == nil {
			bookCreateDTO.UserID = userID
		}

		result := b.bookService.Insert(bookCreateDTO)
		context.JSON(http.StatusBadRequest, gin.H{"book": result})
	}
}

func (b bookController) Update(context *gin.Context) {
	var bookUpdateDTO request.BookUpdateDto
	errDTO := context.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process request"})
		return
	}

	authHeader := context.GetHeader("Authorization")
	userID, err := b.authService.GetUserIDByToken(authHeader)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process request"})
		return
	}

	if b.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) {
		bookUpdateDTO.UserID = userID
		result := b.bookService.Update(bookUpdateDTO)
		context.JSON(http.StatusOK, gin.H{"book": result})
	} else {
		context.JSON(http.StatusForbidden, gin.H{"error": "You dont have permission, You are not the owner"})
	}
}

func (b bookController) Delete(context *gin.Context) {
	var book models.Book
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "failed to get id"})
	}
	book.ID = id

	authHeader := context.GetHeader("Authorization")
	userID, err := b.authService.GetUserIDByToken(authHeader)

	if b.bookService.IsAllowedToEdit(userID, book.ID) {
		b.bookService.Delete(book)
		context.JSON(http.StatusOK, gin.H{"message": "Deleted"})
	} else {
		context.JSON(http.StatusForbidden, gin.H{"error": "You dont have permission, You are not the owner"})
	}
}

func NewBookController(bookService service.BookService, authService service.AuthService) BookController {
	return bookController{
		bookService: bookService,
		authService: authService,
	}
}
