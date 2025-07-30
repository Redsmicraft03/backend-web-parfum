package handlers

import (
	"context"
	"path/filepath"
	"strings"
	"web-parfum/backend/config"
	"web-parfum/backend/database"
	"web-parfum/backend/models"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func DeleteProductHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid product ID",
		})
	}

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Product not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "database error",
		})
	}

	// hapus gambar
	if product.ImageURL != "" {
		// Ekstrak Public ID dari URL Cloudinary
		// Contoh URL: https://res.cloudinary.com/cloud-name/image/upload/v123/parfum_products/123_product.jpg
		// Public ID: parfum_products/123_product
		publicID := extractPublicIDFromURL(product.ImageURL)

		// Hapus gambar menggunakan API Destroy
		_, err := config.Cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
			PublicID: publicID,
		})

		if err != nil {
			// Sebaiknya hanya di-log, jangan sampai gagal hapus gambar membuat seluruh proses gagal
			// Tapi untuk sekarang kita kembalikan error
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to delete image from Cloudinary",
			})
		}
	}
	// hapus data di database
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to delete product",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Product deleted successfully",
		"user_id": userID,
	})
}

func extractPublicIDFromURL(imageUrl string) string {
	// Pisahkan URL berdasarkan "/"
	parts := strings.Split(imageUrl, "/")

	// Cari bagian setelah "upload" dan sebelum nomor versi
	for i, part := range parts {
		if part == "upload" && i+2 < len(parts) {
			// Gabungkan semua bagian setelah nomor versi
			publicIDWithExt := strings.Join(parts[i+2:], "/")
			// Hapus ekstensi file (.jpg, .png, dll.)
			return strings.TrimSuffix(publicIDWithExt, filepath.Ext(publicIDWithExt))
		}
	}
	return ""
}
