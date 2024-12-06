package routes

import (
	"latihan-golang/handlers"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	usercourses := app.Group("/usercourses")
	usercourses.Get("/", handlers.GetUserCourses)
	usercourses.Post("/create", handlers.CreateUserCourse)
	usercourses.Delete("delete/:id", handlers.DeleteUserCourse)
	// usercourses.Get("detail/:id", handlers.Detailusercourses)
	usercourses.Put("/update/:id", handlers.UpdateUserCourse)
}
