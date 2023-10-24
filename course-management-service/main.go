package main

import (
	"context"
	course_management_service "course-management-service/coursemanagement"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	course_management_service.UnimplementedCourseManagementServiceServer
}

func (s *server) GetAllSubject(context.Context, *course_management_service.Empty) (*course_management_service.SubjectList, error) {
	return &course_management_service.SubjectList{
		Subject: []*course_management_service.SubjectItem{
			{
				Id: "1",
				Name: "Math",
				Price: 3,
			},
			{
				Id: "2",
				Name: "English",
				Price: 2,
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
	course_management_service.RegisterCourseManagementServiceServer(s, &server{})

	// Start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}