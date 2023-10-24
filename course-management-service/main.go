package main

import (
	"context"
	course_management "course-management-service/coursemanagement"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	course_management.UnimplementedCourseManagementServiceServer
}

func (s *server) GetAllSubject(context.Context, *course_management.Empty) (*course_management.SubjectList, error) {
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

func main() {
	lis, err := net.Listen("tcp","localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Create a new gRPC server
	s := grpc.NewServer()

	// Attach the CourseManagement service to the server
	course_management.RegisterCourseManagementServiceServer(s, &server{})

	// Start the server
	log.Println("Server is listening on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}