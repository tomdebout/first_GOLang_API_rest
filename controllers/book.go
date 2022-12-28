package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tomdebout/first_GOLang_API_rest/models"
)

var Library []models.Book
var Counter int

func InitDatabase() {
	Counter = 1

	book1 := models.Book{
		ID:     1,
		Title:  "Le langage Go: Les fondamentaux du langage",
		Author: "Frédéric Boulanger",
	}
	Library = append(Library, book1)
}

func FindBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": Library})
}

func CreateBook(c *gin.Context) {
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Counter++
	book := models.Book{ID: Counter, Title: input.Title, Author: input.Author}
	Library = append(Library, book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

func removeIt(book models.Book, bookSlice []models.Book) []models.Book {
	for idx, v := range bookSlice {
		if v.ID == book.ID {
			return append(bookSlice[:idx], bookSlice[idx+1:]...)
		}
	}
	return bookSlice
}

func DeleteBook(c *gin.Context) {
	bookFound := false

	var bookFind models.Book

	for _, book := range Library {
		if c.Param("id") == strconv.Itoa(book.ID) {
			bookFound = true
			bookFind = book
		}
	}

	if !bookFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	Library = removeIt(bookFind, Library)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
