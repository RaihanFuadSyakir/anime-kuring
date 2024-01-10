// main.go
package main

import (
	"anime-kuring/controllers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to PostgreSQL database
	var err error
	db, err = ConnectDB()
	if err != nil {
		panic("Failed to connect to the database")
	}
	// Perform database migrations
	MigrateDB(db, "v1.1")

	// Create a new Fiber app
	app := fiber.New()
	// Define a route
	// Define routes for CRUD operations using the UserController
	userController := controllers.UserController{DB: db}
	app.Get("api/users", userController.GetAllUsers)
	app.Get("api/users/:id", userController.GetUser)
	app.Post("api/users", userController.CreateUser)
	app.Put("api/users/:id", userController.UpdateUser)
	app.Delete("api/users/:id", userController.DeleteUser)

	// Start the server
	port := 3000
	app.Listen(":" + strconv.Itoa(port))
}
