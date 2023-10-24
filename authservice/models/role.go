package models

type Role int

const (
	Admin Role = iota
	Student
	Professor
)
