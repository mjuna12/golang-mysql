package models

import "gorm.io/gorm"

type UserCourse struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	CourseID  uint           `gorm:"not null" json:"course_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Course    Course         `gorm:"foreignKey:CourseID" json:"course"`
	CreatedAt gorm.DeletedAt `json:"created_at"`
	UpdatedAt gorm.DeletedAt `json:"updated_at"`
}
