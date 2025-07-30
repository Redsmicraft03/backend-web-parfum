package handlers

import (
	"os"
	"time"
	"web-parfum/backend/database"
	"web-parfum/backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	result := database.DB.Where("username = ?", data["username"]).First(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "username or password wrong",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "username or password wrong",
		})
	}

	// Buat klaim untuk JWT
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success login",
		"token": token,
	})
}