package services

import (
	"context"
	course_management "course-management-service/coursemanagement"
)

type Server struct {
	course_management.UnimplementedCourseManagementServiceServer
}

func (s *Server) GetAllSubject(context.Context, *course_management.Empty) (*course_management.SubjectList, error) {
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