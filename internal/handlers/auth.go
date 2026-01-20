package handlers

import (
	"net/http"
	"strings"

	"message-sender-bot/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type AuthHandler struct {
	authSvc *services.AuthService
	logger  *zerolog.Logger
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthHandler(app *fiber.App, authSvc *services.AuthService, logger *zerolog.Logger) {
	h := &AuthHandler{authSvc: authSvc, logger: logger}
	app.Post("/api/v1/auth/register", h.Register)
	app.Post("/api/v1/auth/login", h.Login)
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req registerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
	}

	token, err := h.authSvc.Register(c.Context(), req.Email, req.Password)
	if err != nil {
		h.logger.Warn().Err(err).Str("email", req.Email).Msg("Registration failed")
		if strings.Contains(err.Error(), "email already registered") {
			return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "email already registered"})
		}
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
	}

	token, err := h.authSvc.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		h.logger.Warn().Str("email", req.Email).Msg("Login failed")
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}