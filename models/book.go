// models/book.go

package models

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Book Data Structure
type Book struct {
	// ID        string    `bson:"_id, omitempty"`
	Title     string    `bson:"title, omitempty"`
	Author    string    `bson:"author, omitempty"`
	CreatedAt time.Time `bson:"created_at, omitempty"`
	UpdatedAt time.Time `bson:"updated_at, omitempty"`
}

// Database Instance
var collection *mongo.Collection

// BookCollection for define collection used for Book model
func BookCollection(c *mongo.Database) {
	collection = c.Collection("books_collection")
}

// GetAllBooks for get list of all book in db
func GetAllBooks(c *gin.Context) {
	books := []Book{}
	cursor, err := collection.Find(context.TODO(), bson.M{})

	// Handling error
	if err != nil {
		log.Printf("Error while getting all books, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wron",
		})
		return
	}

	// Iterate through the returned curson
	for cursor.Next(context.TODO()) {
		var book Book
		cursor.Decode(&book)
		books = append(books, book)
	}

	log.Printf("Successfully GET all book")
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Todos",
		"data":    books,
	})

	return
}

// CreateBook for inserting book into db
func CreateBook(c *gin.Context) {
	var book Book
	c.BindJSON(&book)
	title := book.Title
	author := book.Author

	newBook := Book{
		// ID:        guuid.New().String(),
		Title:     title,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := collection.InsertOne(context.TODO(), newBook)

	if err != nil {
		log.Printf("Error while inserting new book into db, Reason %v\n", err)

		// Return 500
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})

		return
	}

	log.Printf("Successfully POST a Book")
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully Create Book",
	})

	return
}
