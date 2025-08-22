package handler

import (
    "strings"

    "github.com/gofiber/fiber/v2"
    authuc "workshop-cursor/backend/internal/usecase/auth"
    core "workshop-cursor/backend/internal/core/user"
)

type AuthHandler struct {
    uc *authuc.UseCase
}

func NewAuthHandler(uc *authuc.UseCase) *AuthHandler {
    return &AuthHandler{uc: uc}
}

type loginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
    var req loginRequest
    if err := ctx.BodyParser(&req); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
    }
    token, user, err := h.uc.Login(req.Email, req.Password)
    if err != nil {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
    }
    return ctx.JSON(fiber.Map{"token": token, "user": toPublic(user)})
}

func (h *AuthHandler) Me(ctx *fiber.Ctx) error {
    userID, ok := ctx.Locals("userID").(int64)
    if !ok {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
    }
    user, err := h.uc.GetProfile(userID)
    if err != nil {
        return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
    }
    return ctx.JSON(toPublic(user))
}

func (h *AuthHandler) UpdateMe(ctx *fiber.Ctx) error {
    userID, ok := ctx.Locals("userID").(int64)
    if !ok {
        return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
    }
    var input core.UpdateProfileInput
    if err := ctx.BodyParser(&input); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
    }
    user, err := h.uc.UpdateProfile(userID, input)
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    return ctx.JSON(toPublic(user))
}

func toPublic(u *core.User) core.PublicProfile {
    return core.PublicProfile{
        ID: u.ID,
        Email: u.Email,
        FirstName: u.FirstName,
        LastName: u.LastName,
        Phone: u.Phone,
        MemberCode: u.MemberCode,
        MembershipLevel: u.MembershipLevel,
        Points: u.Points,
        JoinedAt: u.JoinedAt,
    }
}

// AuthMiddleware validates Bearer token and sets userID in context.
func AuthMiddleware(verify func(token string) (int64, error)) fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        header := ctx.Get("Authorization")
        if header == "" || !strings.HasPrefix(header, "Bearer ") {
            return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing token"})
        }
        token := strings.TrimPrefix(header, "Bearer ")
        userID, err := verify(token)
        if err != nil {
            return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
        }
        ctx.Locals("userID", userID)
        return ctx.Next()
    }
}


