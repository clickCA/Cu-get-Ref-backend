package models

import (
	"errors"
	"regexp"
)

type Course struct {
	CourseId          string `gorm:"primaryKey"`
	CourseName        string `gorm:"type:varchar(50)"`
	CourseDescription string `gorm:"type:text"`
	FacultyDepartment string `gorm:"type:varchar(50)"`
	AcademicTerm      string `gorm:"type:varchar(10)"`
	AcademicYear      int32  `gorm:"type:smallint"`
	Professors        string `gorm:"type:text"`
	Prerequisites     string `gorm:"type:text"`
	Status            string `gorm:"type:varchar(10)"`
	CurriculumName    string `gorm:"type:varchar(50)"`
	DegreeLevel       string `gorm:"type:varchar(10)"`
	TeachingHours     int32  `gorm:"type:smallint"`
}

func NewCourse(courseId string, courseName string, courseDescription string, facultyDepartment string, academicTerm string, academicYear int32, professors string, prerequisites string, status string, curriculumName string, degreeLevel string, teachingHours int32) (*Course, error) {
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

func UpdateCourse(oldCourse *Course, courseId string, courseName string, courseDescription string, facultyDepartment string, academicTerm string, academicYear int32, professors string, prerequisites string, status string, curriculumName string, degreeLevel string, teachingHours int32) error {
	match, _ := regexp.MatchString(`\d{7}`, courseId)
	if !match {
		return errors.New("courseId cannot be empty")
	}
	if courseName == "" {
		return errors.New("courseName cannot be empty")
	}
	if courseDescription == "" {
		return errors.New("courseDescription cannot be empty")
	}
	if facultyDepartment == "" {
		return errors.New("facultyDepartment cannot be empty")
	}
	if academicTerm == "" {
		return errors.New("academicTerm cannot be empty")
	}
	if academicYear == 0 {
		return errors.New("academicYear cannot be empty")
	}
	match, _ = regexp.MatchString(`(open)|(close)`, status)
	if !match {
		return errors.New("status must be open or close")
	}
	if curriculumName == "" {
		return errors.New("curriculumName cannot be empty")
	}
	match, _ = regexp.MatchString(`(bachelor)|(master)|(doctoral)`, degreeLevel)
	if !match {
		return errors.New("degreeLevel must be bachelor, master or doctoral")
	}
	if teachingHours <= 0 {
		return errors.New("teachingHours must be greater than 0")
	}
	oldCourse.CourseId = courseId
	oldCourse.CourseName = courseName
	oldCourse.CourseDescription = courseDescription
	oldCourse.FacultyDepartment = facultyDepartment
	oldCourse.AcademicTerm = academicTerm
	oldCourse.AcademicYear = academicYear
	oldCourse.Professors = professors
	oldCourse.Prerequisites = prerequisites
	oldCourse.Status = status
	oldCourse.CurriculumName = curriculumName
	oldCourse.DegreeLevel = degreeLevel
	oldCourse.TeachingHours = teachingHours
	return nil
}
