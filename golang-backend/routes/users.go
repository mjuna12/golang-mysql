package routes

import (
	"latihan-golang/handlers"
	"latihan-golang/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Delete("/delete/:id", handlers.DeleteUser)
	app.Put("/update/:id", handlers.UpdateUser)
	app.Get("/protected", middleware.JWTProtected, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "This is a protected route"})
	})
}
