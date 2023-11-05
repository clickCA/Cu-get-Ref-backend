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

func (s *Server) GetAllSubject(ctx context.Context, req *course_management.Empty) (*course_management.SubjectList, error) {
	log.Println("[Maintenance] GetAllSubject")
	return &course_management.SubjectList{}, nil
}

func (s *Server) GetSubject(ctx context.Context, req *course_management.SubjectId) (*course_management.SubjectItem, error) {
	log.Println("GetSubject", req.GetSubjectId())
	var subject models.Subject
	db.First(&subject, "subject_id = ?", req.GetSubjectId())
	return &course_management.SubjectItem{
		SubjectId:         subject.SubjectId,
		SubjectName:       subject.SubjectName,
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
	}, nil
}

func (s *Server) AddNewSubject(ctx context.Context, req *course_management.SubjectItem) (*course_management.SubjectItem, error) {
	log.Println("AddNewSubject", req.GetSubjectId())
	subject := models.Subject{
		SubjectId:         req.GetSubjectId(),
		SubjectName:       req.GetSubjectName(),
		CourseDescription: req.GetCourseDescription(),
		FacultyDepartment: req.GetFacultyDepartment(),
		AcademicTerm:      req.GetAcademicTerm(),
		AcademicYear:      req.GetAcademicYear(),
		Professors:        "test",
		Prerequisites:     "test",
		Status:            req.GetStatus(),
		CurriculumName:    req.GetCurriculumName(),
		DegreeLevel:       req.GetDegreeLevel(),
		TeachingHours:     req.GetTeachingHours(),
	}
	db.Create(&subject)
	return &course_management.SubjectItem{
		SubjectId:         req.GetSubjectId(),
		SubjectName:       req.GetSubjectName(),
		CourseDescription: req.GetCourseDescription(),
		FacultyDepartment: req.GetFacultyDepartment(),
		AcademicTerm:      req.GetAcademicTerm(),
		AcademicYear:      req.GetAcademicYear(),
		Professors:        req.GetProfessors(),
		Prerequisites:     req.GetPrerequisites(),
		Status:            req.GetStatus(),
		CurriculumName:    req.GetCurriculumName(),
		DegreeLevel:       req.GetDegreeLevel(),
		TeachingHours:     req.GetTeachingHours(),
	}, nil
}

func (s *Server) UpdateSubjectDetail(ctx context.Context, req *course_management.SubjectItem) (*course_management.SubjectItem, error) {
	log.Println("UpdateSubjectDetail")
	log.Println("Updating ", req.String())
	subject := models.Subject{
		SubjectId:         req.GetSubjectId(),
		SubjectName:       req.GetSubjectName(),
		CourseDescription: req.GetCourseDescription(),
		FacultyDepartment: req.GetFacultyDepartment(),
		AcademicTerm:      req.GetAcademicTerm(),
		AcademicYear:      req.GetAcademicYear(),
		Professors:        "test",
		Prerequisites:     "test",
		Status:            req.GetStatus(),
		CurriculumName:    req.GetCurriculumName(),
		DegreeLevel:       req.GetDegreeLevel(),
		TeachingHours:     req.GetTeachingHours(),
	}
	db.Save(&subject)
	return &course_management.SubjectItem{
		SubjectId:         req.GetSubjectId(),
		SubjectName:       req.GetSubjectName(),
		CourseDescription: req.GetCourseDescription(),
		FacultyDepartment: req.GetFacultyDepartment(),
		AcademicTerm:      req.GetAcademicTerm(),
		AcademicYear:      req.GetAcademicYear(),
		Professors:        req.GetProfessors(),
		Prerequisites:     req.GetPrerequisites(),
		Status:            req.GetStatus(),
		CurriculumName:    req.GetCurriculumName(),
		DegreeLevel:       req.GetDegreeLevel(),
		TeachingHours:     req.GetTeachingHours(),
	}, nil
}

func (s *Server) DeleteSubject(ctx context.Context, req *course_management.SubjectId) (*course_management.Empty, error) {
	log.Println("DeleteSubject")
	log.Println("Deleting ", req.GetSubjectId())
	db.Delete(&models.Subject{}, req.GetSubjectId())
	return &course_management.Empty{}, nil
}
