package services

import (
	"ideahive/backend/internal/database"
	"ideahive/backend/internal/models"
)

type Services struct {
	db *database.Database
}

func New(db *database.Database) *Services {
	return &Services{db: db}
}

func (s *Services) CreateIdea(title, description string) (*models.Idea, error) {
	dbIdea := s.db.CreateIdea(title, description)
	return &models.Idea{
		ID:          dbIdea.ID,
		Title:       dbIdea.Title,
		Description: dbIdea.Description,
	}, nil
}
