package main

import (
	"context"
	"log"
	"net/http"
	"online-bookstore-api/handlers"
	"online-bookstore-api/stores"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	log.Println("Stores initialized successfully")

	// Initialize handlers
	handler := handlers.NewHandler(bookStore, authorStore, customerStore, orderStore)

	// Setup routes
	router := handler.SetupRoutes()

	// Add logging middleware
	loggedRouter := loggingMiddleware(router)

	// Create HTTP server
	port := ":8080"
	server := &http.Server{
		Addr:    port,
		Handler: loggedRouter,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", port)
		log.Printf("API endpoints available at http://localhost%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Save data before shutdown
	log.Println("Saving database...")
	if err := stores.SaveDatabase(bookStore, authorStore, customerStore, orderStore, "database.json"); err != nil {
		log.Printf("Error saving database: %v", err)
	} else {
		log.Println("Database saved successfully")
	}

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}

// loggingMiddleware logs HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
