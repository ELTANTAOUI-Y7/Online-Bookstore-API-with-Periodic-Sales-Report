package handlers

import (
	"online-bookstore-api/interfaces"
)

// Handler holds references to all stores
type Handler struct {
	BookStore     interfaces.BookStore
	AuthorStore   interfaces.AuthorStore
	CustomerStore interfaces.CustomerStore
	OrderStore    interfaces.OrderStore
}

// NewHandler creates a new handler instance
func NewHandler(
	bookStore interfaces.BookStore,
	authorStore interfaces.AuthorStore,
	customerStore interfaces.CustomerStore,
	orderStore interfaces.OrderStore,
) *Handler {
	return &Handler{
		BookStore:     bookStore,
		AuthorStore:   authorStore,
		CustomerStore: customerStore,
		OrderStore:    orderStore,
	}
}

