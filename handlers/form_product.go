package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func FormProductHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "anda berada di halaman form product",
		"user_id": userID,
	})
}
