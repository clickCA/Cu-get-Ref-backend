package services

import (
	"context"
	course_management "course-management-service/coursemanagement"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	course_management.UnimplementedCourseManagementServiceServer
}

var dsn = "root:password@tcp(localhost:3307)/db?charset=utf8mb4&parseTime=True"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type Subject struct {
	SubjectId string `gorm:"primaryKey"`
	SubjectName string
	CourseDescription string
	FacultyDepartment string
	AcademicTerm string
	AcademicYear int32
	Professors string
	Prerequisites string
	Status string
	CurriculumName string
	DegreeLevel string
	TeachingHours int32
}

// type CourseManagementServiceClient interface {
// 	GetAllSubject(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SubjectList, error)
// 	GetSubject(ctx context.Context, in *SubjectId, opts ...grpc.CallOption) (*SubjectItem, error)
// 	AddNewSubject(ctx context.Context, in *SubjectItem, opts ...grpc.CallOption) (*SubjectItem, error)
// 	UpdateSubjectDetail(ctx context.Context, in *SubjectItem, opts ...grpc.CallOption) (*SubjectItem, error)
// 	DeleteSubject(ctx context.Context, in *SubjectId, opts ...grpc.CallOption) (*Empty, error)
// }
func (s *Server) GetAllSubject(ctx context.Context,req *course_management.Empty) (*course_management.SubjectList, error) {
	log.Println("[Maintenance] GetAllSubject")
	return &course_management.SubjectList{}, nil
}

func (s *Server) GetSubject(ctx context.Context,req *course_management.SubjectId) (*course_management.SubjectItem, error) {
	log.Println("GetSubject", req.GetSubjectId())
	var subject Subject
	db.First(&subject, "subject_id = ?", req.GetSubjectId())
	return &course_management.SubjectItem{
		SubjectId: subject.SubjectId,
		SubjectName: subject.SubjectName,
		CourseDescription: subject.CourseDescription,
		FacultyDepartment: subject.FacultyDepartment,
		AcademicTerm: subject.AcademicTerm,
		AcademicYear: subject.AcademicYear,
		Professors: []*course_management.Professor{},
		Prerequisites: []*course_management.Prerequisite{},
		Status: subject.Status,
		CurriculumName: subject.CurriculumName,
		DegreeLevel: subject.DegreeLevel,
		TeachingHours: subject.TeachingHours,
	}, nil
}

func (s *Server) AddNewSubject(ctx context.Context,req *course_management.SubjectItem) (*course_management.SubjectItem, error) {
	log.Println("AddNewSubject", req.GetSubjectId())
	subject := Subject{
		SubjectId: req.GetSubjectId(),
		SubjectName: req.GetSubjectName(),
		CourseDescription: req.GetCourseDescription(),
		FacultyDepartment: req.GetFacultyDepartment(),
		AcademicTerm: req.GetAcademicTerm(),
		AcademicYear: req.GetAcademicYear(),
		Professors: "test",
		Prerequisites: "test",
		Status: req.GetStatus(),
		CurriculumName: req.GetCurriculumName(),
		DegreeLevel: req.GetDegreeLevel(),
		TeachingHours: req.GetTeachingHours(),
	}
	db.Create(&subject)
	return &course_management.SubjectItem{
		SubjectId: req.GetSubjectId(),
		SubjectName: req.GetSubjectName(),
		CourseDescription: req.GetCourseDescription(),
		FacultyDepartment: req.GetFacultyDepartment(),
		AcademicTerm: req.GetAcademicTerm(),
		AcademicYear: req.GetAcademicYear(),
		Professors: req.GetProfessors(),
		Prerequisites: req.GetPrerequisites(),
		Status: req.GetStatus(),
		CurriculumName: req.GetCurriculumName(),
		DegreeLevel: req.GetDegreeLevel(),
		TeachingHours: req.GetTeachingHours(),
	}, nil
}

func (s *Server) UpdateSubjectDetail(ctx context.Context,req *course_management.SubjectItem) (*course_management.SubjectItem, error) {
	log.Println("UpdateSubjectDetail")
	log.Println("Updating ",req.String())
	subject := Subject{
		SubjectId: req.GetSubjectId(),
		SubjectName: req.GetSubjectName(),
		CourseDescription: req.GetCourseDescription(),
		FacultyDepartment: req.GetFacultyDepartment(),
		AcademicTerm: req.GetAcademicTerm(),
		AcademicYear: req.GetAcademicYear(),
		Professors: "test",
		Prerequisites: "test",
		Status: req.GetStatus(),
		CurriculumName: req.GetCurriculumName(),
		DegreeLevel: req.GetDegreeLevel(),
		TeachingHours: req.GetTeachingHours(),
	}
	db.Save(&subject)
	return &course_management.SubjectItem{
		SubjectId: req.GetSubjectId(),
		SubjectName: req.GetSubjectName(),
		CourseDescription: req.GetCourseDescription(),
		FacultyDepartment: req.GetFacultyDepartment(),
		AcademicTerm: req.GetAcademicTerm(),
		AcademicYear: req.GetAcademicYear(),
		Professors: req.GetProfessors(),
		Prerequisites: req.GetPrerequisites(),
		Status: req.GetStatus(),
		CurriculumName: req.GetCurriculumName(),
		DegreeLevel: req.GetDegreeLevel(),
		TeachingHours: req.GetTeachingHours(),
	}, nil
}

func (s *Server) DeleteSubject(ctx context.Context,req *course_management.SubjectId) (*course_management.Empty, error) {
	log.Println("DeleteSubject")
	log.Println("Deleting ",req.GetSubjectId())
	db.Delete(&Subject{}, req.GetSubjectId())
	return &course_management.Empty{}, nil
}