# Online Bookstore API with Periodic Sales Report

## Project Status

This project implements an Online Bookstore API with Periodic Sales Report generation. Currently, **Parts 1, 2, and 3** are completed.

## Completed Parts

### ✅ Part 1: Project Setup and Data Models
- [x] Go module initialized (`online-bookstore-api`)
- [x] Project structure organized (models, interfaces, stores packages)
- [x] All data models defined:
  - `Book` struct with nested `Author`
  - `Author` struct
  - `Customer` struct with nested `Address`
  - `Order` struct with nested `Customer` and `OrderItem[]`
  - `OrderItem` struct with nested `Book`
  - `Address` struct
  - `SalesReport` struct with nested `BookSales[]`
  - `BookSales` struct
  - `SearchCriteria` struct
  - `ErrorResponse` struct

### ✅ Part 2: Interfaces and In-Memory Stores
- [x] All interfaces defined:
  - `BookStore` interface
  - `AuthorStore` interface
  - `CustomerStore` interface
  - `OrderStore` interface (with `GetOrdersInTimeRange` method)
- [x] In-memory stores implemented with thread-safe access:
  - `InMemoryBookStore` with `sync.RWMutex`
  - `InMemoryAuthorStore` with `sync.RWMutex`
  - `InMemoryCustomerStore` with `sync.RWMutex`
  - `InMemoryOrderStore` with `sync.RWMutex`
- [x] Persistence layer implemented:
  - Save/load functionality to/from `database.json`
  - Automatic data loading on application start
  - Thread-safe data access throughout

### ✅ Part 3: RESTful API Endpoints
- [x] HTTP handlers package created (`handlers/`)
- [x] Books endpoints implemented:
  - [x] `POST /books` - Create a new book
  - [x] `GET /books/{id}` - Retrieve a book by ID
  - [x] `PUT /books/{id}` - Update a book
  - [x] `DELETE /books/{id}` - Delete a book
  - [x] `GET /books` - Search for books using query parameters (supports title, author_id, genre, min_price, max_price)
- [x] Authors endpoints implemented:
  - [x] `POST /authors` - Create a new author
  - [x] `GET /authors/{id}` - Retrieve an author by ID
  - [x] `PUT /authors/{id}` - Update an author
  - [x] `DELETE /authors/{id}` - Delete an author
  - [x] `GET /authors` - List all authors
- [x] Customers endpoints implemented:
  - [x] `POST /customers` - Create a new customer
  - [x] `GET /customers/{id}` - Retrieve a customer by ID
  - [x] `PUT /customers/{id}` - Update a customer
  - [x] `DELETE /customers/{id}` - Delete a customer
  - [x] `GET /customers` - List all customers
- [x] Orders endpoints implemented:
  - [x] `POST /orders` - Place a new order (validates customer and books exist, calculates total)
  - [x] `GET /orders/{id}` - Retrieve an order by ID
  - [x] `PUT /orders/{id}` - Update an order
  - [x] `DELETE /orders/{id}` - Delete an order
  - [x] `GET /orders` - List all orders
- [x] HTTP router set up using `net/http`
- [x] Server configured to listen on port `:8080`
- [x] Request logging middleware implemented
- [x] JSON request/response handling
- [x] Error handling with appropriate HTTP status codes

## Remaining Parts - To Do List

### 📋 Part 4: Concurrency and Context Handling
- [ ] Accept `context.Context` from HTTP requests in handlers
- [ ] Use contexts to handle client cancellations
- [ ] Check `ctx.Done()` in long-running operations
- [ ] Implement goroutines for concurrent order processing
- [ ] Ensure all handlers can handle multiple concurrent requests
- [ ] Add request timeouts where appropriate

### 📋 Part 5: Error Handling and Logging
- [ ] Implement consistent error response structure
- [ ] Return appropriate HTTP status codes:
  - [ ] `200 OK` for successful GET requests
  - [ ] `201 Created` for successful POST requests
  - [ ] `400 Bad Request` for invalid input
  - [ ] `404 Not Found` for missing resources
  - [ ] `500 Internal Server Error` for server errors
- [ ] Use `log` package to record:
  - [ ] API requests (method, path, status)
  - [ ] Errors and exceptions
  - [ ] Significant events (orders placed, reports generated)
- [ ] Create helper functions for JSON error responses

### 📋 Part 6: Periodic Sales Report Generation
- [ ] Create `reports/` package for report generation logic
- [ ] Implement `generateSalesReport()` function:
  - [ ] Fetch orders within the last 24 hours using `GetOrdersInTimeRange`
  - [ ] Calculate total revenue
  - [ ] Count total number of orders
  - [ ] Calculate total books sold
  - [ ] Identify top-selling books
  - [ ] Create `SalesReport` struct with aggregated data
- [ ] Implement report storage:
  - [ ] Create `output-reports/` directory if it doesn't exist
  - [ ] Save reports as JSON files with timestamp in filename (e.g., `report_090120250000.json`)
  - [ ] Format: `report_MMDDYYYYHHMM.json`
- [ ] Set up periodic background task:
  - [ ] Use `time.Ticker` to schedule daily execution (every 24 hours)
  - [ ] Run as a goroutine that doesn't block the main server
  - [ ] Handle context cancellation for graceful shutdown
- [ ] Implement Sales Report API endpoint:
  - [ ] `GET /reports/sales?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD`
  - [ ] Parse query parameters for date range
  - [ ] Load and return matching reports from `output-reports/` directory
  - [ ] Return JSON array of reports
- [ ] Integrate background task with main server:
  - [ ] Start report generator when server starts
  - [ ] Handle graceful shutdown using context

### 📋 Part 7: Documentation
- [ ] Update README.md with:
  - [ ] How to build and run the application
  - [ ] API endpoint documentation with examples
  - [ ] Request/response examples for each endpoint
  - [ ] Environment variables or configuration options
  - [ ] Manual test cases showcasing functionality
- [ ] Create OpenAPI/Swagger specification file:
  - [ ] Define all endpoints
  - [ ] Document request/response schemas
  - [ ] Include example requests and responses
  - [ ] Save as `openapi.yaml` or `swagger.json`

### 📋 Part 8: Testing and Finalization
- [ ] Test all CRUD operations for each entity
- [ ] Test concurrent request handling
- [ ] Test periodic report generation
- [ ] Test data persistence (save/load from `database.json`)
- [ ] Verify graceful shutdown
- [ ] Clean up any temporary files or test data
- [ ] Ensure code compiles without warnings
- [ ] Review code quality and organization

## Project Structure

```
online-bookstore-api/
├── main.go                 # Application entry point
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies checksum
├── database.json           # Persistent data storage (created at runtime)
├── output-reports/         # Sales reports directory (created at runtime)
├── models/
│   └── models.go          # Data model definitions
├── interfaces/
│   └── interfaces.go      # Interface definitions
├── stores/
│   ├── bookstore.go       # In-memory book store implementation
│   ├── authorstore.go     # In-memory author store implementation
│   ├── customerstore.go   # In-memory customer store implementation
│   ├── orderstore.go      # In-memory order store implementation
│   └── persistence.go    # Save/load functionality
├── handlers/              # HTTP handlers (to be implemented)
├── reports/               # Report generation logic (to be implemented)
└── README.md              # This file
```

## How to Build

```bash
go build -o bookstore.exe .
```

## How to Run

```bash
./bookstore.exe
```

The server will start on `http://localhost:8080`. You can test the API endpoints using tools like `curl` or Postman.

### Example API Calls

**Create an Author:**
```bash
curl -X POST http://localhost:8080/authors \
  -H "Content-Type: application/json" \
  -d '{"first_name": "John", "last_name": "Doe", "bio": "Software Engineer"}'
```

**Create a Book:**
```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title": "Effective Go Concurrency", "author": {"id": 1}, "genres": ["Programming"], "published_at": "2021-07-15T00:00:00Z", "price": 39.99, "stock": 100}'
```

**Search Books:**
```bash
curl http://localhost:8080/books?title=Go&genre=Programming
```

**Create a Customer:**
```bash
curl -X POST http://localhost:8080/customers \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Smith", "email": "jane@example.com", "address": {"street": "123 Main St", "city": "New York", "state": "NY", "postal_code": "10001", "country": "USA"}}'
```

**Create an Order:**
```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"customer": {"id": 1}, "items": [{"book": {"id": 1}, "quantity": 2}], "status": "pending"}'
```

## Next Steps

1. Start with **Part 3** to implement the RESTful API endpoints
2. Test each endpoint as you implement it
3. Move to **Part 4** for concurrency and context handling
4. Implement error handling and logging in **Part 5**
5. Add the periodic sales report in **Part 6**
6. Complete documentation in **Part 7**
7. Finalize with testing in **Part 8**

## Notes

- All stores are thread-safe using `sync.RWMutex`
- Data is automatically saved to `database.json` (you'll need to implement the save trigger)
- Data is automatically loaded from `database.json` on application start
- The project uses only Go standard library packages
