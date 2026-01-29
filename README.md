# Online Bookstore API with Periodic Sales Report

## Project Status

This project implements an Online Bookstore API with Periodic Sales Report generation. Currently, **Parts 1, 2, 3, 5, 6, 7, 8, and 11** are completed.

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
  - `SalesReport` struct (timestamp, total revenue, total orders, total books sold, top-selling books) with nested `BookSales[]`
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

### ✅ Part 11: Periodic Sales Report Generation
- [x] `reports/` package created with report generation logic
- [x] `GenerateSalesReport()` function:
  - [x] Fetches orders within the last 24 hours using `OrderStore.GetOrdersInTimeRange`
  - [x] Aggregates total revenue, total orders, total books sold
  - [x] Identifies top-selling books (top 10)
  - [x] Returns `SalesReport` with all aggregated data
- [x] Report storage:
  - [x] Reports saved to `output-reports/` (directory created if needed)
  - [x] Filename format: `report_DDMMYYYYHHmm.json` (e.g. `report_250120261430.json`)
  - [x] JSON content with timestamp, revenue, orders, books sold, top-selling books
- [x] Background task:
  - [x] Goroutine runs report at startup, then every 24 hours via `time.Ticker`
  - [x] Context-based cancellation for graceful shutdown (`reportCancel()` on SIGINT/SIGTERM)
  - [x] Concurrency-safe: uses store's existing mutex via `GetOrdersInTimeRange`
- [x] Sales Report API endpoint:
  - [x] `GET /reports/sales?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD` (query params optional; default last 30 days)
  - [x] Loads matching reports from `output-reports/` and returns JSON array
- [x] Integration in `main.go`: report generator started with server, stopped before shutdown

## Remaining Parts - To Do List

### 📋 Documentation
- [x] README with build/run instructions and API overview
- [ ] Environment variables or configuration options (if needed)
- [ ] Manual test cases section
- [x] OpenAPI specification in `openapi.yaml` (endpoints, schemas, examples)

### 📋 Testing and Finalization
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
├── handlers/              # HTTP handlers (books, authors, customers, orders, reports)
├── reports/               # Sales report generation and file loader
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

## API Overview

Base URL: `http://localhost:8080`

### Main Resources

- `Authors` – `/authors`, `/authors/{id}`
- `Books` – `/books`, `/books/{id}`
- `Customers` – `/customers`, `/customers/{id}`
- `Orders` – `/orders`, `/orders/{id}`
- `Reports` – `/reports/sales` (GET, optional `start_date` & `end_date` query params)

## API Endpoints (Summary)

### Authors

- `GET /authors` – List all authors
- `POST /authors` – Create a new author
- `GET /authors/{id}` – Get a specific author
- `PUT /authors/{id}` – Update an author
- `DELETE /authors/{id}` – Delete an author

**Example – Create Author (request):**

```bash
curl -X POST http://localhost:8080/authors \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "bio": "Software Engineer with expertise in Go."
  }'
```

**Example – Create Author (response):**

```json
{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "bio": "Software Engineer with expertise in Go."
}
```

### Books

- `GET /books` – List or search books
  - Query params: `title`, `author_id`, `genre`, `min_price`, `max_price`
- `POST /books` – Create a new book
- `GET /books/{id}` – Get a specific book
- `PUT /books/{id}` – Update a book
- `DELETE /books/{id}` – Delete a book

**Example – Create Book (request):**

```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Effective Go Concurrency",
    "author": { "id": 1 },
    "genres": ["Programming", "Technology"],
    "published_at": "2021-07-15T00:00:00Z",
    "price": 39.99,
    "stock": 100
  }'
```

**Example – Get Book (response):**

```json
{
  "id": 1,
  "title": "Effective Go Concurrency",
  "author": {
    "id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "bio": "Software Engineer with expertise in Go."
  },
  "genres": ["Programming", "Technology"],
  "published_at": "2021-07-15T00:00:00Z",
  "price": 39.99,
  "stock": 100
}
```

### Customers

- `GET /customers` – List all customers
- `POST /customers` – Create a new customer
- `GET /customers/{id}` – Get a specific customer
- `PUT /customers/{id}` – Update a customer
- `DELETE /customers/{id}` – Delete a customer

**Example – Create Customer (request):**

```bash
curl -X POST http://localhost:8080/customers \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Smith",
    "email": "jane@example.com",
    "address": {
      "street": "123 Main St",
      "city": "New York",
      "state": "NY",
      "postal_code": "10001",
      "country": "USA"
    }
  }'
```

### Orders

- `GET /orders` – List all orders
- `POST /orders` – Create a new order
- `GET /orders/{id}` – Get a specific order
- `PUT /orders/{id}` – Update an order
- `DELETE /orders/{id}` – Delete an order

### Sales Reports

- `GET /reports/sales` – List sales reports in a date range
  - Query params (optional): `start_date=YYYY-MM-DD`, `end_date=YYYY-MM-DD`
  - If omitted, returns reports from the last 30 days
  - Reports are generated automatically every 24 hours and at startup; each report covers the previous 24 hours (total revenue, total orders, total books sold, top-selling books)

**Example – Get Sales Reports (request):**

```bash
# All reports in default range (last 30 days)
curl http://localhost:8080/reports/sales

# Reports between two dates
curl "http://localhost:8080/reports/sales?start_date=2026-01-01&end_date=2026-01-25"
```

**Example – Sales Report (response item):**

```json
{
  "timestamp": "2026-01-25T14:30:00Z",
  "total_revenue": 159.96,
  "total_orders": 2,
  "total_books_sold": 4,
  "top_selling_books": [
    {
      "book": { "id": 1, "title": "Effective Go Concurrency", "author": {...}, "genres": ["Programming"], "published_at": "2021-07-15T00:00:00Z", "price": 39.99, "stock": 98 },
      "quantity_sold": 3
    }
  ]
}
```

**Example – Create Order (request):**

```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{
    "customer": { "id": 1 },
    "items": [
      {
        "book": { "id": 1 },
        "quantity": 2
      }
    ],
    "status": "pending"
  }'
```

**Example – Create Order (response):**

```json
{
  "id": 1,
  "customer": {
    "id": 1,
    "name": "Jane Smith",
    "email": "jane@example.com",
    "address": {
      "street": "123 Main St",
      "city": "New York",
      "state": "NY",
      "postal_code": "10001",
      "country": "USA"
    },
    "created_at": "2026-01-25T10:00:00Z"
  },
  "items": [
    {
      "book": {
        "id": 1,
        "title": "Effective Go Concurrency",
        "author": {
          "id": 1,
          "first_name": "John",
          "last_name": "Doe",
          "bio": "Software Engineer with expertise in Go."
        },
        "genres": ["Programming", "Technology"],
        "published_at": "2021-07-15T00:00:00Z",
        "price": 39.99,
        "stock": 100
      },
      "quantity": 2
    }
  ],
  "total_price": 79.98,
  "created_at": "2026-01-25T10:05:00Z",
  "status": "pending"
}
```

## Swagger / OpenAPI Specification

An OpenAPI 3.0 specification is provided in `openapi.yaml` at the root of the project.

You can use this file with tools like:

- Swagger UI
- Postman
- Insomnia

To visualize and interact with the API.

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

**Get Sales Reports:**
```bash
curl http://localhost:8080/reports/sales
curl "http://localhost:8080/reports/sales?start_date=2026-01-01&end_date=2026-01-25"
```

## Next Steps

1. Run the server and test all CRUD and report endpoints
2. Add or extend OpenAPI spec for `/reports/sales` if desired
3. Add automated tests (unit/integration) and manual test cases
4. Optionally add environment variables for port and report output directory

## Notes

- All stores are thread-safe using `sync.RWMutex`
- Data is automatically saved to `database.json` on graceful shutdown (SIGINT/SIGTERM)
- Data is automatically loaded from `database.json` on application start
- Sales reports run at startup and every 24 hours; reports are written to `output-reports/` as JSON files
- The project uses only Go standard library packages
