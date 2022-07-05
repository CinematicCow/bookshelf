package handler

import (
	"net/http"

	"github.com/Cinematiccow/bookshelf/db"
	"github.com/Cinematiccow/bookshelf/models"
	"github.com/gin-gonic/gin"
)

// structs to seralize and validate input body
type AddBookSchema struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookSchema struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

/*
	@method: GET
	@route: /books
	@description: retuns all books
*/
func GetAllBooks(ctx *gin.Context) {

	var Book []models.Book

	db.DB.Find(&Book)

	ctx.JSON(http.StatusOK, gin.H{"data": Book})

}

/*
	@method: GET
	@route: /books/:id
	@description: retuns a book by id
*/
func GetOneBook(ctx *gin.Context) {

	book, bookExists := CheckBookExists(ctx.Param("id"))

	if !bookExists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid Book id"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

/*
	@method: POST
	@route: /books
	@description: add a book
*/
func AddBook(ctx *gin.Context) {

	// validate input
	var input AddBookSchema

	if !ValidateInput(ctx, &input) {
		return
	}

	// Add book to database
	book := models.Book{Title: input.Title, Author: input.Author}
	db.DB.Create(&book)

	ctx.JSON(http.StatusCreated, gin.H{"data": book})

}

/*
	@method: PATCH
	@route: /books/:id
	@description: update a book by id
*/
func UpdateBook(ctx *gin.Context) {

	// check if book exists
	book, bookExists := CheckBookExists(ctx.Param("id"))

	if !bookExists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid Book id"})
		return
	}

	// validate update schema
	var input UpdateBookSchema

	if !ValidateInput(ctx, &input) {
		return
	}

	updatedBook := models.Book{Title: input.Title, Author: input.Author}
	db.DB.Model(&book).Updates(updatedBook)

	ctx.JSON(http.StatusCreated, gin.H{"data": book})

}

/*
	@method: DELETE
	@route: /books/:id
	@description: delete a book by id
*/
func DeleteBook(ctx *gin.Context) {

	// check if book exists
	book, bookExists := CheckBookExists(ctx.Param("id"))

	if !bookExists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid Book id"})
		return
	}

	// delete the book

	db.DB.Delete(book)
	ctx.JSON(http.StatusOK, gin.H{"status": "book deleted"})

}
