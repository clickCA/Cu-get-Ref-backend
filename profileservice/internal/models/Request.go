package models

import "gorm.io/gorm"

type CreateRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType string `json:"userType" binding:"required"`
}
type ReadRequest struct {
	ID       uint   `uri:"id" binding:"required"`
	UserType string `form:"userType" binding:"required"`
}

type UpdateRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	UserType string `form:"userType" binding:"required"`
}

type ResponseModel struct {
	response gorm.Model // Change "models.Base" to the appropriate model type for the user.
}
