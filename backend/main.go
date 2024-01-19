// main.go
package main

import (
	"anime-kuring/controllers"
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect to PostgreSQL database
	client, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Perform database migrations
	MigrateDB(client, "v1.0")
	// Create a new Fiber app
	app := fiber.New()
	db := client.Database("anime-kuring")
	// Define a route
	// Define routes for CRUD operations using the UserController
	//userController := controllers.UserController{DB: db}
	animeController := controllers.NewAnimeController(db.Collection("animes"))
	defineAnimesRoute(app, animeController)
	// Start the server
	port := 5000
	app.Listen(":" + strconv.Itoa(port))
}

func defineAnimesRoute(app *fiber.App, controller *controllers.AnimeController) {
	animeGroup := app.Group("")
	animeGroup.Post("/api/animes", controller.CreateAnime)
	animeGroup.Get("/api/animes", controller.GetAnimesPropagated)
	animeGroup.Get("/api/animes/:index", controller.GetAnime)
	animeGroup.Put("/api/animes/:id", controller.UpdateAnime)
	animeGroup.Delete("/api/animes/:id", controller.DeleteAnime)
}
