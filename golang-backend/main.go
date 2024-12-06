package main

import (
	"latihan-golang/database"
	"latihan-golang/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Koneksi ke database dan migrasi tabel
	database.ConnectDatabase()

	// Membuat aplikasi Fiber
	app := fiber.New()

	// Tambahkan middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",       // Domain frontend yang diizinkan
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS", // Metode HTTP yang diizinkan
		AllowHeaders:     "Content-Type,Authorization",  // Header yang diizinkan
		AllowCredentials: true,                          // Jika menggunakan cookie atau kredensial
	}))

	// Setup routes
	routes.SetupRoutes(app)
	routes.ProductRoutes(app)
	routes.CategoryRoutes(app)

	// Jalankan server pada port 3000
	app.Listen(":3000")
}
