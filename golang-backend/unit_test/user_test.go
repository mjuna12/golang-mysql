package unit_test

import (
	"latihan-golang/handlers"
	"latihan-golang/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestLoginValidation menguji login validasi
func TestLoginValidation(t *testing.T) {
	// Membuat Fiber app
	app := fiber.New()

	// Test Login Handler
	app.Post("/login", handlers.Login)

	// Test case untuk valid login
	t.Run("Valid Login", func(t *testing.T) {
		reqBody := `{"email":"valid@example.com", "password":"validpassword123"}`
		req := httptest.NewRequest("POST", "/login", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Contains(t, resp.Body, "Login successful")
	})

	// Test case untuk invalid login (email tidak ditemukan)
	t.Run("Invalid Login - Email Not Found", func(t *testing.T) {
		reqBody := `{"email":"nonexistent@example.com", "password":"validpassword123"}`
		req := httptest.NewRequest("POST", "/login", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		assert.Contains(t, resp.Body, "Invalid email or password")
	})

	// Test case untuk invalid login (password salah)
	t.Run("Invalid Login - Incorrect Password", func(t *testing.T) {
		reqBody := `{"email":"valid@example.com", "password":"wrongpassword"}`
		req := httptest.NewRequest("POST", "/login", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		assert.Contains(t, resp.Body, "Invalid email or password")
	})
}

func TestRegisterUserValidation(t *testing.T) {
	validate := validator.New()

	// Test case untuk registrasi dengan input yang valid
	validUser := models.RegisterUser{
		Username: "validuser",
		Email:    "validuser@example.com",
		Password: "validpassword123",
	}

	err := validate.Struct(validUser)
	assert.Nil(t, err, "Valid user should not produce an error")

	// Test case untuk registrasi dengan email kosong
	invalidUser1 := models.RegisterUser{
		Username: "validuser",
		Email:    "",
		Password: "validpassword123",
	}

	err = validate.Struct(invalidUser1)
	assert.NotNil(t, err, "Email cannot be empty")

	// Test case untuk registrasi dengan password terlalu pendek
	invalidUser2 := models.RegisterUser{
		Username: "validuser",
		Email:    "validuser@example.com",
		Password: "short",
	}

	err = validate.Struct(invalidUser2)
	assert.NotNil(t, err, "Password should be at least 5 characters long")

	// Test case untuk registrasi dengan username yang terlalu pendek
	invalidUser3 := models.RegisterUser{
		Username: "us",
		Email:    "validuser@example.com",
		Password: "validpassword123",
	}

	err = validate.Struct(invalidUser3)
	assert.NotNil(t, err, "Username should be between 3 and 30 characters")
}
