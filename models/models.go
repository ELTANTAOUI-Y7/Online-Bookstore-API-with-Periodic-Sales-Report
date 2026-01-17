package models

import "time"

// Book represents a book in the bookstore
type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      Author    `json:"author"`
	Genres      []string  `json:"genres"`
	PublishedAt time.Time `json:"published_at"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
}

// Author represents an author
type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
}

// Address represents a customer's address
type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

// Customer represents a customer
type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   Address   `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	Book     Book `json:"book"`
	Quantity int  `json:"quantity"`
}

// Order represents an order
type Order struct {
	ID         int         `json:"id"`
	Customer   Customer    `json:"customer"`
	Items      []OrderItem `json:"items"`
	TotalPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	Status     string      `json:"status"`
}

// BookSales represents book sales data for reports
type BookSales struct {
	Book     Book `json:"book"`
	Quantity int  `json:"quantity_sold"`
}

// SalesReport represents a sales report
type SalesReport struct {
	Timestamp       time.Time   `json:"timestamp"`
	TotalRevenue    float64     `json:"total_revenue"`
	TotalOrders     int         `json:"total_orders"`
	TopSellingBooks []BookSales `json:"top_selling_books"`
}

// SearchCriteria represents search parameters for books
type SearchCriteria struct {
	Title    string
	AuthorID int
	Genre    string
	MinPrice float64
	MaxPrice float64
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}
