package main

import (
	"course-management-service/config"
	course_management "course-management-service/coursemanagement"
	"course-management-service/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Create a TCP listener on port
	port := config.EnvServerPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	s := grpc.NewServer()
	services := &services.Server{}

	// Attach the CourseManagement service to the server
	course_management.RegisterCourseManagementServiceServer(s, services)

	// Start the server
	log.Printf("Starting server on port %v...", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
