// db_migrations.go
package main

import (
	"anime-kuring/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Migration represents a database migration
type Migration struct {
	ID      primitive.ObjectID `bson:"_id"`
	Version string             `bson:"version"`
}

// MigrateDB performs database migrations
func MigrateDB(client *mongo.Client, desiredVersion string) error {
	fmt.Println("migrating ")
	// Get the database from the MongoDB client
	db := client.Database("anime-kuring")
	// Check if migration collection exists
	migrationCollection := db.Collection("migrations")
	animeCollection := db.Collection("animes")
	var result Migration
	var res models.Anime
	filter := bson.D{{}}
	data_err := animeCollection.FindOne(context.TODO(), filter).Decode(&res)
	if data_err != nil {
		data, _ := parseData()
		animeCollection.InsertMany(context.TODO(), convertSlice[models.Anime](data))
		fmt.Println("insert animes data")
	}
	err := migrationCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("no migrations collection")
		migrationCollection.InsertOne(context.TODO(), Migration{Version: desiredVersion})
		return nil
	}
	_ = migrationCollection.FindOne(context.TODO(), filter).Decode(&result)
	if result.Version != desiredVersion {
		update := bson.D{{"$set", Migration{Version: desiredVersion}}}
		// Updates the first document that has the specified "_id" value
		_, err = migrationCollection.UpdateOne(context.TODO(), result, update)
		if err != nil {
			panic(err)
		}
		return nil
	}
	fmt.Println("same ver migration")
	return nil
}
func convertSlice[T any](data []T) []interface{} {
	output := make([]interface{}, len(data))
	for idx, item := range data {
		output[idx] = item
	}
	return output
}
func parseData() ([]models.Anime, error) {
	//open json file
	jsonFile, err := os.Open("anime-offline-database-minified.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return []models.Anime{}, err
	}
	fmt.Println("Successfully Opened metadata animes")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var data models.AnimeData
	json.Unmarshal([]byte(byteValue), &data)

	fmt.Println(data.License)
	return data.Data, err
}
