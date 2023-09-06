package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ReviewerName string
	ReceiverID int
	Title string
	Text string
	Rating int
}