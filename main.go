package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){

	gin.SetMode(gin.ReleaseMode)
	
	r := gin.New()

	r.GET("/", indexHandler)
	r.GET("/books", getBookHandler)
	r.POST("/books", createBookHandler)

	r.Run()

}

type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
}

var books = []Book{
	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
}

func indexHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func getBookHandler(c *gin.Context){
	c.JSON(http.StatusOK, books)
}

func createBookHandler(c *gin.Context) {
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	books = append(books, book)

	c.JSON(http.StatusCreated, book)
}

