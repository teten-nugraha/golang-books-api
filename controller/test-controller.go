package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestAPI(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello, ini adalah private endpoint "})
}
