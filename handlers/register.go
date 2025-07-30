package handlers

import (
	"web-parfum/backend/database"
	"web-parfum/backend/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengenkripsi password",
		})
	}

	user := models.User{
		Username: data["username"],
		Password: string(password),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal membuat akun",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
