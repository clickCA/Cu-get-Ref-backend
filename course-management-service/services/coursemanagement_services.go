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

var dsn = "root:password@localhost:3307/db?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

type Subject struct {
	SubjectId string `gorm:"primaryKey"`
	SubjectName string
	CourseDescription string
	FacultyDepartment string
	AcademicTerm string
	AcademicYear int32
	Professors []*string
	Prerequisites []*string
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
	log.Println("GetAllSubject")
	return &course_management.SubjectList{
		Subject: []*course_management.SubjectItem{
			{
				SubjectId: "2110415",
				SubjectName: "Software Architecture",
				CourseDescription: "1. Introduction to Software Architecture",
				FacultyDepartment: "Computer Engineering",
				AcademicTerm: "1",
				AcademicYear: 2563,
				Professors: []*course_management.Professor{
					{
						ProfessorName: "Dr. Somchai",
					},
				},
				Prerequisites: []*course_management.Prerequisite{
					{
						SubjectId: "2110211",
					},
				},
				Status: "open",
				CurriculumName: "Computer Engineering",
				DegreeLevel: "bachelor",
				TeachingHours: 3,
			},
		},
		}, nil
}

func (s *Server) GetSubject(ctx context.Context,req *course_management.SubjectId) (*course_management.SubjectItem, error) {
	log.Println("GetSubject")
	log.Println("Getting ",req.GetSubjectId())
	return &course_management.SubjectItem{
		SubjectId: "2110415",
		SubjectName: "Software Architecture",
		CourseDescription: "1. Introduction to Software Architecture",
		FacultyDepartment: "Computer Engineering",
		AcademicTerm: "1",
		AcademicYear: 2563,
		Professors: []*course_management.Professor{
			{
				ProfessorName: "Dr. Somchai",
			},
		},
		Prerequisites: []*course_management.Prerequisite{
			{
				SubjectId: "2110211",
			},
		},
		Status: "open",
		CurriculumName: "Computer Engineering",
		DegreeLevel: "bachelor",
		TeachingHours: 3,
	}, nil
}

func (s *Server) AddNewSubject(ctx context.Context,req *course_management.SubjectItem) (*course_management.SubjectItem, error) {
	log.Println("AddNewSubject")
	log.Println("Adding ",req.String())
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
	return &course_management.Empty{}, nil
}