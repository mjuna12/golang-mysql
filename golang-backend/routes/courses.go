package routes

import (
	"latihan-golang/handlers"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	courses := app.Group("/courses")
	courses.Post("/create", handlers.CreateCourse)
	courses.Delete("/delete/:id", handlers.DeleteCourse)
	courses.Put("/update/:id", handlers.UpdateCourse)
	courses.Get("/", handlers.GetAllCourses)
	courses.Get("/:id", handlers.GetCourseByID)
}
