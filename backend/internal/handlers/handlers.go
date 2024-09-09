package handlers

import (
	"encoding/json"
	"idea-repository-backend/internal/services"
	"net/http"
)

type Handlers struct {
	services *services.Services
}

func New(s *services.Services) *Handlers {
	return &Handlers{services: s}
}

func (h *Handlers) CreateIdeaHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateIdeaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idea, err := h.services.CreateIdea(req.Title, req.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := IdeaResponse{
		ID:          idea.ID,
		Title:       idea.Title,
		Description: idea.Description,
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type CreateIdeaRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type IdeaResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
