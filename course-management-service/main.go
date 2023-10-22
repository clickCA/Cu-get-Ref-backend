package main

import (
	"log"
)

func main() {
	// Create a new gRPC server
	s := grpc.NewServer()

	// Attach the CourseManagement service to the server
	pb.RegisterCourseManagementServer(s, &server{})

	// Start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}