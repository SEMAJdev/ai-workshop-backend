package di

import (
    memrepo "workshop-cursor/backend/internal/adapter/repository/memory"
    "workshop-cursor/backend/internal/adapter/http/handler"
    usecase "workshop-cursor/backend/internal/usecase/hello"
)

type Container struct {
    HelloHandler *handler.HelloHandler
}

func NewContainer() *Container {
    // repositories
    greeterRepo := memrepo.NewInMemoryGreeterRepository()

    // use cases
    helloUC := usecase.NewUseCase(greeterRepo)

    // handlers
    helloHandler := handler.NewHelloHandler(helloUC)

    return &Container{
        HelloHandler: helloHandler,
    }
}


