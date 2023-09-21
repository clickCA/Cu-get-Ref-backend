package models

type Review struct {
	Base
	CourseID    uint   `json:"course_id"`
	StudentID   uint   `json:"student_id"`
	ProfessorID uint   `json:"professor_id"`
	Content     string `json:"content"`
	Rating      uint   `json:"rating"`
	Anonymous   bool   `json:"anonymous"`
}
