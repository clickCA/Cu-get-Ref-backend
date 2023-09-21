package models

type Student struct {
	Base
	StudentID        string   `gorm:"type:varchar(255);uniqueIndex;notNull" json:"student_id"`
	Name             string   `gorm:"notNull" json:"name"`
	Email            string   `gorm:"type:varchar(255);uniqueIndex;notNull" json:"email"`
	Password         string   `gorm:"notNull" json:"password"`
	Year             string   `json:"year"`
	Major            string   `json:"major"`
	ElectiveCourses  []Course `gorm:"many2many:student_electives;" json:"elective_courses"`
	MandatoryCourses []Course `gorm:"many2many:student_mandatories;" json:"mandatory_courses"`
	ReviewsWritten   []Review `gorm:"foreignKey:StudentID;references:ID" json:"reviews_written"`
}
