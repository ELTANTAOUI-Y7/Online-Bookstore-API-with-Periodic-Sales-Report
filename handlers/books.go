package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"online-bookstore-api/models"
	"strconv"
	"strings"
)

// CreateBook handles POST /books
func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate required fields
	if book.Title == "" {
		respondWithError(w, http.StatusBadRequest, "Title is required")
		return
	}

	createdBook, err := h.BookStore.CreateBook(book)
	if err != nil {
		log.Printf("Error creating book: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create book")
		return
	}

	log.Printf("Book created: ID=%d, Title=%s", createdBook.ID, createdBook.Title)
	respondWithJSON(w, http.StatusCreated, createdBook)
}

// GetBook handles GET /books/{id}
func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := extractID(r.URL.Path, "/books/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := h.BookStore.GetBook(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}

// UpdateBook handles PUT /books/{id}
func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := extractID(r.URL.Path, "/books/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updatedBook, err := h.BookStore.UpdateBook(id, book)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "Book not found")
		} else {
			log.Printf("Error updating book: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to update book")
		}
		return
	}

	log.Printf("Book updated: ID=%d", updatedBook.ID)
	respondWithJSON(w, http.StatusOK, updatedBook)
}

// DeleteBook handles DELETE /books/{id}
func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := extractID(r.URL.Path, "/books/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	if err := h.BookStore.DeleteBook(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "Book not found")
		} else {
			log.Printf("Error deleting book: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to delete book")
		}
		return
	}

	log.Printf("Book deleted: ID=%d", id)
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Book deleted successfully"})
}

// SearchBooks handles GET /books with query parameters
func (h *Handler) SearchBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse query parameters
	criteria := models.SearchCriteria{
		Title: r.URL.Query().Get("title"),
		Genre: r.URL.Query().Get("genre"),
	}

	// Parse author_id
	if authorIDStr := r.URL.Query().Get("author_id"); authorIDStr != "" {
		if authorID, err := strconv.Atoi(authorIDStr); err == nil {
			criteria.AuthorID = authorID
		}
	}

	// Parse min_price
	if minPriceStr := r.URL.Query().Get("min_price"); minPriceStr != "" {
		if minPrice, err := strconv.ParseFloat(minPriceStr, 64); err == nil {
			criteria.MinPrice = minPrice
		}
	}

	// Parse max_price
	if maxPriceStr := r.URL.Query().Get("max_price"); maxPriceStr != "" {
		if maxPrice, err := strconv.ParseFloat(maxPriceStr, 64); err == nil {
			criteria.MaxPrice = maxPrice
		}
	}

	// If no search criteria provided, return all books
	if criteria.Title == "" && criteria.AuthorID == 0 && criteria.Genre == "" &&
		criteria.MinPrice == 0 && criteria.MaxPrice == 0 {
		books, err := h.BookStore.GetAllBooks()
		if err != nil {
			log.Printf("Error getting all books: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to retrieve books")
			return
		}
		respondWithJSON(w, http.StatusOK, books)
		return
	}

	books, err := h.BookStore.SearchBooks(criteria)
	if err != nil {
		log.Printf("Error searching books: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to search books")
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}
