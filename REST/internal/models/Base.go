package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreatedAt" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdatedAt" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggertype:"string" format:"date-time"`
}
