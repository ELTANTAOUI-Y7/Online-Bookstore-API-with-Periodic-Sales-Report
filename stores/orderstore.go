package stores

import (
	"fmt"
	"online-bookstore-api/interfaces"
	"online-bookstore-api/models"
	"sync"
	"time"
)

// InMemoryOrderStore implements OrderStore interface
type InMemoryOrderStore struct {
	mu     sync.RWMutex
	orders map[int]models.Order
	nextID int
}

// NewInMemoryOrderStore creates a new in-memory order store
func NewInMemoryOrderStore() *InMemoryOrderStore {
	return &InMemoryOrderStore{
		orders: make(map[int]models.Order),
		nextID: 1,
	}
}

// CreateOrder creates a new order
func (s *InMemoryOrderStore) CreateOrder(order models.Order) (models.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order.ID = s.nextID
	s.nextID++
	s.orders[order.ID] = order
	return order, nil
}

// GetOrder retrieves an order by ID
func (s *InMemoryOrderStore) GetOrder(id int) (models.Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	order, exists := s.orders[id]
	if !exists {
		return models.Order{}, fmt.Errorf("order with ID %d not found", id)
	}
	return order, nil
}

// UpdateOrder updates an existing order
func (s *InMemoryOrderStore) UpdateOrder(id int, order models.Order) (models.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.orders[id]; !exists {
		return models.Order{}, fmt.Errorf("order with ID %d not found", id)
	}

	order.ID = id
	s.orders[id] = order
	return order, nil
}

// DeleteOrder deletes an order by ID
func (s *InMemoryOrderStore) DeleteOrder(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.orders[id]; !exists {
		return fmt.Errorf("order with ID %d not found", id)
	}

	delete(s.orders, id)
	return nil
}

// GetAllOrders returns all orders
func (s *InMemoryOrderStore) GetAllOrders() ([]models.Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	orders := make([]models.Order, 0, len(s.orders))
	for _, order := range s.orders {
		orders = append(orders, order)
	}
	return orders, nil
}

// GetOrdersInTimeRange retrieves orders within a time range
func (s *InMemoryOrderStore) GetOrdersInTimeRange(start, end time.Time) ([]models.Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var results []models.Order
	for _, order := range s.orders {
		if (order.CreatedAt.After(start) || order.CreatedAt.Equal(start)) &&
			(order.CreatedAt.Before(end) || order.CreatedAt.Equal(end)) {
			results = append(results, order)
		}
	}
	return results, nil
}

// GetData returns the internal data for persistence
func (s *InMemoryOrderStore) GetData() map[int]models.Order {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data := make(map[int]models.Order)
	for k, v := range s.orders {
		data[k] = v
	}
	return data
}

// LoadData loads data from persistence
func (s *InMemoryOrderStore) LoadData(data map[int]models.Order, nextID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orders = data
	s.nextID = nextID
}

// GetNextID returns the next ID that will be used
func (s *InMemoryOrderStore) GetNextID() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.nextID
}

// Verify interface implementation
var _ interfaces.OrderStore = (*InMemoryOrderStore)(nil)
