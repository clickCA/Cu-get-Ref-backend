package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	CourseName string `gorm:"type:varchar(255);notNull" json:"course_name"`
	CourseCode string `gorm:"type:varchar(255);uniqueIndex;notNull" json:"course_code"`
}
