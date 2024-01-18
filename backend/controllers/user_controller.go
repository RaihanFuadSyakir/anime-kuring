// user_controller.go
package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// UserController contains CRUD methods for User
type UserController struct {
	DB *mongo.Client
}
