# Online Bookstore API with Periodic Sales Report

## Project Status

This project implements an Online Bookstore API with Periodic Sales Report generation. Currently, **Parts 1, 2, 3, 5, 6, 7, and 8** are completed.

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

### ✅ Part 5: Concurrency and Synchronization
- [x] All stores use `sync.RWMutex` for thread-safe access
- [x] Handlers can handle multiple concurrent requests without data corruption
- [x] Goroutines implemented for concurrent order processing (CreateOrder, GetOrder, GetAllOrders)
- [x] Mutex synchronization verified in all stores (BookStore, AuthorStore, CustomerStore, OrderStore)
- [x] Read locks (RLock) for read operations, write locks (Lock) for write operations

### ✅ Part 6: Context for Cancellation and Timeouts
- [x] All handlers accept `context.Context` from HTTP requests
- [x] Context timeouts implemented (5-10 seconds depending on operation)
- [x] `ctx.Done()` checks in all handlers and long-running operations
- [x] Proper handling of context cancellation and deadline exceeded
- [x] Context-aware error responses for timeouts and cancellations
- [x] Context checks before and during operations

### ✅ Part 7: Error Handling and Responses
- [x] Consistent error response structure using `ErrorResponse` struct
- [x] Appropriate HTTP status codes:
  - [x] `200 OK` for successful GET requests
  - [x] `201 Created` for successful POST requests
  - [x] `400 Bad Request` for invalid input
  - [x] `404 Not Found` for missing resources
  - [x] `408 Request Timeout` for context timeouts/cancellations
  - [x] `500 Internal Server Error` for server errors
- [x] `log` package used to record:
  - [x] API requests (method, path) via middleware
  - [x] Errors and exceptions
  - [x] Significant events (orders placed, books created, etc.)
- [x] Helper functions for JSON error responses (`respondWithError`, `respondWithJSON`)

### ✅ Part 8: Logging
- [x] Enhanced logging middleware with detailed request information:
  - [x] HTTP method, path, protocol
  - [x] Response status codes
  - [x] Response size (bytes written)
  - [x] Request duration/timing
  - [x] Client IP address
  - [x] Error logging for 4xx and 5xx responses
- [x] Comprehensive logging utility functions (`handlers/logging.go`):
  - [x] `LogEvent()` - Log significant events with timestamps
  - [x] `LogError()` - Log errors with context
  - [x] `LogInfo()` - Log informational messages
  - [x] Specialized logging functions for specific events
- [x] All significant events logged:
  - [x] Orders placed (`LogOrderPlaced`) - includes order ID, customer ID, total price, item count
  - [x] Books created (`LogBookCreated`) - includes book ID, title, author ID
  - [x] Authors created (`LogAuthorCreated`) - includes author ID, name
  - [x] Customers created (`LogCustomerCreated`) - includes customer ID, name, email
  - [x] Updates logged (`LogUpdate`) - includes entity type, ID, and details
  - [x] Deletes logged (`LogDelete`) - includes entity type and ID
- [x] Error logging throughout:
  - [x] All error paths log errors with context
  - [x] Not found errors logged as info
  - [x] Server errors logged with full error details
  - [x] HTTP error responses logged automatically
- [x] Request logging:
  - [x] All API requests logged via middleware
  - [x] Search operations logged with criteria and result counts
  - [x] List operations logged with result counts

## Remaining Parts - To Do List

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
