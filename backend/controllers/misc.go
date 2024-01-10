package controllers

import "github.com/gofiber/fiber/v2"

// generateResponse generates a standardized JSON response
func generateResponse(status int, msg string, data interface{}) fiber.Map {
	return fiber.Map{
		"status": status,
		"msg":    msg,
		"data":   data,
	}
}
