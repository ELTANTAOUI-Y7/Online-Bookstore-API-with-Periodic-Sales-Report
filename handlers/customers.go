package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"online-bookstore-api/models"
	"strings"
	"time"
)

// CreateCustomer handles POST /customers
func (h *Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate required fields
	if customer.Name == "" || customer.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Name and email are required")
		return
	}

	// Set CreatedAt if not provided
	if customer.CreatedAt.IsZero() {
		customer.CreatedAt = time.Now()
	}

	createdCustomer, err := h.CustomerStore.CreateCustomer(customer)
	if err != nil {
		log.Printf("Error creating customer: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create customer")
		return
	}

	log.Printf("Customer created: ID=%d, Name=%s", createdCustomer.ID, createdCustomer.Name)
	respondWithJSON(w, http.StatusCreated, createdCustomer)
}

// GetCustomer handles GET /customers/{id}
func (h *Handler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := extractID(r.URL.Path, "/customers/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	customer, err := h.CustomerStore.GetCustomer(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Customer not found")
		return
	}

	respondWithJSON(w, http.StatusOK, customer)
}

// UpdateCustomer handles PUT /customers/{id}
func (h *Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := extractID(r.URL.Path, "/customers/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	var customer models.Customer
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updatedCustomer, err := h.CustomerStore.UpdateCustomer(id, customer)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "Customer not found")
		} else {
			log.Printf("Error updating customer: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to update customer")
		}
		return
	}

	log.Printf("Customer updated: ID=%d", updatedCustomer.ID)
	respondWithJSON(w, http.StatusOK, updatedCustomer)
}

// DeleteCustomer handles DELETE /customers/{id}
func (h *Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := extractID(r.URL.Path, "/customers/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	if err := h.CustomerStore.DeleteCustomer(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "Customer not found")
		} else {
			log.Printf("Error deleting customer: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to delete customer")
		}
		return
	}

	log.Printf("Customer deleted: ID=%d", id)
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Customer deleted successfully"})
}

// GetAllCustomers handles GET /customers
func (h *Handler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	customers, err := h.CustomerStore.GetAllCustomers()
	if err != nil {
		log.Printf("Error getting all customers: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve customers")
		return
	}

	respondWithJSON(w, http.StatusOK, customers)
}

