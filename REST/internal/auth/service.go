package auth

import (
	"REST/internal/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

// New creates a new authentication service instance.
func New(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

// Register a new user.
func (s *AuthService) Register(email, password, userType string) (*gorm.Model, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var base gorm.Model
	switch userType {
	case "student":
		student := models.Student{
			Email:    email,
			Password: string(hashedPassword),
		}
		if err := s.db.Create(&student).Error; err != nil {
			return nil, err
		}
		base = student.Model

	case "professor":
		professor := models.Professor{
			Email:    email,
			Password: string(hashedPassword),
		}
		if err := s.db.Create(&professor).Error; err != nil {
			return nil, err
		}
		base = professor.Model

	default:
		return nil, errors.New("invalid user type")
	}

	return &base, nil
}

// Login checks the credentials and returns the user if they're correct.
func (s *AuthService) Login(email, password, userType string) (*gorm.Model, error) {
	var userPassword string
	var base gorm.Model

	switch userType {
	case "student":
		var student models.Student
		if err := s.db.Where("email = ?", email).First(&student).Error; err != nil {
			return nil, err
		}
		userPassword = student.Password
		base = student.Model
	case "professor":
		var professor models.Professor
		if err := s.db.Where("email = ?", email).First(&professor).Error; err != nil {
			return nil, err
		}
		userPassword = professor.Password
		base = professor.Model
	default:
		return nil, errors.New("invalid user type")
	}

	// Check the password
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	if err != nil {
		return nil, errors.New("invalid login credentials")
	}

	return &base, nil
}

//
