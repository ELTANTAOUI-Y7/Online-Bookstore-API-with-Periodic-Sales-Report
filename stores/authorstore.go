package stores

import (
	"fmt"
	"online-bookstore-api/interfaces"
	"online-bookstore-api/models"
	"sync"
)

// InMemoryAuthorStore implements AuthorStore interface
type InMemoryAuthorStore struct {
	mu      sync.RWMutex
	authors map[int]models.Author
	nextID  int
}

// NewInMemoryAuthorStore creates a new in-memory author store
func NewInMemoryAuthorStore() *InMemoryAuthorStore {
	return &InMemoryAuthorStore{
		authors: make(map[int]models.Author),
		nextID:  1,
	}
}

// CreateAuthor creates a new author
func (s *InMemoryAuthorStore) CreateAuthor(author models.Author) (models.Author, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	author.ID = s.nextID
	s.nextID++
	s.authors[author.ID] = author
	return author, nil
}

// GetAuthor retrieves an author by ID
func (s *InMemoryAuthorStore) GetAuthor(id int) (models.Author, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	author, exists := s.authors[id]
	if !exists {
		return models.Author{}, fmt.Errorf("author with ID %d not found", id)
	}
	return author, nil
}

// UpdateAuthor updates an existing author
func (s *InMemoryAuthorStore) UpdateAuthor(id int, author models.Author) (models.Author, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.authors[id]; !exists {
		return models.Author{}, fmt.Errorf("author with ID %d not found", id)
	}

	author.ID = id
	s.authors[id] = author
	return author, nil
}

// DeleteAuthor deletes an author by ID
func (s *InMemoryAuthorStore) DeleteAuthor(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.authors[id]; !exists {
		return fmt.Errorf("author with ID %d not found", id)
	}

	delete(s.authors, id)
	return nil
}

// GetAllAuthors returns all authors
func (s *InMemoryAuthorStore) GetAllAuthors() ([]models.Author, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	authors := make([]models.Author, 0, len(s.authors))
	for _, author := range s.authors {
		authors = append(authors, author)
	}
	return authors, nil
}

// GetData returns the internal data for persistence
func (s *InMemoryAuthorStore) GetData() map[int]models.Author {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data := make(map[int]models.Author)
	for k, v := range s.authors {
		data[k] = v
	}
	return data
}

// LoadData loads data from persistence
func (s *InMemoryAuthorStore) LoadData(data map[int]models.Author, nextID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.authors = data
	s.nextID = nextID
}

// GetNextID returns the next ID that will be used
func (s *InMemoryAuthorStore) GetNextID() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.nextID
}

// Verify interface implementation
var _ interfaces.AuthorStore = (*InMemoryAuthorStore)(nil)
