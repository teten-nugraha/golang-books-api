package controller

import (
	"net/http"
	"strconv"

	"books_api/models"

	"books_api/service"

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
	//var bookCreateDTO request.BookCreateDto
	//errDTO := context.ShouldBind(&bookCreateDTO)
	//if errDTO != nil {
	//	context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process request"})
	//} else {
	//	authHeader := context.GetHeader("Authorization")
	//	userID := b.getUserIDByToken(authHeader)
	//	convertedID, err := strconv.ParseUint(userID, 10, 64)
	//	if err != nil {
	//		bookCreateDTO.UserID = convertedID
	//	}
	//
	//	result := b.bookService.Insert(bookCreateDTO)
	//	response := helpers.BuildResponse(true, "OK", result)
	//	context.JSON(http.StatusCreated, response)
	//}
}

func (b bookController) Update(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (b bookController) Delete(context *gin.Context) {
	//TODO implement me
	panic("implement me")
}

//func (b bookController) getUserIDByToken(token string) string {
//	aToken, err := b.authService.ValidateToken(token)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	claims := aToken.Claims.(jwt.MapClaims)
//	id := fmt.Sprintf("%v", claims["user_id"])
//	return id
//}

func NewBookController(bookService service.BookService, authService service.AuthService) BookController {
	return bookController{
		bookService: bookService,
		authService: authService,
	}
}
