package CRUD
import (
	"REST/internal/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)
type CRUD struct{
	db *gorm.DB
}
func New(db *gorm.DB) *Create {
	return &Create{db: db}
}
// Register a new user.
func (s *CRUD) Create(courseName,courseCode) (*models.Course, error) {
	var course models.Course
	course:= models.Course{
		CourseName: courseName,
		CourseCode: courseCode
	}
	if err := s.db.Create(&course).Error; err != nil {
			return nil, err
	}
	return &course,nil
}
func (s *CRUD) Update(courseCode,newCourseName)(*models.Course, error) {
	var course models.Course
	if err := s.db.Where("courseCode = ?", courseCode).First(&course).Error; err != nil {
		return nil, err
	}
	code = course.CourseCode
	s.db.Update()
	// Very not understand please do it
}