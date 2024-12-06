package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"type:varchar(100);not null" json:"name"`
	Email       string       `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password    string       `gorm:"type:varchar(255);not null" json:"-"`
	PhoneNumber string       `gorm:"type:varchar(15)" json:"phone_number"`
	Address     string       `gorm:"type:text" json:"address"`
	UserCourses []UserCourse `gorm:"foreignKey:UserID" json:"user_courses"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type RegisterUser struct {
	Username string `json:"username" validate:"required,min=3,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (u *User) Validate() error {
	return validate.Struct(u)
}

func (l *LoginUser) Validate() error {
	return validate.Struct(l)
}

func (l *RegisterUser) Validate() error {
	return validate.Struct(l)
}
