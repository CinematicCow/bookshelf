package handler

import (
	"net/http"

	"github.com/Cinematiccow/bookshelf/db"
	"github.com/Cinematiccow/bookshelf/models"
	"github.com/gin-gonic/gin"
)

// check if book exists
func CheckBookExists(id string) (models.Book, bool) {

	var Book models.Book

	if err := db.DB.Where("id=?", id).First(&Book).Error; err != nil {
		return Book, false
	}

	return Book, true

}

// validate input schema
// the function takes 2 parameters, one is the gin context and another is a generic schema
func ValidateInput[inputSchema *AddBookSchema | *UpdateBookSchema](ctx *gin.Context, schema inputSchema) bool {

	if err := ctx.ShouldBindJSON(&schema); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return false
	}

	return true
}
