// based on the email id provided, finds the user object
// can be seen as the main constructor to start validation
package services

import (
	"authservice/models"

	"gorm.io/gorm"
)

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, email, username, passwordHash, fullname string, role int) (*models.User, error) {
	user := &models.User{
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
	}
	result := db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetUserByEmail retrieves a user from the database by email
func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	user := &models.User{}
	result := db.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// UpdateUser updates a user in the database
func UpdateUser(db *gorm.DB, user *models.User) error {
	result := db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser deletes a user from the database
func DeleteUser(db *gorm.DB, user *models.User) error {
	result := db.Delete(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
