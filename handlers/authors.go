package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"online-bookstore-api/models"
	"strings"
	"time"
)

// CreateAuthor handles POST /authors with context support
func (h *Handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if checkContext(ctx, w) {
		return
	}

	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate required fields
	if author.FirstName == "" || author.LastName == "" {
		respondWithError(w, http.StatusBadRequest, "First name and last name are required")
		return
	}

	if checkContext(ctx, w) {
		return
	}

	createdAuthor, err := h.AuthorStore.CreateAuthor(author)
	if err != nil {
		log.Printf("Error creating author: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create author")
		return
	}

	log.Printf("Author created: ID=%d, Name=%s %s", createdAuthor.ID, createdAuthor.FirstName, createdAuthor.LastName)
	respondWithJSON(w, http.StatusCreated, createdAuthor)
}

// GetAuthor handles GET /authors/{id} with context support
func (h *Handler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx := r.Context()
	if checkContext(ctx, w) {
		return
	}

	id, err := extractID(r.URL.Path, "/authors/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid author ID")
		return
	}

	author, err := h.AuthorStore.GetAuthor(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "Author not found")
		} else {
			log.Printf("Error getting author: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to retrieve author")
		}
		return
	}

	respondWithJSON(w, http.StatusOK, author)
}

// UpdateAuthor handles PUT /authors/{id} with context support
func (h *Handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if checkContext(ctx, w) {
		return
	}

	id, err := extractID(r.URL.Path, "/authors/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid author ID")
		return
	}

	var author models.Author
	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if checkContext(ctx, w) {
		return
	}

	updatedAuthor, err := h.AuthorStore.UpdateAuthor(id, author)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "Author not found")
		} else {
			log.Printf("Error updating author: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to update author")
		}
		return
	}

	log.Printf("Author updated: ID=%d", updatedAuthor.ID)
	respondWithJSON(w, http.StatusOK, updatedAuthor)
}

// DeleteAuthor handles DELETE /authors/{id} with context support
func (h *Handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if checkContext(ctx, w) {
		return
	}

	id, err := extractID(r.URL.Path, "/authors/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid author ID")
		return
	}

	if checkContext(ctx, w) {
		return
	}

	if err := h.AuthorStore.DeleteAuthor(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "Author not found")
		} else {
			log.Printf("Error deleting author: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to delete author")
		}
		return
	}

	log.Printf("Author deleted: ID=%d", id)
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Author deleted successfully"})
}

// GetAllAuthors handles GET /authors with context support
func (h *Handler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx := r.Context()
	if checkContext(ctx, w) {
		return
	}

	authors, err := h.AuthorStore.GetAllAuthors()
	if err != nil {
		log.Printf("Error getting all authors: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve authors")
		return
	}

	respondWithJSON(w, http.StatusOK, authors)
}
