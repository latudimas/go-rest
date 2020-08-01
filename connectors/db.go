// models/setup.go

package connectors

import (
	"context"
	"log"
	"time"

	"example.com/go-rest/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectDatabase for create conenction to mongodb
func ConnectDatabase() {
	// Database Config
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)

	// Setup Context, required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	// Cancel Context to avoid memory leak
	defer cancel()

	// Test connection to db
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to database", err)
	} else {
		log.Println("Connected!")
	}

	db := client.Database("go_rest_db")
	models.BookCollection(db)
	return
}
