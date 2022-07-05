package main

import (
	"net/http"

	"github.com/Cinematiccow/bookshelf/db"
	"github.com/Cinematiccow/bookshelf/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "hello fuckaroo"})
	})

	r.GET("/books", handler.GetAllBooks)
	r.GET("/books/:id", handler.GetOneBook)
	r.POST("/books", handler.AddBook)
	r.PATCH("/books/:id", handler.UpdateBook)
	r.DELETE("/books/:id", handler.DeleteBook)

	db.ConnectDatabase()

	r.Run(":4000")
}
