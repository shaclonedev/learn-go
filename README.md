# Learn Go - REST API Project

A beginner-friendly Go project for learning RESTful API development using Go's standard library.

## 📖 Overview

This project is a learning exercise to understand Go programming and REST API development. It implements a simple Category management system with CRUD operations using only Go's `net/http` package without any external frameworks.

## 🚀 Features

- **RESTful API** endpoints for Category management
- Pure Go implementation using standard library
- JSON-based request/response handling
- HTTP routing and method handling
- ID parsing from URL paths

## 📦 Data Model

### Category

```go
type Category struct {
    ID          int64  `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
}
```

## 🔗 API Endpoints

Currently implemented endpoints:

| Method | Endpoint           | Description                 | Status         |
| ------ | ------------------ | --------------------------- | -------------- |
| GET    | `/categories`      | Get all categories          | ✅ Implemented |
| GET    | `/categories/{id}` | Get a single category by ID | ✅ Implemented |
| POST   | `/categories`      | Create a new category       | 🔄 In Progress |
| PUT    | `/categories/{id}` | Update an existing category | 🔄 In Progress |
| DELETE | `/categories/{id}` | Delete a category           | 🔄 In Progress |

## 🛠️ Installation & Setup

### Prerequisites

- Go 1.x or higher installed on your system
- Basic understanding of Go syntax and HTTP concepts

### Running the Application

1. Clone or navigate to the project directory:

```bash
cd /path/to/learn-go
```

2. Run the application:

```bash
go run main.go
```

3. The server will start on `http://localhost:8080`

## 📝 Usage Examples

### Get All Categories

```bash
curl http://localhost:8080/categories
```

**Response:**

```json
[
  {
    "id": 1,
    "name": "A",
    "description": "DESC A"
  },
  {
    "id": 2,
    "name": "B",
    "description": "DESC B"
  },
  {
    "id": 3,
    "name": "C",
    "description": "DESC C"
  }
]
```

### Get Single Category

```bash
curl http://localhost:8080/categories/1
```

**Response:**

```json
{
  "id": 1,
  "name": "A",
  "description": "DESC A"
}
```

## 📚 Learning Objectives

This project helps you learn:

- ✅ Go project structure and module management
- ✅ HTTP server creation using `net/http`
- ✅ HTTP request routing and handling
- ✅ JSON encoding/decoding
- ✅ HTTP methods (GET, POST, PUT, DELETE)
- ✅ Error handling in web applications
- ✅ RESTful API design principles
- ✅ URL path parsing and parameter extraction

## 🗂️ Project Structure

```
learn-go/
├── main.go           # Main application file with handlers
├── go.mod            # Go module definition
├── docs/
│   └── assignment.md # Learning assignment details
└── README.md         # This file
```

## 🎯 Assignment (Session 1)

Implement full CRUD operations for the Category model:

- [x] GET `/categories` - Retrieve all categories
- [x] GET `/categories/{id}` - Retrieve a single category
- [x] POST `/categories` - Create a new category
- [x] PUT `/categories/{id}` - Update an existing category
- [x] DELETE `/categories/{id}` - Delete a category
- [ ] sqlite
- [ ] unit test

## 📖 Resources

- [Go Documentation](https://golang.org/doc/)
- [net/http Package](https://pkg.go.dev/net/http)
- [JSON Package](https://pkg.go.dev/encoding/json)
- [RESTful API Design Best Practices](https://restfulapi.net/)

## 📄 License

This is a learning project and is free to use for educational purposes.

---

**Happy Learning! 🎉**
