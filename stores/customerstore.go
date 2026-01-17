package stores

import (
	"fmt"
	"online-bookstore-api/interfaces"
	"online-bookstore-api/models"
	"sync"
)

// InMemoryCustomerStore implements CustomerStore interface
type InMemoryCustomerStore struct {
	mu        sync.RWMutex
	customers map[int]models.Customer
	nextID    int
}

// NewInMemoryCustomerStore creates a new in-memory customer store
func NewInMemoryCustomerStore() *InMemoryCustomerStore {
	return &InMemoryCustomerStore{
		customers: make(map[int]models.Customer),
		nextID:    1,
	}
}

// CreateCustomer creates a new customer
func (s *InMemoryCustomerStore) CreateCustomer(customer models.Customer) (models.Customer, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	customer.ID = s.nextID
	s.nextID++
	s.customers[customer.ID] = customer
	return customer, nil
}

// GetCustomer retrieves a customer by ID
func (s *InMemoryCustomerStore) GetCustomer(id int) (models.Customer, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	customer, exists := s.customers[id]
	if !exists {
		return models.Customer{}, fmt.Errorf("customer with ID %d not found", id)
	}
	return customer, nil
}

// UpdateCustomer updates an existing customer
func (s *InMemoryCustomerStore) UpdateCustomer(id int, customer models.Customer) (models.Customer, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.customers[id]; !exists {
		return models.Customer{}, fmt.Errorf("customer with ID %d not found", id)
	}

	customer.ID = id
	s.customers[id] = customer
	return customer, nil
}

// DeleteCustomer deletes a customer by ID
func (s *InMemoryCustomerStore) DeleteCustomer(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.customers[id]; !exists {
		return fmt.Errorf("customer with ID %d not found", id)
	}

	delete(s.customers, id)
	return nil
}

// GetAllCustomers returns all customers
func (s *InMemoryCustomerStore) GetAllCustomers() ([]models.Customer, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	customers := make([]models.Customer, 0, len(s.customers))
	for _, customer := range s.customers {
		customers = append(customers, customer)
	}
	return customers, nil
}

// GetData returns the internal data for persistence
func (s *InMemoryCustomerStore) GetData() map[int]models.Customer {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data := make(map[int]models.Customer)
	for k, v := range s.customers {
		data[k] = v
	}
	return data
}

// LoadData loads data from persistence
func (s *InMemoryCustomerStore) LoadData(data map[int]models.Customer, nextID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.customers = data
	s.nextID = nextID
}

// GetNextID returns the next ID that will be used
func (s *InMemoryCustomerStore) GetNextID() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.nextID
}

// Verify interface implementation
var _ interfaces.CustomerStore = (*InMemoryCustomerStore)(nil)
