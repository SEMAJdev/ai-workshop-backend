package router

import (
    "github.com/gofiber/fiber/v2"
    "workshop-cursor/backend/internal/adapter/http/handler"
)

func Register(app *fiber.App, helloHandler *handler.HelloHandler) {
    app.Get("/", func(ctx *fiber.Ctx) error {
        return ctx.SendString("Hello world")
    })
}
