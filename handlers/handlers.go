package handlers

import (
	"online-bookstore-api/interfaces"
)

// Handler holds references to all stores and report output directory
type Handler struct {
	BookStore        interfaces.BookStore
	AuthorStore      interfaces.AuthorStore
	CustomerStore    interfaces.CustomerStore
	OrderStore       interfaces.OrderStore
	ReportOutputDir  string
}

// NewHandler creates a new handler instance
func NewHandler(
	bookStore interfaces.BookStore,
	authorStore interfaces.AuthorStore,
	customerStore interfaces.CustomerStore,
	orderStore interfaces.OrderStore,
	reportOutputDir string,
) *Handler {
	if reportOutputDir == "" {
		reportOutputDir = "output-reports"
	}
	return &Handler{
		BookStore:       bookStore,
		AuthorStore:     authorStore,
		CustomerStore:   customerStore,
		OrderStore:      orderStore,
		ReportOutputDir: reportOutputDir,
	}
}

