package router

import (
	_ "workshop-cursor/backend/docs"
	"workshop-cursor/backend/internal/adapter/http/handler"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func Register(app *fiber.App, helloHandler *handler.HelloHandler, authHandler *handler.AuthHandler, verify func(token string) (int64, error)) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello world")
	})
	api := app.Group("/api")
	api.Post("/login", authHandler.Login)

	authGroup := api.Group("", handler.AuthMiddleware(verify))
	authGroup.Get("/me", authHandler.Me)
	authGroup.Put("/me", authHandler.UpdateMe)

	// swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
}
