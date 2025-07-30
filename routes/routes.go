package routes

import (
	"os"
	"web-parfum/backend/handlers"
	"web-parfum/backend/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("CLIENT_URL"),
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	api := app.Group("/api")

	// publik
	api.Post("/login", handlers.LoginHandler)
	api.Get("/dashboard", handlers.DashboardHandler)

	// admin
	admin := api.Group("/admin", middleware.Protected())

	admin.Get("/", handlers.AdminHandler) // /api/admin/
	admin.Get("/form-product", handlers.FormProductHandler) // /api/admin/form-product
	admin.Post("/create-product", handlers.CreateProductHandler) // /api/admin/create-product
	admin.Patch("/update-product/:id", handlers.UpdateProductHandler) // /api/admin/update-product/:id
	admin.Delete("/delete-product/:id", handlers.DeleteProductHandler) // /api/admin/delete-product/:id
}
