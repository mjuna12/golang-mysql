package handlers

import (
	"latihan-golang/database"
	"latihan-golang/models"
	"latihan-golang/services"
	"latihan-golang/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user models.RegisterUser
	// Parsing body request ke struct RegisterUser
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	// Validasi input user
	if err := user.Validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	// Cek apakah email sudah terdaftar
	var existingUser models.User
	if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Email is already registered",
		})
	}

	// Hash password sebelum disimpan
	hashedPassword := services.HashingPassword(user.Password)

	// Simpan user ke database
	newUser := models.User{
		Name:     user.Username, // Menggunakan Username dari RegisterUser
		Email:    user.Email,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}

	// Response sukses
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"data": fiber.Map{
			"username": newUser.Name,
			"email":    newUser.Email,
		},
	})
}

func Login(c *fiber.Ctx) error {
	var loginReq models.LoginUser
	if err := c.BodyParser(&loginReq); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Invalid input", "error": err.Error()})
	}

	// Cek user berdasarkan email
	var user models.User
	if err := database.DB.Where("email = ?", loginReq.Email).First(&user).Error; err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	// Periksa password
	if !services.CheckPasswordHash(loginReq.Password, user.Password) {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid email or password"})
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to generate token", "error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	// Mencari user berdasarkan ID
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Menghapus user dari database
	if err := database.DB.Delete(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete user",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	// Mendapatkan ID dari URL params
	userID := c.Params("id")

	// Mencari user berdasarkan ID
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Memparsing body request untuk update data
	var updateData models.User
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	user.Name = updateData.Name
	user.Email = updateData.Email
	user.PhoneNumber = updateData.PhoneNumber
	user.Address = updateData.Address

	// Menyimpan perubahan di database
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update user",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
		"user":    user,
	})
}
