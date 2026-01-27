package handlers

import (
	"log"
	"time"
)

// LogEvent logs a significant event with timestamp
func LogEvent(eventType, message string, details ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("[%s] EVENT: %s - %s %v", timestamp, eventType, message, details)
}

// LogError logs an error with context
func LogError(operation, message string, err error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	if err != nil {
		log.Printf("[%s] ERROR: %s - %s: %v", timestamp, operation, message, err)
	} else {
		log.Printf("[%s] ERROR: %s - %s", timestamp, operation, message)
	}
}

// LogInfo logs informational messages
func LogInfo(operation, message string, details ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("[%s] INFO: %s - %s %v", timestamp, operation, message, details)
}

// LogOrderPlaced logs when an order is successfully placed
func LogOrderPlaced(orderID int, customerID int, totalPrice float64, itemCount int) {
	LogEvent("ORDER_PLACED", 
		"Order successfully created",
		map[string]interface{}{
			"order_id": orderID,
			"customer_id": customerID,
			"total_price": totalPrice,
			"item_count": itemCount,
		},
	)
}

// LogBookCreated logs when a book is created
func LogBookCreated(bookID int, title string, authorID int) {
	LogEvent("BOOK_CREATED",
		"Book successfully created",
		map[string]interface{}{
			"book_id": bookID,
			"title": title,
			"author_id": authorID,
		},
	)
}

// LogAuthorCreated logs when an author is created
func LogAuthorCreated(authorID int, firstName, lastName string) {
	LogEvent("AUTHOR_CREATED",
		"Author successfully created",
		map[string]interface{}{
			"author_id": authorID,
			"name": firstName + " " + lastName,
		},
	)
}

// LogCustomerCreated logs when a customer is created
func LogCustomerCreated(customerID int, name, email string) {
	LogEvent("CUSTOMER_CREATED",
		"Customer successfully created",
		map[string]interface{}{
			"customer_id": customerID,
			"name": name,
			"email": email,
		},
	)
}

// LogUpdate logs when an entity is updated
func LogUpdate(entityType string, entityID int, details map[string]interface{}) {
	LogEvent("UPDATE",
		entityType+" updated",
		map[string]interface{}{
			"entity_type": entityType,
			"entity_id": entityID,
			"details": details,
		},
	)
}

// LogDelete logs when an entity is deleted
func LogDelete(entityType string, entityID int) {
	LogEvent("DELETE",
		entityType+" deleted",
		map[string]interface{}{
			"entity_type": entityType,
			"entity_id": entityID,
		},
	)
}
