package database

import (
	"idea-repository-backend/internal/models"

	"gorm.io/gorm"
)

type Database struct {
	GormDB *gorm.DB
}

func New(gormDB *gorm.DB) *Database {
	return &Database{
		GormDB: gormDB,
	}
}

func (db *Database) CreateIdea(title, description string) *models.Idea {
	idea := &models.Idea{
		Title:       title,
		Description: description,
	}
	db.GormDB.Create(idea)
	return idea
}
