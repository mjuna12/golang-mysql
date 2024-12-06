package handlers

import (
	"latihan-golang/database"
	"latihan-golang/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Create UserCourse
func CreateUserCourse(c *fiber.Ctx) error {
	var userCourse models.UserCourse
	if err := c.BodyParser(&userCourse); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input", "error": err.Error()})
	}

	if err := database.DB.Create(&userCourse).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create userCourse", "error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "UserCourse created successfully", "data": userCourse})
}
