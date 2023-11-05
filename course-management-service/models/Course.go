package models

type Subject struct {
	SubjectId         string `gorm:"primaryKey"`
	SubjectName       string
	CourseDescription string
	FacultyDepartment string
	AcademicTerm      string
	AcademicYear      int32
	Professors        string
	Prerequisites     string
	Status            string
	CurriculumName    string
	DegreeLevel       string
	TeachingHours     int32
}
