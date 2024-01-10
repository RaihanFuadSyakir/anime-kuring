// user_controller.go
package controllers

import (
	"anime-kuring/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UserController contains CRUD methods for User
type UserController struct {
	DB *gorm.DB
}

// GetAllUsers retrieves all users from the database
// GetAllUsers retrieves all users from the database
func (u *UserController) GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	u.DB.Find(&users)
	return c.JSON(generateResponse(200, "Success", users))
}

// GetUser retrieves a user by ID from the database
func (u *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := u.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(generateResponse(404, "User not found", nil))
	}
	return c.JSON(generateResponse(200, "Success", user))
}

// CreateUser creates a new user in the database
func (u *UserController) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(generateResponse(400, "Invalid request payload", nil))
	}
	u.DB.Create(&user)
	return c.JSON(generateResponse(201, "User created successfully", user))
}

// UpdateUser updates a user by ID in the database
func (u *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := u.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(generateResponse(404, "User not found", nil))
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(generateResponse(400, "Invalid request payload", nil))
	}
	u.DB.Save(&user)
	return c.JSON(generateResponse(200, "User updated successfully", user))
}

// DeleteUser deletes a user by ID from the database
func (u *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := u.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(generateResponse(404, "User not found", nil))
	}
	u.DB.Delete(&user)
	return c.JSON(generateResponse(200, "User deleted successfully", nil))
}
