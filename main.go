package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type book struct {
	ID     string `json:"id,omitempty" binding:"-"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books = []book{
	{ID: uuid.New().String(), Title: "The Go Programming Language", Author: "Alan Donovan, Brian Kernighan"},
	{ID: uuid.New().String(), Title: "Docker: Up & Running", Author: "Sean P. Kane"},
	{ID: uuid.New().String(), Title: "The Kubernetes Book", Author: "Nigel Poulton"},
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"status": "UP"})
}
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}
func postBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	newBook.ID = uuid.New().String()
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}
func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/health", healthCheck)
	router.POST("/books", postBook)
	router.Run("0.0.0.0:8080")
}
