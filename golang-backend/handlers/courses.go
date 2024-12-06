package handlers

import (
	"latihan-golang/database"
	"latihan-golang/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Create Course
func CreateCourse(c *fiber.Ctx) error {
	var course models.Course
	if err := c.BodyParser(&course); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input", "error": err.Error()})
	}

	if err := database.DB.Create(&course).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create course", "error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "Course created successfully", "data": course})
}

// Get All Courses
func GetAllCourses(c *fiber.Ctx) error {
	var courses []models.Course
	if err := database.DB.Find(&courses).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch courses", "error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Courses fetched successfully", "data": courses})
}

// Update Course
func UpdateCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	var course models.Course
	if err := database.DB.Where("id = ?", id).First(&course).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Course not found"})
	}

	if err := c.BodyParser(&course); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input", "error": err.Error()})
	}

	if err := database.DB.Save(&course).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to update course", "error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Course updated successfully", "data": course})
}

// Delete Course
func DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.Course{}, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "Course not found"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Course deleted successfully"})
}
