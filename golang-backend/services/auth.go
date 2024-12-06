package services

import (
	"golang.org/x/crypto/bcrypt"
)

// HashingPassword untuk mengenkripsi password
func HashingPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

// CheckPasswordHash untuk memeriksa password
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
