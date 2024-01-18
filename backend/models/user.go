package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User model
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string
	Password string
	Email    string
	Token    string
}
