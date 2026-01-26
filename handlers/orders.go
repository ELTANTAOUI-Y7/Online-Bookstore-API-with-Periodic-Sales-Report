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

// CreateOrder handles POST /orders with context support and concurrent processing
func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Get context from request with timeout
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	// Check if context is already done
	if checkContext(ctx, w) {
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate required fields
	if len(order.Items) == 0 {
		respondWithError(w, http.StatusBadRequest, "Order must contain at least one item")
		return
	}

	// Verify customer exists (with context check)
	if checkContext(ctx, w) {
		return
	}
	customer, err := h.CustomerStore.GetCustomer(order.Customer.ID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Customer not found")
		return
	}
	order.Customer = customer

	// Verify all books exist and calculate total price (with context checks)
	totalPrice := 0.0
	for i, item := range order.Items {
		// Check context before each book lookup
		if checkContext(ctx, w) {
			return
		}

		book, err := h.BookStore.GetBook(item.Book.ID)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Book not found")
			return
		}
		// Update the book in the item with full book details
		order.Items[i].Book = book
		totalPrice += book.Price * float64(item.Quantity)
	}

	// Set order details
	order.TotalPrice = totalPrice
	if order.CreatedAt.IsZero() {
		order.CreatedAt = time.Now()
	}
	if order.Status == "" {
		order.Status = "pending"
	}

	// Final context check before creating order
	if checkContext(ctx, w) {
		return
	}

	// Create order in a goroutine for concurrent processing
	orderChan := make(chan models.Order, 1)
	errChan := make(chan error, 1)

	go func() {
		createdOrder, err := h.OrderStore.CreateOrder(order)
		if err != nil {
			errChan <- err
			return
		}
		orderChan <- createdOrder
	}()

	// Wait for order creation or context cancellation
	select {
	case <-ctx.Done():
		respondWithError(w, http.StatusRequestTimeout, "Request timeout while processing order")
		return
	case err := <-errChan:
		log.Printf("Error creating order: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create order")
		return
	case createdOrder := <-orderChan:
		log.Printf("Order created: ID=%d, Customer=%d, Total=%.2f", createdOrder.ID, createdOrder.Customer.ID, createdOrder.TotalPrice)
		respondWithJSON(w, http.StatusCreated, createdOrder)
	}
}

// GetOrder handles GET /orders/{id} with context support
func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx := r.Context()
	if checkContext(ctx, w) {
		return
	}

	id, err := extractID(r.URL.Path, "/orders/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	// Use goroutine for concurrent order retrieval
	orderChan := make(chan models.Order, 1)
	errChan := make(chan error, 1)

	go func() {
		order, err := h.OrderStore.GetOrder(id)
		if err != nil {
			errChan <- err
			return
		}
		orderChan <- order
	}()

	select {
	case <-ctx.Done():
		respondWithError(w, http.StatusRequestTimeout, "Request was canceled")
		return
	case err := <-errChan:
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "Order not found")
		} else {
			log.Printf("Error getting order: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to retrieve order")
		}
		return
	case order := <-orderChan:
		respondWithJSON(w, http.StatusOK, order)
	}
}

// UpdateOrder handles PUT /orders/{id} with context support
func (h *Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if checkContext(ctx, w) {
		return
	}

	id, err := extractID(r.URL.Path, "/orders/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if checkContext(ctx, w) {
		return
	}

	updatedOrder, err := h.OrderStore.UpdateOrder(id, order)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "Order not found")
		} else {
			log.Printf("Error updating order: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to update order")
		}
		return
	}

	log.Printf("Order updated: ID=%d", updatedOrder.ID)
	respondWithJSON(w, http.StatusOK, updatedOrder)
}

// DeleteOrder handles DELETE /orders/{id} with context support
func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if checkContext(ctx, w) {
		return
	}

	id, err := extractID(r.URL.Path, "/orders/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	if checkContext(ctx, w) {
		return
	}

	if err := h.OrderStore.DeleteOrder(id); err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "Order not found")
		} else {
			log.Printf("Error deleting order: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Failed to delete order")
		}
		return
	}

	log.Printf("Order deleted: ID=%d", id)
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Order deleted successfully"})
}

// GetAllOrders handles GET /orders with context support
func (h *Handler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	if checkContext(ctx, w) {
		return
	}

	// Use goroutine for concurrent order retrieval
	ordersChan := make(chan []models.Order, 1)
	errChan := make(chan error, 1)

	go func() {
		orders, err := h.OrderStore.GetAllOrders()
		if err != nil {
			errChan <- err
			return
		}
		ordersChan <- orders
	}()

	select {
	case <-ctx.Done():
		respondWithError(w, http.StatusRequestTimeout, "Request timeout while retrieving orders")
		return
	case err := <-errChan:
		log.Printf("Error getting all orders: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve orders")
		return
	case orders := <-ordersChan:
		respondWithJSON(w, http.StatusOK, orders)
	}
}

