package handlers

import (
	"web-parfum/backend/database"
	"web-parfum/backend/models"

	"github.com/gofiber/fiber/v2"
)

func DashboardHandler(c *fiber.Ctx) error {
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
	})
}
