package interfaces

import (
	"online-bookstore-api/models"
	"time"
)

// BookStore defines operations for book management
type BookStore interface {
	CreateBook(book models.Book) (models.Book, error)
	GetBook(id int) (models.Book, error)
	UpdateBook(id int, book models.Book) (models.Book, error)
	DeleteBook(id int) error
	SearchBooks(criteria models.SearchCriteria) ([]models.Book, error)
	GetAllBooks() ([]models.Book, error)
}

// AuthorStore defines operations for author management
type AuthorStore interface {
	CreateAuthor(author models.Author) (models.Author, error)
	GetAuthor(id int) (models.Author, error)
	UpdateAuthor(id int, author models.Author) (models.Author, error)
	DeleteAuthor(id int) error
	GetAllAuthors() ([]models.Author, error)
}

// CustomerStore defines operations for customer management
type CustomerStore interface {
	CreateCustomer(customer models.Customer) (models.Customer, error)
	GetCustomer(id int) (models.Customer, error)
	UpdateCustomer(id int, customer models.Customer) (models.Customer, error)
	DeleteCustomer(id int) error
	GetAllCustomers() ([]models.Customer, error)
}

// OrderStore defines operations for order management
type OrderStore interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetOrder(id int) (models.Order, error)
	UpdateOrder(id int, order models.Order) (models.Order, error)
	DeleteOrder(id int) error
	GetAllOrders() ([]models.Order, error)
	GetOrdersInTimeRange(start, end time.Time) ([]models.Order, error)
}
