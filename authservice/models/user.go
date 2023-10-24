package models

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string
	Role         int
}
