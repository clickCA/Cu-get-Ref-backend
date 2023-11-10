package services

import (
	"authservice/config"
	"authservice/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
func (s *AuthService) Register(email, password string, role models.Role) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
		Role:         role,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Login checks the credentials and returns the user if they're correct.
func (s *AuthService) Login(email, password string, role models.Role) (*models.User, error) {
	var user models.User

	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	// Check the password
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid login credentials")
	}

	return &user, nil
}

func (s *AuthService) GetSignedToken() (string, error) {
	// Your shared secret
	secret, key := config.GetJWTSecret()
	// Create the Claims
	claims := jwt.MapClaims{
		"iss": key,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(30 * time.Minute).Unix(),
		"aud": "cugetref.org",
	}

	// Create the token with HS256 algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your shared secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
