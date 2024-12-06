package routes

import (
	"latihan-golang/handlers"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	category := app.Group("/courses")
	category.Post("/create", handlers.CreateCourse)
	category.Delete("/delete/:id", handlers.DeleteCourse)
	category.Put("/update/:id", handlers.UpdateCourse)
	category.Get("/", handlers.GetAllCourses)
	// category.Get("/:id", handlers.GetAllProductByCategory)
}
