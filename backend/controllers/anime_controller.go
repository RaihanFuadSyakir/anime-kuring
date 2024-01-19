package controllers

import (
	"anime-kuring/models"
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AnimeController struct {
	Col *mongo.Collection
}

func NewAnimeController(col *mongo.Collection) *AnimeController {
	return &AnimeController{Col: col}
}

// CreateAnime creates a new Anime.
func (ac *AnimeController) CreateAnime(c *fiber.Ctx) error {
	var newAnime models.Anime
	if err := c.BodyParser(&newAnime); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "invalid request body", nil)
	}
	result, err := ac.Col.InsertOne(context.TODO(), newAnime)
	if err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
	}
	return jsonResponse(c, fiber.StatusCreated, "anime info inserted", result)
}

// GetAnime retrieves an Anime by ID.
func (ac *AnimeController) GetAnime(c *fiber.Ctx) error {
	animeIndexStr := c.Params("index")
	animeIndex, err := strconv.Atoi(animeIndexStr)
	if err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid propogate value", nil)
	}
	// Find the anime by ID in the database
	var result models.Anime
	err = ac.Col.FindOne(context.TODO(), bson.M{"index": animeIndex}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return jsonResponse(c, fiber.StatusNotFound, "Anime not found", nil)
		}
		return jsonResponse(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
	}

	// Return the anime as JSON
	return c.JSON(result)
}

// GetAnimesPropagated retrieves a list of animes based on query parameters.
func (ac *AnimeController) GetAnimesPropagated(c *fiber.Ctx) error {
	propogateStr := c.Query("propogate") // Use your specific query parameters
	pagesStr := c.Query("page")
	// Convert propogate to int
	propogate, err := strconv.Atoi(propogateStr)
	if err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid propogate value", nil)
	}

	// Convert pages to int
	pages, err := strconv.Atoi(pagesStr)
	if err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid pages value", nil)
	}
	index := propogate*(pages-1) + 1
	// Find animes in the database based on query parameters
	filter := bson.D{{"index", bson.D{{"$gte", index}, {"$lt", index + propogate}}}}
	cursor, err := ac.Col.Find(context.TODO(), filter)
	if err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
	}
	defer cursor.Close(context.TODO())

	// Decode the results into a slice of Anime
	var animes []models.Anime
	if err := cursor.All(context.TODO(), &animes); err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
	}

	// Return the animes as JSON
	return c.JSON(animes)
}

// UpdateAnime updates an existing Anime by ID.
func (ac *AnimeController) UpdateAnime(c *fiber.Ctx) error {
	animeID := c.Params("id")

	var updateAnime models.Anime
	if err := c.BodyParser(&updateAnime); err != nil {
		return jsonResponse(c, fiber.StatusBadRequest, "Invalid request body", nil)
	}

	// Update the anime in the database
	update := bson.M{
		"$set": updateAnime,
	}

	result, err := ac.Col.UpdateOne(context.TODO(), bson.M{"_id": animeID}, update)
	if err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
	}

	return jsonResponse(c, fiber.StatusOK, "Anime updated successfully", result)
}

// DeleteAnime deletes an Anime by ID.
func (ac *AnimeController) DeleteAnime(c *fiber.Ctx) error {
	animeID := c.Params("id")

	// Delete the anime from the database
	result, err := ac.Col.DeleteOne(context.TODO(), bson.M{"_id": animeID})
	if err != nil {
		return jsonResponse(c, fiber.StatusInternalServerError, "Internal Server Error", nil)
	}

	return jsonResponse(c, fiber.StatusOK, "Anime deleted successfully", result)
}
