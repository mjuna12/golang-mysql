package models

import "gorm.io/gorm"

type Course struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CourseName  string         `gorm:"type:varchar(100);not null" json:"course_name"`
	MentorName  string         `gorm:"type:varchar(100);not null" json:"mentor_name"`
	MentorTitle string         `gorm:"type:varchar(50)" json:"mentor_title"`
	UserCourses []UserCourse   `gorm:"foreignKey:CourseID" json:"user_courses"`
	CreatedAt   gorm.DeletedAt `json:"created_at"`
	UpdatedAt   gorm.DeletedAt `json:"updated_at"`
}
