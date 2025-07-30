package handlers

import (
	"context"
	"web-parfum/backend/config"
	"web-parfum/backend/database"
	"web-parfum/backend/models"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func UpdateProductHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	var existingProduct models.Product

	if err := database.DB.First(&existingProduct, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "product not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "database error",
		})
	}

	var payload models.ProductPayload
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	// cek ada gambar baru atau tidak
	if payload.ImageData != "" {
		uploadResult, err := config.Cld.Upload.Upload(
			context.Background(),
			payload.ImageData,
			uploader.UploadParams{
				Folder: "parfum_products",
			},
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "failed to upload image",
			})
		}

		// timpa url lama dengan baru
		existingProduct.ImageURL = uploadResult.SecureURL
	}

	existingProduct.Name = payload.Name
	existingProduct.Description = payload.Description
	existingProduct.Price = payload.Price
	existingProduct.Link = payload.Link

	if err := database.DB.Save(&existingProduct).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success update product",
		"data":    existingProduct,
		"user_id": userID,
	})
}
