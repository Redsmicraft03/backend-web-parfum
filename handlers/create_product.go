	package handlers

	import (
		"context"
		"web-parfum/backend/config"
		"web-parfum/backend/database"
		"web-parfum/backend/models"

		"github.com/cloudinary/cloudinary-go/v2/api/uploader"
		"github.com/gofiber/fiber/v2"
		"github.com/golang-jwt/jwt/v5"
	)

	func CreateProductHandler(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(float64)

		var product models.ProductPayload

		if err := c.BodyParser(&product); err != nil {
			return err
		}

		uploadResult, err := config.Cld.Upload.Upload(
			context.Background(),
			product.ImageData,
			uploader.UploadParams{
				Folder: "parfum_products",
			},
		)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "failed to upload image",
			})
		}

		newProduct := models.Product{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Link:        product.Link,
			ImageURL:    uploadResult.SecureURL,
		}

		if err := database.DB.Create(&newProduct).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "failed to create product",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "success create product",
			"data":    newProduct,
			"user_id": userID,
		})
	}
