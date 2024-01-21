// db_connector.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB initializes a connection to the PostgreSQL database
func ConnectDB() (*mongo.Client, error) {
	// Load environment variables from .env file
	// if err := godotenv.Load(); err != nil {
	// 	return nil, err
	// }
	uri := os.Getenv("MONGODB_URI")
	fmt.Println("this is", uri)
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return client, nil
}
