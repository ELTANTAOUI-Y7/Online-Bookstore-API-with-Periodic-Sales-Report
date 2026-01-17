package main

import (
	"log"
	"online-bookstore-api/stores"
)

func main() {
	// Initialize stores
	bookStore := stores.NewInMemoryBookStore()
	authorStore := stores.NewInMemoryAuthorStore()
	customerStore := stores.NewInMemoryCustomerStore()
	orderStore := stores.NewInMemoryOrderStore()

	// Load data from persistence if it exists
	if err := stores.LoadDatabase(bookStore, authorStore, customerStore, orderStore, "database.json"); err != nil {
		log.Printf("Warning: Failed to load database: %v", err)
	}

	// TODO: Initialize HTTP server and handlers (Part 3)
	

	log.Println("Stores initialized successfully")
	log.Println("API server not yet implemented - see README for remaining tasks")

	// Now, I will just keep the stores in memory
	// In Part 3, I'll add the HTTP server
	_ = bookStore
	_ = authorStore
	_ = customerStore
	_ = orderStore
}
