package handler

import (
    "github.com/gofiber/fiber/v2"
    usecase "workshop-cursor/backend/internal/usecase/hello"
)

type HelloHandler struct {
    uc *usecase.UseCase
}

func NewHelloHandler(uc *usecase.UseCase) *HelloHandler {
    return &HelloHandler{uc: uc}
}

func (h *HelloHandler) GetHello(ctx *fiber.Ctx) error {
    greeting, err := h.uc.GetGreeting()
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return ctx.JSON(fiber.Map{"message": greeting})
}


