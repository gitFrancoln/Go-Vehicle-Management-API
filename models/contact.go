package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone"`
	Message string `json:"message" binding:"required"`
}
