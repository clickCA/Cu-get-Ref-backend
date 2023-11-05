package models

type Course struct {
	CourseId          string `gorm:"primaryKey"`
	CourseName        string
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
