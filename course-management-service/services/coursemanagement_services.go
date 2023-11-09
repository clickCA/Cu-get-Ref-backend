package services

import (
	"context"
	"course-management-service/config"
	course_management "course-management-service/coursemanagement"
	"course-management-service/models"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	course_management.UnimplementedCourseManagementServiceServer
}

var db = config.ConnectDB()

func (s *Server) GetAllCourses(ctx context.Context, req *course_management.Empty) (*course_management.CourseList, error) {
	log.Println("GetAllCourses")
	// Get all course in database
	var courses []models.Course
	result := db.Find(&courses)
	if result.Error != nil {
		err := status.Errorf(
			codes.Internal,
			"Could not find courses",
		)
		return nil, err
	}
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
	result := db.First(&course, req.GetCourseId())
	if result.Error != nil {
		err := status.Errorf(
			codes.NotFound,
			"Could not find course with ID: %s",
			req.GetCourseId(),
		)
		return nil, err
	}
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
	// Create new course
	// Map professors
	var professors []models.Professor
	for _, professor := range req.GetProfessors() {
		professors = append(professors, models.Professor{
			ProfessorName: professor.GetProfessorName(),
		})
	}
	// Map prerequisites
	var prerequisites []models.Prerequisite
	for _, prerequisite := range req.GetPrerequisites() {
		prerequisites = append(prerequisites, models.Prerequisite{
			PrerequisiteId: prerequisite.GetCourseId(),
		})
	}
	course, err := models.NewCourse(
		req.GetCourseId(),
		req.GetCourseName(),
		req.GetCourseDescription(),
		req.GetFacultyDepartment(),
		req.GetAcademicTerm(),
		req.GetAcademicYear(),
		professors,
		prerequisites,
		req.GetStatus(),
		req.GetCurriculumName(),
		req.GetDegreeLevel(),
		req.GetTeachingHours(),
	)
	if err != nil {
		err = status.Errorf(
			codes.InvalidArgument,
			err.Error(),
		)
		return nil, err
	}
	// Add new course to database
	result := db.Create(&course)
	if result.Error != nil {
		err = status.Errorf(
			codes.Internal,
			"Could not create new course with ID: %s",
			req.GetCourseId(),
		)
		return nil, err
	}
	// Return
	return &course_management.CourseItem{
		CourseId:          course.CourseId,
		CourseName:        course.CourseName,
		CourseDescription: course.CourseDescription,
		FacultyDepartment: course.FacultyDepartment,
		AcademicTerm:      course.AcademicTerm,
		AcademicYear:      course.AcademicYear,
		Professors:        req.GetProfessors(),
		Prerequisites:     req.GetPrerequisites(),
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
	result := db.First(&course, req.GetCourseId())
	if result.Error != nil {
		err := status.Errorf(
			codes.NotFound,
			"Could not find course with ID: %s",
			req.GetCourseId(),
		)
		return nil, err
	}
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
	result = db.Save(&course)
	if result.Error != nil {
		err := status.Errorf(
			codes.Internal,
			"Could not update course with ID: %s",
			req.GetCourseId(),
		)
		return nil, err
	}
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
	result := db.First(&course, req.GetCourseId())
	if result.Error != nil {
		err := status.Errorf(
			codes.NotFound,
			"Could not find course with ID: %s",
			req.GetCourseId(),
		)
		return nil, err
	}
	// Delete course
	result = db.Delete(&course)
	if result.Error != nil {
		err := status.Errorf(
			codes.Internal,
			"Could not delete course with ID: %s",
			req.GetCourseId(),
		)
		return nil, err
	}
	// Return
	return &course_management.Empty{}, nil
}
