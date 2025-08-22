package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"

    "workshop-cursor/backend/internal/adapter/http/router"
    "workshop-cursor/backend/internal/di"
)

func main() {
    app := fiber.New()

    // middlewares
    app.Use(recover.New())
    app.Use(logger.New())

    // DI container
    container := di.NewContainer()

    // routes
    router.Register(app, container.HelloHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    if err := app.Listen(":" + port); err != nil {
        log.Fatalf("failed to start server: %v", err)
    }
}


