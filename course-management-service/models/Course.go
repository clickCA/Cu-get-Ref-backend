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

func NewCourse(courseId string, courseName string, courseDescription string, facultyDepartment string, academicTerm string, academicYear int32, professors string, prerequisites string, status string, curriculumName string, degreeLevel string, teachingHours int32) *Course {
	course := new(Course)
	course.CourseId = courseId
	course.CourseName = courseName
	course.CourseDescription = courseDescription
	course.FacultyDepartment = facultyDepartment
	course.AcademicTerm = academicTerm
	course.AcademicYear = academicYear
	course.Professors = professors
	course.Prerequisites = prerequisites
	course.Status = status
	course.CurriculumName = curriculumName
	course.DegreeLevel = degreeLevel
	course.TeachingHours = teachingHours
	return course
}
