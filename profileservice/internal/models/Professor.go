package models

import "gorm.io/gorm"

type Professor struct {
	gorm.Model
	Name            string   `gorm:"notNull" json:"name"`
	Email           string   `gorm:"type:varchar(255);uniqueIndex;notNull" json:"email"`
	Password        string   `gorm:"notNull" json:"password"`
	SubjectsTaught  []Course `gorm:"many2many:professor_subjects;" json:"subjects_taught"`
	ReviewsReceived []Review `gorm:"foreignKey:ProfessorID;references:ID" json:"reviews_received"`
}
