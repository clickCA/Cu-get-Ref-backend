package services

import (
	"context"
	"course-management-service/config"
	course_management "course-management-service/coursemanagement"
	"course-management-service/models"
	"log"
	"strings"

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
			Professors:        convertStringToProfessors(subject.Professors),
			Prerequisites:     convertStringToPrerequisites(subject.Prerequisites),
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
		Professors:        convertStringToProfessors(course.Professors),
		Prerequisites:     convertStringToPrerequisites(course.Prerequisites),
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
	professors_string := convertProfessorsToString(req.GetProfessors())
	prerequisites_string := convertPrerequisitesToString(req.GetPrerequisites())
	// Create new course
	course, err := models.NewCourse(
		req.GetCourseId(),
		req.GetCourseName(),
		req.GetCourseDescription(),
		req.GetFacultyDepartment(),
		req.GetAcademicTerm(),
		req.GetAcademicYear(),
		professors_string,
		prerequisites_string,
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
	// Map to string
	professors_string := convertProfessorsToString(req.GetProfessors())
	prerequisites_string := convertPrerequisitesToString(req.GetPrerequisites())
	// Create new course
	err := models.UpdateCourse(
		&course,
		req.GetCourseId(),
		req.GetCourseName(),
		req.GetCourseDescription(),
		req.GetFacultyDepartment(),
		req.GetAcademicTerm(),
		req.GetAcademicYear(),
		professors_string,
		prerequisites_string,
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
	// Update course
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
		Professors:        convertStringToProfessors(course.Professors),
		Prerequisites:     convertStringToPrerequisites(course.Prerequisites),
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

func convertProfessorsToString(professors []*course_management.Professor) string {
	var professors_string string
	for _, professor := range professors {
		// if the last element skip ||
		if professor == professors[len(professors)-1] {
			professors_string += professor.GetProfessorName()
			continue
		}
		professors_string += professor.GetProfessorName() + "||"
	}
	return professors_string
}

func convertPrerequisitesToString(prerequisites []*course_management.Prerequisite) string {
	var prerequisites_string string
	for _, prerequisite := range prerequisites {
		// if the last element skip ||
		if prerequisite == prerequisites[len(prerequisites)-1] {
			prerequisites_string += prerequisite.GetCourseId()
			continue
		}
		prerequisites_string += prerequisite.GetCourseId() + "||"
	}
	return prerequisites_string
}

func convertStringToProfessors(professors string) []*course_management.Professor {
	var professors_list []*course_management.Professor
	professors_array := strings.Split(professors, "||")
	for _, professor := range professors_array {
		// Skip empty string
		if professor == "" {
			continue
		}
		professors_list = append(professors_list, &course_management.Professor{
			ProfessorName: professor,
		})
	}
	return professors_list
}

func convertStringToPrerequisites(prerequisites string) []*course_management.Prerequisite {
	var prerequisites_list []*course_management.Prerequisite
	prerequisites_array := strings.Split(prerequisites, "||")
	for _, prerequisite := range prerequisites_array {
		// Skip empty string
		if prerequisite == "" {
			continue
		}
		prerequisites_list = append(prerequisites_list, &course_management.Prerequisite{
			CourseId: prerequisite,
		})
	}
	return prerequisites_list
}
