package services

import (
	"errors"
	"profileservice/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ProfileService struct {
	db *gorm.DB
}

// New creates a new authentication service instance.
func New(db *gorm.DB) *ProfileService {
	return &ProfileService{db: db}
}

// Create a new student or professor.
func (s *ProfileService) Create(email, password, userType string) (*gorm.Model, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var user gorm.Model
	switch userType {
	case "student":
		student := models.Student{
			Email:    email,
			Password: string(hashedPassword),
		}
		if err := s.db.Create(&student).Error; err != nil {
			return nil, err
		}
		user = student.Model

	case "professor":
		professor := models.Professor{
			Email:    email,
			Password: string(hashedPassword),
		}
		if err := s.db.Create(&professor).Error; err != nil {
			return nil, err
		}
		user = professor.Model

	default:
		return nil, errors.New("invalid user type")
	}

	return &user, nil
}

// Read returns a student or professor by their ID.
func (s *ProfileService) Read(id uint, userType string) (*gorm.Model, error) {
	var user gorm.Model

	switch userType {
	case "student":
		var student models.Student
		if err := s.db.First(&student, id).Error; err != nil {
			return nil, err
		}
		user = student.Model

	case "professor":
		var professor models.Professor
		if err := s.db.First(&professor, id).Error; err != nil {
			return nil, err
		}
		user = professor.Model

	default:
		return nil, errors.New("invalid user type")
	}

	return &user, nil
}

// Update updates the user's information.
func (s *ProfileService) Update(id uint, email, password, userType string) (*gorm.Model, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var user gorm.Model

	switch userType {
	case "student":
		var student models.Student
		if err := s.db.First(&student, id).Error; err != nil {
			return nil, err
		}
		student.Email = email
		student.Password = string(hashedPassword)
		if err := s.db.Save(&student).Error; err != nil {
			return nil, err
		}
		user = student.Model

	case "professor":
		var professor models.Professor
		if err := s.db.First(&professor, id).Error; err != nil {
			return nil, err
		}
		professor.Email = email
		professor.Password = string(hashedPassword)
		if err := s.db.Save(&professor).Error; err != nil {
			return nil, err
		}
		user = professor.Model

	default:
		return nil, errors.New("invalid user type")
	}

	return &user, nil
}

// Delete deletes a student or professor by their ID.
func (s *ProfileService) Delete(id uint, userType string) error {
	switch userType {
	case "student":
		var student models.Student
		if err := s.db.First(&student, id).Error; err != nil {
			return err
		}
		if err := s.db.Delete(&student).Error; err != nil {
			return err
		}

	case "professor":
		var professor models.Professor
		if err := s.db.First(&professor, id).Error; err != nil {
			return err
		}
		if err := s.db.Delete(&professor).Error; err != nil {
			return err
		}

	default:
		return errors.New("invalid user type")
	}

	return nil
}
