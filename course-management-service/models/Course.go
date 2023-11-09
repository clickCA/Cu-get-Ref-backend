package models

import (
	"errors"
	"regexp"
)

type Course struct {
	CourseId          string `gorm:"primaryKey"`
	CourseName        string
	CourseDescription string
	FacultyDepartment string
	AcademicTerm      string
	AcademicYear      int32
	Professors        []Professor    `gorm:"many2many:course_professors;"`
	Prerequisites     []Prerequisite `gorm:"many2many:course_prerequisites;"`
	Status            string
	CurriculumName    string
	DegreeLevel       string
	TeachingHours     int32
}
type Professor struct {
	ProfessorName string `gorm:"primaryKey"`
}

type Prerequisite struct {
	PrerequisiteId string `gorm:"primaryKey"`
}

func NewCourse(courseId string, courseName string, courseDescription string, facultyDepartment string, academicTerm string, academicYear int32, professors []Professor, prerequisites []Prerequisite, status string, curriculumName string, degreeLevel string, teachingHours int32) (*Course, error) {
	match, _ := regexp.MatchString(`\d{7}`, courseId)
	if !match {
		return nil, errors.New("courseId must be 7 digits")
	}
	if courseName == "" {
		return nil, errors.New("courseName cannot be empty")
	}
	if courseDescription == "" {
		return nil, errors.New("courseDescription cannot be empty")
	}
	if facultyDepartment == "" {
		return nil, errors.New("facultyDepartment cannot be empty")
	}
	if academicTerm == "" {
		return nil, errors.New("academicTerm cannot be empty")
	}
	if academicYear == 0 {
		return nil, errors.New("academicYear cannot be empty")
	}
	match, _ = regexp.MatchString(`(open)|(close)`, status)
	if !match {
		return nil, errors.New("status must be open or close")
	}
	if curriculumName == "" {
		return nil, errors.New("curriculumName cannot be empty")
	}
	match, _ = regexp.MatchString(`(bachelor)|(master)|(doctoral)`, degreeLevel)
	if !match {
		return nil, errors.New("degreeLevel must be bachelor, master or doctoral")
	}
	if teachingHours <= 0 {
		return nil, errors.New("teachingHours must be greater than 0")
	}
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
	return course, nil
}
