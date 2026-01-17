package stores

import (
	"encoding/json"
	"fmt"
	"online-bookstore-api/models"
	"os"
)

// DatabaseData represents the complete database structure for persistence
type DatabaseData struct {
	Books     map[int]models.Book     `json:"books"`
	Authors   map[int]models.Author   `json:"authors"`
	Customers map[int]models.Customer `json:"customers"`
	Orders    map[int]models.Order    `json:"orders"`
	NextIDs   struct {
		Book     int `json:"book"`
		Author   int `json:"author"`
		Customer int `json:"customer"`
		Order    int `json:"order"`
	} `json:"next_ids"`
}

// SaveDatabase saves all stores to a JSON file
func SaveDatabase(
	bookStore *InMemoryBookStore,
	authorStore *InMemoryAuthorStore,
	customerStore *InMemoryCustomerStore,
	orderStore *InMemoryOrderStore,
	filename string,
) error {
	data := DatabaseData{
		Books:     bookStore.GetData(),
		Authors:   authorStore.GetData(),
		Customers: customerStore.GetData(),
		Orders:    orderStore.GetData(),
	}

	// Get next IDs from stores
	data.NextIDs.Book = bookStore.GetNextID()
	data.NextIDs.Author = authorStore.GetNextID()
	data.NextIDs.Customer = customerStore.GetNextID()
	data.NextIDs.Order = orderStore.GetNextID()

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode data: %w", err)
	}

	return nil
}

// LoadDatabase loads all stores from a JSON file
func LoadDatabase(
	bookStore *InMemoryBookStore,
	authorStore *InMemoryAuthorStore,
	customerStore *InMemoryCustomerStore,
	orderStore *InMemoryOrderStore,
	filename string,
) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, start with empty stores
			return nil
		}
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var data DatabaseData
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return fmt.Errorf("failed to decode data: %w", err)
	}

	// Load data into stores
	if data.Books != nil {
		bookStore.LoadData(data.Books, data.NextIDs.Book)
	}
	if data.Authors != nil {
		authorStore.LoadData(data.Authors, data.NextIDs.Author)
	}
	if data.Customers != nil {
		customerStore.LoadData(data.Customers, data.NextIDs.Customer)
	}
	if data.Orders != nil {
		orderStore.LoadData(data.Orders, data.NextIDs.Order)
	}

	return nil
}
