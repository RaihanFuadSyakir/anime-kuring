// main.go
package main

import (
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to PostgreSQL database
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Perform database migrations
	MigrateDB(db, "v1.0")
	// Create a new Fiber app
	app := fiber.New()
	// Define a route
	// Define routes for CRUD operations using the UserController
	//userController := controllers.UserController{DB: db}

	// Start the server
	port := 5000
	app.Listen(":" + strconv.Itoa(port))
}
