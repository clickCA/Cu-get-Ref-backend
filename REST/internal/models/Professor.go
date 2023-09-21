package models

type Professor struct {
	Base
	Name            string   `gorm:"notNull" json:"name"`
	Email           string   `gorm:"uniqueIndex;notNull" json:"email"`
	Password        string   `gorm:"notNull" json:"password"`
	SubjectsTaught  []Course `gorm:"many2many:professor_subjects;" json:"subjects_taught"`
	ReviewsReceived []Review `gorm:"foreignKey:ProfessorID;references:ID" json:"reviews_received"`
}
