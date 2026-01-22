package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"online-bookstore-api/models"
	"strings"
	"time"
)

// CreateOrder handles POST /orders
func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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

	// Verify customer exists
	if _, err := h.CustomerStore.GetCustomer(order.Customer.ID); err != nil {
		respondWithError(w, http.StatusBadRequest, "Customer not found")
		return
	}

	// Verify all books exist and calculate total price
	totalPrice := 0.0
	for i, item := range order.Items {
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

	createdOrder, err := h.OrderStore.CreateOrder(order)
	if err != nil {
		log.Printf("Error creating order: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create order")
		return
	}

	log.Printf("Order created: ID=%d, Customer=%d, Total=%.2f", createdOrder.ID, createdOrder.Customer.ID, createdOrder.TotalPrice)
	respondWithJSON(w, http.StatusCreated, createdOrder)
}

// GetOrder handles GET /orders/{id}
func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := extractID(r.URL.Path, "/orders/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
		return
	}

	order, err := h.OrderStore.GetOrder(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Order not found")
		return
	}

	respondWithJSON(w, http.StatusOK, order)
}

// UpdateOrder handles PUT /orders/{id}
func (h *Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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

// DeleteOrder handles DELETE /orders/{id}
func (h *Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := extractID(r.URL.Path, "/orders/")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid order ID")
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

// GetAllOrders handles GET /orders
func (h *Handler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	orders, err := h.OrderStore.GetAllOrders()
	if err != nil {
		log.Printf("Error getting all orders: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve orders")
		return
	}

	respondWithJSON(w, http.StatusOK, orders)
}

