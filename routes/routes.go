package routes

import (
	"net/http"

	"example.com/go-rest/models"
	"github.com/gin-gonic/gin"
)

// Routes list
func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("/books", models.GetAllBooks)
	router.POST("/book", models.CreateBook)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome to api go-rest",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  200,
		"message": "Route Not Found",
	})
	return
}
