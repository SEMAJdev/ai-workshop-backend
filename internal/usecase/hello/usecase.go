package hello

import (
    core "workshop-cursor/backend/internal/core/hello"
)

// UseCase provides application logic to get a greeting.
type UseCase struct {
    greeterRepository core.GreeterRepository
}

func NewUseCase(greeterRepository core.GreeterRepository) *UseCase {
    return &UseCase{greeterRepository: greeterRepository}
}

func (u *UseCase) GetGreeting() (string, error) {
    return u.greeterRepository.GetGreeting()
}


