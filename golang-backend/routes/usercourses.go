package routes

import (
	"latihan-golang/handlers"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	product := app.Group("/usercourses")
	// product.Get("/", handlers.GetAllProduct)
	product.Post("/create", handlers.CreateUserCourse)
	// product.Delete("delete/:id", handlers.DeleteProduct)
	// product.Get("detail/:id", handlers.DetailProduct)
	// product.Put("/update/:id", handlers.UpdateProduct)
}
