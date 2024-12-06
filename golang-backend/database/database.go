package database

import (
	"latihan-golang/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase menginisialisasi koneksi ke database MySQL
func ConnectDatabase() {
	dsn := "root@tcp(127.0.0.1:3306)/diskominfo-test?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	// Coba membuka koneksi ke database MySQL
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return // Menghentikan proses lebih lanjut jika gagal terhubung
	}

	// Jika berhasil terkoneksi, log informasi
	log.Println("Successfully connected to the database!")

	// Auto-migrate models untuk memastikan tabel sudah ada di database
	err = DB.AutoMigrate(&models.User{}, &models.Course{}, &models.UserCourse{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
		return
	}

	log.Println("Auto-migration complete.")
}
