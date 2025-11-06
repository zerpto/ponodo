# Ponodo

A lightweight Go framework for building CLI applications with integrated HTTP server capabilities. 
Ponodo provides a clean architecture foundation with built-in support for database connections, 
configuration management, logging, validation, and structured HTTP responses.

## Features

- **CLI Framework**: Built on Cobra for powerful command-line interface development
- **HTTP Server**: Integrated Gin web server with graceful shutdown handling
- **Database**: GORM integration with PostgreSQL support
- **Configuration**: Environment-based configuration using Viper
- **Logging**: Zerolog integration for structured logging
- **Validation**: Request validation with user-friendly error messages
- **Response Helpers**: Standardized HTTP response helpers for consistent API responses

## Installation

```bash
go get github.com/zerpto/ponodo
```

## Usage

### Basic Setup

```go
package main

import (
    "github.com/zerpto/ponodo"
    "github.com/zerpto/ponodo/cli/handlers"
    "github.com/zerpto/ponodo/config"
)

func main() {
    // Create a new application instance
    app := ponodo.NewApp()
    
    // Setup configuration loader
    configLoader, err := config.NewLoader()
    if err != nil {
        panic(err)
    }
    app.SetConfigLoader(configLoader)
    
    // Setup base dependencies (logger, database, etc.)
    app.SetupBaseDependencies()
    
    // Add HTTP command handler
    app.AddCommand(func(app contracts.AppContract) clicontracts.CommandContract {
        return handlers.NewHttpHandler(app, setupRoutes)
    })
    
    // Run the application
    app.Run()
}

func setupRoutes(app contracts.AppContract) {
    router := app.GetGin()
    
    // Add your routes here
    router.GET("/health", func(c *gin.Context) {
        response.Ok(c, map[string]string{"status": "ok"})
    })
}
```

### Configuration

Create a `.env` file in your project root:

```env
APP_NAME=myapp
ENV=development
DEBUG=true
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_DATABASE=myapp_db
```

### Implementing Custom Commands

```go
type MyCommand struct {
    App contracts.AppContract
}

func (c *MyCommand) Use() string {
    return "mycommand"
}

func (c *MyCommand) Short() string {
    return "Description of my command"
}

func (c *MyCommand) Long() string {
    return "Detailed description of my command"
}

func (c *MyCommand) Example() string {
    return "myapp mycommand --option value"
}

func (c *MyCommand) Run(cmd *cobra.Command, args []string) {
    // Your command logic here
}

// Add to app
app.AddCommand(func(app contracts.AppContract) clicontracts.CommandContract {
    return &MyCommand{App: app}
})
```

### Using Response Helpers

```go
import "github.com/zerpto/ponodo/response"

// Success responses
response.Ok(ctx, data)              // 200 OK
response.Created(ctx, data)         // 201 Created
response.NoContent(ctx)             // 204 No Content

// Error responses
response.BadRequest(ctx, err)       // 400 Bad Request
response.Unauthorized(ctx, err)     // 401 Unauthorized
response.Forbidden(ctx, err)         // 403 Forbidden
response.NotFound(ctx, err)          // 404 Not Found
response.InternalServerError(ctx, err) // 500 Internal Server Error
```

### Request Validation

```go
import (
    "github.com/zerpto/ponodo/request"
    "github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
    request.BaseRequest
    Email string `json:"email" validate:"required,email"`
    Name  string `json:"name" validate:"required,min=3"`
}

func createUser(ctx *gin.Context) {
    var req CreateUserRequest
    req.Ctx = ctx
    req.Validator = app.GetValidator()
    
    if err := ctx.ShouldBindJSON(&req); err != nil {
        response.BadRequest(ctx, err)
        return
    }
    
    if err := req.Validator.Struct(req); err != nil {
        response.BadRequest(ctx, err)
        return
    }
    
    // Process request...
    response.Created(ctx, user)
}
```

### Database Access

```go
db := app.GetDb()

// Use GORM normally
type User struct {
    ID   uint
    Name string
}

db.Create(&User{Name: "John"})
```

## Development

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -cover

# Generate coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Regenerating Mocks

The project uses [gomock](https://github.com/golang/mock) for generating test mocks. Mock generation directives are defined at the top of each interface file using `//go:generate` comments.

To regenerate all mocks:

```bash
go generate ./...
```

This will regenerate mocks for:
- `contracts/AppContract` → `mocks/mock_app_contract.go`
- `cli/contracts/CommandContract` → `mocks/mock_command_contract.go`
- `config/contracts/ConfigContract` and `DbConfigContract` → `mocks/mock_config_contract.go`

**Prerequisites for mock generation:**
```bash
go install go.uber.org/mock/mockgen@latest
```

## Requirements

- Go 1.24.0 or higher
- PostgreSQL (for database features)
- mockgen (for generating test mocks during development)

## License

[Add your license here]

