package stores

import (
	"fmt"
	"online-bookstore-api/interfaces"
	"online-bookstore-api/models"
	"strings"
	"sync"
)

// InMemoryBookStore implements BookStore interface
type InMemoryBookStore struct {
	mu     sync.RWMutex
	books  map[int]models.Book
	nextID int
}

// NewInMemoryBookStore creates a new in-memory book store
func NewInMemoryBookStore() *InMemoryBookStore {
	return &InMemoryBookStore{
		books:  make(map[int]models.Book),
		nextID: 1,
	}
}

// CreateBook creates a new book
func (s *InMemoryBookStore) CreateBook(book models.Book) (models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	book.ID = s.nextID
	s.nextID++
	s.books[book.ID] = book
	return book, nil
}

// GetBook retrieves a book by ID
func (s *InMemoryBookStore) GetBook(id int) (models.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	book, exists := s.books[id]
	if !exists {
		return models.Book{}, fmt.Errorf("book with ID %d not found", id)
	}
	return book, nil
}

// UpdateBook updates an existing book
func (s *InMemoryBookStore) UpdateBook(id int, book models.Book) (models.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[id]; !exists {
		return models.Book{}, fmt.Errorf("book with ID %d not found", id)
	}

	book.ID = id
	s.books[id] = book
	return book, nil
}

// DeleteBook deletes a book by ID
func (s *InMemoryBookStore) DeleteBook(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.books[id]; !exists {
		return fmt.Errorf("book with ID %d not found", id)
	}

	delete(s.books, id)
	return nil
}

// SearchBooks searches for books based on criteria
func (s *InMemoryBookStore) SearchBooks(criteria models.SearchCriteria) ([]models.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var results []models.Book
	for _, book := range s.books {
		matches := true

		// Case-insensitive partial match for title
		if criteria.Title != "" {
			titleMatch := strings.Contains(strings.ToLower(book.Title), strings.ToLower(criteria.Title))
			if !titleMatch {
				matches = false
			}
		}
		if criteria.AuthorID != 0 && book.Author.ID != criteria.AuthorID {
			matches = false
		}
		if criteria.Genre != "" {
			genreFound := false
			for _, genre := range book.Genres {
				if strings.EqualFold(genre, criteria.Genre) {
					genreFound = true
					break
				}
			}
			if !genreFound {
				matches = false
			}
		}
		if criteria.MinPrice > 0 && book.Price < criteria.MinPrice {
			matches = false
		}
		if criteria.MaxPrice > 0 && book.Price > criteria.MaxPrice {
			matches = false
		}

		if matches {
			results = append(results, book)
		}
	}

	return results, nil
}

// GetAllBooks returns all books
func (s *InMemoryBookStore) GetAllBooks() ([]models.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	books := make([]models.Book, 0, len(s.books))
	for _, book := range s.books {
		books = append(books, book)
	}
	return books, nil
}

// GetData returns the internal data for persistence
func (s *InMemoryBookStore) GetData() map[int]models.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data := make(map[int]models.Book)
	for k, v := range s.books {
		data[k] = v
	}
	return data
}

// LoadData loads data from persistence
func (s *InMemoryBookStore) LoadData(data map[int]models.Book, nextID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.books = data
	s.nextID = nextID
}

// GetNextID returns the next ID that will be used
func (s *InMemoryBookStore) GetNextID() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.nextID
}

// Verify interface implementation
var _ interfaces.BookStore = (*InMemoryBookStore)(nil)
