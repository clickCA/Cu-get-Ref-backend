package models

import (
	"time"
)

type ReviewRequest struct {
	Id       string `json:"id"`
	Reviewer string `json:"reviewer"`
	Message  string `json:"message"`
	Rating   int    `json:"rating"`
}

type Review struct {
	ID       string
	REVIEWER string
	DATE     time.Time
	MESSAGE  string
	RATING   int
}

func NewReview(id, reviewer, message string, rating int, date time.Time) *Review {
	d := new(Review)
	d.ID = id
	d.REVIEWER = reviewer
	d.DATE = date
	d.MESSAGE = message
	d.RATING = rating

	return d
}
