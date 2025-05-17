# Eagle Tutorials Platform Backend

This is the Go (Gin) + MongoDB backend for the Eagle Tutorials Platform.

## Structure

- `cmd/server/` - Entrypoint for the API server
- `internal/api/` - HTTP handlers
- `internal/db/` - MongoDB connection and models
- `internal/service/` - Business logic
- `internal/middleware/` - Middleware (auth, logging, etc.)
- `pkg/` - Shared utilities

## Quick Start

1. Copy `.env.example` to `.env` and fill in your MongoDB URI and other secrets.
2. Run:
   ```sh
   go mod tidy
   go run ./cmd/server
   ```

## Dependencies
- Gin
- MongoDB Go Driver
- godotenv
- logrus
