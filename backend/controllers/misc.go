package controllers

import "github.com/gofiber/fiber/v2"

// generateResponse generates a standardized JSON response
type Response struct {
	OK     bool        `json:"ok"`
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func jsonResponse(c *fiber.Ctx, status int, msg string, data interface{}) error {
	response := Response{
		OK:     status >= 200 && status < 300,
		Status: status,
		Msg:    msg,
		Data:   data,
	}
	return c.Status(status).JSON(response)
}
