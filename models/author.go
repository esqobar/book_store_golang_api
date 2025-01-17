package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string `json:"name" binding:"required"`
	Books []Book `json:"books,omitempty"`
}
