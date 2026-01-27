package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"online-bookstore-api/models"
	"strings"
	"time"
)

// CreateCustomer handles POST /customers with context support
func (h *Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if checkContext(ctx, w) {
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

	if checkContext(ctx, w) {
		return
	}

	createdCustomer, err := h.CustomerStore.CreateCustomer(customer)
	if err != nil {
		LogError("CreateCustomer", "Failed to create customer", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create customer")
		return
	}

	LogCustomerCreated(createdCustomer.ID, createdCustomer.Name, createdCustomer.Email)
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

// UpdateCustomer handles PUT /customers/{id} with context support
func (h *Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if checkContext(ctx, w) {
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

	if checkContext(ctx, w) {
		return
	}

	updatedCustomer, err := h.CustomerStore.UpdateCustomer(id, customer)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			LogInfo("UpdateCustomer", "Customer not found", map[string]interface{}{"customer_id": id})
			respondWithError(w, http.StatusNotFound, "Customer not found")
		} else {
			LogError("UpdateCustomer", "Failed to update customer", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to update customer")
		}
		return
	}

	LogUpdate("Customer", updatedCustomer.ID, map[string]interface{}{
		"name": updatedCustomer.Name,
		"email": updatedCustomer.Email,
	})
	respondWithJSON(w, http.StatusOK, updatedCustomer)
}

// DeleteCustomer handles DELETE /customers/{id} with context support
func (h *Handler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if checkContext(ctx, w) {
		return
	}

	id, err := extractID(r.URL.Path, "/customers/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	if checkContext(ctx, w) {
		return
	}

	if err := h.CustomerStore.DeleteCustomer(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			LogInfo("DeleteCustomer", "Customer not found", map[string]interface{}{"customer_id": id})
			respondWithError(w, http.StatusNotFound, "Customer not found")
		} else {
			LogError("DeleteCustomer", "Failed to delete customer", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to delete customer")
		}
		return
	}

	LogDelete("Customer", id)
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Customer deleted successfully"})
}

// GetAllCustomers handles GET /customers with context support
func (h *Handler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx := r.Context()
	if checkContext(ctx, w) {
		return
	}

	customers, err := h.CustomerStore.GetAllCustomers()
	if err != nil {
		LogError("GetAllCustomers", "Failed to retrieve customers", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve customers")
		return
	}

	LogInfo("GetAllCustomers", "Retrieved all customers", map[string]interface{}{"count": len(customers)})
	respondWithJSON(w, http.StatusOK, customers)
}

