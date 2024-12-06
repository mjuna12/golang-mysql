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

func GetUserCourses(c *fiber.Ctx) error {
	userID := c.Params("user_id")

	var userCourses []models.UserCourse
	if err := database.DB.Where("user_id = ?", userID).Preload("Course").Find(&userCourses).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "UserCourses not found", "error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "UserCourses found", "data": userCourses})
}

func UpdateUserCourse(c *fiber.Ctx) error {
	userCourseID := c.Params("id")

	var userCourse models.UserCourse
	if err := database.DB.First(&userCourse, userCourseID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "UserCourse not found", "error": err.Error()})
	}

	// Parsing updated data from request body
	if err := c.BodyParser(&userCourse); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input", "error": err.Error()})
	}

	// Save the updated UserCourse data
	if err := database.DB.Save(&userCourse).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update userCourse", "error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "UserCourse updated successfully", "data": userCourse})
}

func DeleteUserCourse(c *fiber.Ctx) error {
	userCourseID := c.Params("id")

	var userCourse models.UserCourse
	if err := database.DB.First(&userCourse, userCourseID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "UserCourse not found", "error": err.Error()})
	}

	// Delete UserCourse from database
	if err := database.DB.Delete(&userCourse).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to delete userCourse", "error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "UserCourse deleted successfully"})
}
