package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"workshop-cursor/backend/internal/adapter/http/router"
	"workshop-cursor/backend/internal/config"
	"workshop-cursor/backend/internal/di"
)

func main() {
	cfg := config.Load()
	app := fiber.New()

	// middlewares
	app.Use(recover.New())
	app.Use(logger.New())

	// DI container
	container := di.NewContainer(cfg)

	// routes
	router.Register(app, container.HelloHandler, container.AuthHandler, container.AuthMW)

	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
