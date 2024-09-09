package models

import (
	"gorm.io/gorm"
)

type Idea struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Add more models here
