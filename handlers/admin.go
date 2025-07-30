package handlers

import (
	"web-parfum/backend/database"
	"web-parfum/backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AdminHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	var products []models.Product

	result := database.DB.Find(&products)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Tidak bisa mengambil data produk dari database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    products,
		"user_id": userID,
	})
}
