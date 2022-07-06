package main

import (
	"net/http"

	"github.com/Cinematiccow/bookshelf/auth"
	"github.com/Cinematiccow/bookshelf/db"
	"github.com/Cinematiccow/bookshelf/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "hello fuckaroo"})
	})

	auth.Auth()
	// Add middlewares

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "*"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     append([]string{"content-type"}, supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	}))

	r.Use(func(ctx *gin.Context) {
		supertokens.Middleware(http.HandlerFunc(
			func(rw http.ResponseWriter, r *http.Request) {
				ctx.Next()
			},
		)).ServeHTTP(ctx.Writer, ctx.Request)

		ctx.Abort()
	})

	// User route handlers
	r.GET("/getuserinfo", handler.VerifySession(nil), handler.GetLoggedUser)

	// Books route handlers
	r.GET("/books", handler.GetAllBooks)
	r.GET("/books/:id", handler.GetOneBook)
	r.POST("/books", handler.AddBook)
	r.PATCH("/books/:id", handler.UpdateBook)
	r.DELETE("/books/:id", handler.DeleteBook)

	db.ConnectDatabase()

	r.Run(":4000")
}
