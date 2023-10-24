package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	CourseID    uint   `json:"course_id"`
	StudentID   uint   `json:"student_id"`
	ProfessorID uint   `json:"professor_id"`
	Content     string `json:"content"`
	Rating      uint   `json:"rating"`
	Anonymous   bool   `json:"anonymous"`
}
