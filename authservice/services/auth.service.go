package services

import (
	"authservice/jwt"
	"authservice/models"
	"errors"
	"fmt"
	"time"

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
	// we make a JWT Token here with signing method of ES256 and claims.
	// claims are attributes.
	// Aud - audience
	// Iss - issuer
	// Exp - expiration of the Token
	claimsMap := jwt.ClaimsMap{
		Aud: "cugetref.org",
		Iss: "cugetref.org",
		Exp: fmt.Sprint(time.Now().Add(time.Minute * 1).Unix()),
	}

	secret := jwt.GetSecret()
	if secret == "" {
		return "", errors.New("empty JWT secret")
	}

	header := "HS256"
	tokenString, err := jwt.GenerateToken(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
