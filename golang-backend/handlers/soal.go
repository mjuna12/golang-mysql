package handlers

import (
	"latihan-golang/database"
	"latihan-golang/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Menampilkan daftar peserta didik dan mata kuliah yang diikutinya dengan mentor bergelar sarjana
func GetStudentsWithSkomOrST(c *fiber.Ctx) error {
	var userCourses []models.UserCourse

	// Mengambil data user_courses yang mentornya bergelar 'S.Kom' atau 'S.T.'
	if err := database.DB.
		Preload("User").
		Preload("Course").
		Joins("JOIN courses c ON c.id = user_courses.course_id").
		Where("c.mentor_title IN ('S.Kom', 'S.T.')").
		Find(&userCourses).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve students with Sarjana mentor",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Students with Sarjana mentor",
		"data":    userCourses,
	})
}

// Menampilkan daftar peserta didik dan mata kuliah yang diikutinya, yang mentornya bergelar selain sarjana
func GetStudentsWithNonSkomMentors(c *fiber.Ctx) error {
	var userCourses []models.UserCourse

	// Mengambil data user_courses yang mentornya tidak bergelar 'S.Kom' atau 'S.T.'
	if err := database.DB.
		Preload("User").
		Preload("Course").
		Joins("JOIN courses c ON c.id = user_courses.course_id").
		Where("c.mentor_title NOT IN ('S.Kom', 'S.T.')").
		Find(&userCourses).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve students with non-Sarjana mentor",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Students with non-Sarjana mentor",
		"data":    userCourses,
	})
}

// Menampilkan jumlah peserta didik per mata kuliah
func GetStudentCountPerCourse(c *fiber.Ctx) error {
	var result []struct {
		CourseName   string `json:"course_name"`
		StudentCount int    `json:"student_count"`
	}

	// Mengambil jumlah peserta didik per mata kuliah
	if err := database.DB.
		Select("courses.course_name, COUNT(user_courses.user_id) AS student_count").
		Joins("JOIN courses ON courses.id = user_courses.course_id").
		Group("courses.course_name").
		Scan(&result).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve student count per course",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Student count per course",
		"data":    result,
	})
}

// Menampilkan jumlah peserta didik dan total fee per mentor
func GetStudentCountAndTotalFeePerMentor(c *fiber.Ctx) error {
	var result []struct {
		MentorName   string `json:"mentor_name"`
		MentorTitle  string `json:"mentor_title"`
		StudentCount int    `json:"student_count"`
		TotalFee     int    `json:"total_fee"`
	}

	// Mengambil jumlah peserta didik dan total fee per mentor
	if err := database.DB.
		Select("courses.mentor_name, courses.mentor_title, COUNT(user_courses.user_id) AS student_count, COUNT(user_courses.user_id) * 2000000 AS total_fee").
		Joins("JOIN courses ON courses.id = user_courses.course_id").
		Group("courses.mentor_name, courses.mentor_title").
		Scan(&result).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve student count and total fee per mentor",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Student count and total fee per mentor",
		"data":    result,
	})
}
