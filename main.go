package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tomdebout/first_GOLang_API_rest/controllers"
)

func main() {
	controllers.InitDatabase()
	r := gin.Default()
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.Run()
}
