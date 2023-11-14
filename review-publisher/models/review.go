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

type ReviewRequest struct {
	Id       string `json:"id"`
	Reviewer string `json:"reviewer"`
	Message  string `json:"message"`
	Rating   int    `json:"rating"`
}

type ReviewReq struct {
	COURSE_ID string
	METHOD    string
	REVIEWER  string
	DATE      time.Time
	MESSAGE   string
	RATING    int
}

func NewReviewReq(id, method, reviewer, message string, rating int, date time.Time) *ReviewReq {
	d := new(ReviewReq)
	d.COURSE_ID = id
	d.METHOD = method
	d.REVIEWER = reviewer
	d.DATE = date
	d.MESSAGE = message
	d.RATING = rating

	return d
}
