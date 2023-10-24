package models

import (
	"time"
)

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
