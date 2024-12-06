package routes

import (
	"latihan-golang/handlers"
	"latihan-golang/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Register Routes
	app.Post("/register", handlers.Register)
	// Login Routes
	app.Post("/login", handlers.Login)
	// Route yang memerlukan autentikasi
	app.Get("/protected", middleware.JWTProtected, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "This is a protected route"})
	})
}
