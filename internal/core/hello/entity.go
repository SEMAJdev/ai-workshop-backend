package hello

// GreeterRepository defines the contract to fetch greeting messages.
type GreeterRepository interface {
    GetGreeting() (string, error)
}


