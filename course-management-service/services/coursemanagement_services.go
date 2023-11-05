package services

import (
	"context"
	"course-management-service/config"
	course_management "course-management-service/coursemanagement"
	"course-management-service/models"
	"log"
)

type Server struct {
	course_management.UnimplementedCourseManagementServiceServer
}

var db = config.ConnectDB()

func (s *Server) GetAllCourses(ctx context.Context, req *course_management.Empty) (*course_management.CourseList, error) {
	log.Println("GetAllCourses")
	// Get all course in database
	var courses []models.Course
	db.Find(&courses)
	// Convert to protobuf
	var courseList []*course_management.CourseItem
	for _, subject := range courses {
		courseList = append(courseList, &course_management.CourseItem{
			CourseId:          subject.CourseId,
			CourseName:        subject.CourseName,
			CourseDescription: subject.CourseDescription,
			FacultyDepartment: subject.FacultyDepartment,
			AcademicTerm:      subject.AcademicTerm,
			AcademicYear:      subject.AcademicYear,
			Professors:        []*course_management.Professor{},
			Prerequisites:     []*course_management.Prerequisite{},
			Status:            subject.Status,
			CurriculumName:    subject.CurriculumName,
			DegreeLevel:       subject.DegreeLevel,
			TeachingHours:     subject.TeachingHours,
		})
	}
	// Return
	return &course_management.CourseList{
		Course: courseList,
	}, nil
}

func (s *Server) GetCourse(ctx context.Context, req *course_management.CourseId) (*course_management.CourseItem, error) {
	log.Println("GetCourse", req.GetCourseId())
	// Get course in database
	var course models.Course
	db.First(&course, req.GetCourseId())
	// Convert to protobuf
	courseItem := &course_management.CourseItem{
		CourseId:          course.CourseId,
		CourseName:        course.CourseName,
		CourseDescription: course.CourseDescription,
		FacultyDepartment: course.FacultyDepartment,
		AcademicTerm:      course.AcademicTerm,
		AcademicYear:      course.AcademicYear,
		Professors:        []*course_management.Professor{},
		Prerequisites:     []*course_management.Prerequisite{},
		Status:            course.Status,
		CurriculumName:    course.CurriculumName,
		DegreeLevel:       course.DegreeLevel,
		TeachingHours:     course.TeachingHours,
	}
	// Return
	return courseItem, nil
}

func (s *Server) AddNewCourse(ctx context.Context, req *course_management.CourseItem) (*course_management.CourseItem, error) {
	log.Print("AddNewCourse")
	course := models.NewCourse(
		req.GetCourseId(),
		req.GetCourseName(),
		req.GetCourseDescription(),
		req.GetFacultyDepartment(),
		req.GetAcademicTerm(),
		req.GetAcademicYear(),
		"test",
		"test",
		req.GetStatus(),
		req.GetCurriculumName(),
		req.GetDegreeLevel(),
		req.GetTeachingHours(),
	)
	// Add new course to database
	db.Create(&course)
	// Return
	return &course_management.CourseItem{
		CourseId:          course.CourseId,
		CourseName:        course.CourseName,
		CourseDescription: course.CourseDescription,
		FacultyDepartment: course.FacultyDepartment,
		AcademicTerm:      course.AcademicTerm,
		AcademicYear:      course.AcademicYear,
		Professors:        []*course_management.Professor{},
		Prerequisites:     []*course_management.Prerequisite{},
		Status:            course.Status,
		CurriculumName:    course.CurriculumName,
		DegreeLevel:       course.DegreeLevel,
		TeachingHours:     course.TeachingHours,
	}, nil
}

func (s *Server) UpdateCourseDetail(ctx context.Context, req *course_management.CourseItem) (*course_management.CourseItem, error) {
	log.Println("UpdateCourseDetail")
	// Get course in database
	var course models.Course
	db.First(&course, req.GetCourseId())
	// Update course
	course.CourseName = req.GetCourseName()
	course.CourseDescription = req.GetCourseDescription()
	course.FacultyDepartment = req.GetFacultyDepartment()
	course.AcademicTerm = req.GetAcademicTerm()
	course.AcademicYear = req.GetAcademicYear()
	course.Status = req.GetStatus()
	course.CurriculumName = req.GetCurriculumName()
	course.DegreeLevel = req.GetDegreeLevel()
	course.TeachingHours = req.GetTeachingHours()
	db.Save(&course)
	// Return
	return &course_management.CourseItem{
		CourseId:          course.CourseId,
		CourseName:        course.CourseName,
		CourseDescription: course.CourseDescription,
		FacultyDepartment: course.FacultyDepartment,
		AcademicTerm:      course.AcademicTerm,
		AcademicYear:      course.AcademicYear,
		Professors:        []*course_management.Professor{},
		Prerequisites:     []*course_management.Prerequisite{},
		Status:            course.Status,
		CurriculumName:    course.CurriculumName,
		DegreeLevel:       course.DegreeLevel,
		TeachingHours:     course.TeachingHours,
	}, nil
}

func (s *Server) DeleteCourse(ctx context.Context, req *course_management.CourseId) (*course_management.Empty, error) {
	log.Println("DeleteCourse", req.GetCourseId())
	// Get course in database
	var course models.Course
	db.First(&course, req.GetCourseId())
	// Delete course
	db.Delete(&course)
	// Return
	return &course_management.Empty{}, nil
}
