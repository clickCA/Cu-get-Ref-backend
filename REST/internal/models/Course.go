package models

type Course struct {
	Base
	CourseName string `gorm:"type:varchar(255);notNull" json:"course_name"`
	CourseCode string `gorm:"type:varchar(255);uniqueIndex;notNull" json:"course_code"`
}
