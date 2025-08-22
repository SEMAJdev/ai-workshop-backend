## Backend (Go + Fiber)

### Structure
```
backend/
  cmd/api/main.go                # app entry
  internal/
    core/hello/                  # domain contracts
    usecase/hello/               # application logic
    adapter/
      repository/memory/         # in-memory repo
      http/handler/              # HTTP handlers
      http/router/               # route registration
    di/                          # dependency wiring
```

### Run

1. Install Go 1.21+.
2. Install deps and run:
```bash
cd backend
go build ./...
PORT=8080 go run cmd/api/main.go
```

### Endpoint

- GET `/api/hello`
  - Response: `{ "message": "Hello world" }`


