package routes

import (
	"latihan-golang/handlers"

	"github.com/gofiber/fiber/v2"
)

func SoalRoutes(app *fiber.App) {
	courses := app.Group("/soal")
	courses.Get("/1", handlers.GetStudentsWithSkomOrST)
	courses.Get("/2", handlers.GetStudentsWithNonSkomMentors)
	courses.Get("/3", handlers.GetStudentCountPerCourse)
	courses.Get("/4", handlers.GetStudentCountAndTotalFeePerMentor)
}
