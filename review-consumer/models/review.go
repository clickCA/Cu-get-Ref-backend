package models

import (
	"time"
)

type Review struct {
	COURSE_ID string
	REVIEWER  string
	DATE      time.Time
	MESSAGE   string
	RATING    int
}

type ReviewReq struct {
	COURSE_ID string
	METHOD    string
	REVIEWER  string
	DATE      time.Time
	MESSAGE   string
	RATING    int
}

func NewReview(id, reviewer, message string, rating int, date time.Time) *Review {
	d := new(Review)
	d.COURSE_ID = id
	d.REVIEWER = reviewer
	d.DATE = date
	d.MESSAGE = message
	d.RATING = rating

	return d
}
