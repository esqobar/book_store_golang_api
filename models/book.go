package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title    string `json:"title" binding:"required"`
	ISBN     string `json:"isbn" binding:"required"`
	AuthorID uint   `json:"author_id" binding:"required"`
	Author   Author `json:"author,omitempty" binding:"-"`
}

type CreateBookInput struct {
	Title    string `json:"title" binding:"required"`
	ISBN     string `json:"isbn" binding:"required"`
	AuthorID uint   `json:"author_id" binding:"required"`
}
