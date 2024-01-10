package models

import "gorm.io/gorm"

// User model
type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Token    string
}
