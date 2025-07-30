package middleware

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: JwtError,
	})
}

func JwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Token JWT hilang atau salah format",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Token tidak valid atau kadaluarsa",
	})
}