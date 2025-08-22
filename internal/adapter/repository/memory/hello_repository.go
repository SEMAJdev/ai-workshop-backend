package memory

import "workshop-cursor/backend/internal/core/hello"

// InMemoryGreeterRepository is a simple in-memory implementation.
type InMemoryGreeterRepository struct{}

func NewInMemoryGreeterRepository() *InMemoryGreeterRepository {
    return &InMemoryGreeterRepository{}
}

func (r *InMemoryGreeterRepository) GetGreeting() (string, error) {
    return "Hello world", nil
}

var _ hello.GreeterRepository = (*InMemoryGreeterRepository)(nil)


