# Online Bookstore API with Periodic Sales Report

## Project Status

This project implements an Online Bookstore API with Periodic Sales Report generation. Currently, **Parts 1 and 2** are completed.

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

## Remaining Parts - To Do List

### 📋 Part 3: RESTful API Endpoints
- [ ] Create HTTP handlers package (`handlers/`)
- [ ] Implement Books endpoints:
  - [ ] `POST /books` - Create a new book
  - [ ] `GET /books/{id}` - Retrieve a book by ID
  - [ ] `PUT /books/{id}` - Update a book
  - [ ] `DELETE /books/{id}` - Delete a book
  - [ ] `GET /books` - Search for books using query parameters
- [ ] Implement Authors endpoints:
  - [ ] `POST /authors` - Create a new author
  - [ ] `GET /authors/{id}` - Retrieve an author by ID
  - [ ] `PUT /authors/{id}` - Update an author
  - [ ] `DELETE /authors/{id}` - Delete an author
  - [ ] `GET /authors` - List all authors
- [ ] Implement Customers endpoints:
  - [ ] `POST /customers` - Create a new customer
  - [ ] `GET /customers/{id}` - Retrieve a customer by ID
  - [ ] `PUT /customers/{id}` - Update a customer
  - [ ] `DELETE /customers/{id}` - Delete a customer
  - [ ] `GET /customers` - List all customers
- [ ] Implement Orders endpoints:
  - [ ] `POST /orders` - Place a new order
  - [ ] `GET /orders/{id}` - Retrieve an order by ID
  - [ ] `PUT /orders/{id}` - Update an order
  - [ ] `DELETE /orders/{id}` - Delete an order
  - [ ] `GET /orders` - List all orders
- [ ] Set up HTTP router (using `net/http` or a router library)
- [ ] Configure server to listen on a port (e.g., `:8080`)

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

**Note:** The API server is not yet implemented. Once Part 3 is completed, the server will start and listen on a configured port.

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
