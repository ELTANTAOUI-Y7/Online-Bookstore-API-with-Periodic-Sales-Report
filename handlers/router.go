package handlers

import (
	"net/http"
	"strings"
)

// SetupRoutes sets up all API routes
func (h *Handler) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Books routes
	mux.HandleFunc("/books", h.handleBooks)
	mux.HandleFunc("/books/", h.handleBookByID)

	// Authors routes
	mux.HandleFunc("/authors", h.handleAuthors)
	mux.HandleFunc("/authors/", h.handleAuthorByID)

	// Customers routes
	mux.HandleFunc("/customers", h.handleCustomers)
	mux.HandleFunc("/customers/", h.handleCustomerByID)

	// Orders routes
	mux.HandleFunc("/orders", h.handleOrders)
	mux.HandleFunc("/orders/", h.handleOrderByID)

	return mux
}

// handleBooks routes requests to /books
func (h *Handler) handleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateBook(w, r)
	case http.MethodGet:
		h.SearchBooks(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleBookByID routes requests to /books/{id}
func (h *Handler) handleBookByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetBook(w, r)
	case http.MethodPut:
		h.UpdateBook(w, r)
	case http.MethodDelete:
		h.DeleteBook(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleAuthors routes requests to /authors
func (h *Handler) handleAuthors(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateAuthor(w, r)
	case http.MethodGet:
		h.GetAllAuthors(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleAuthorByID routes requests to /authors/{id}
func (h *Handler) handleAuthorByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetAuthor(w, r)
	case http.MethodPut:
		h.UpdateAuthor(w, r)
	case http.MethodDelete:
		h.DeleteAuthor(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleCustomers routes requests to /customers
func (h *Handler) handleCustomers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateCustomer(w, r)
	case http.MethodGet:
		h.GetAllCustomers(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleCustomerByID routes requests to /customers/{id}
func (h *Handler) handleCustomerByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetCustomer(w, r)
	case http.MethodPut:
		h.UpdateCustomer(w, r)
	case http.MethodDelete:
		h.DeleteCustomer(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleOrders routes requests to /orders
func (h *Handler) handleOrders(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateOrder(w, r)
	case http.MethodGet:
		h.GetAllOrders(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleOrderByID routes requests to /orders/{id}
func (h *Handler) handleOrderByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetOrder(w, r)
	case http.MethodPut:
		h.UpdateOrder(w, r)
	case http.MethodDelete:
		h.DeleteOrder(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Helper function to check if path matches pattern (not used but kept for reference)
func _matchPath(path, pattern string) bool {
	return strings.HasPrefix(path, pattern)
}
