package handlers

import (
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

func NewAuthHandler(
	router fiber.Router,
	authSvc *services.AuthService,
	logger *zerolog.Logger,
) *AuthHandler {
	handler := &AuthHandler{authSvc: authSvc, logger: logger}
	authGroup := router.Group("/api/v1/auth")
	authGroup.Post("/register", handler.Register)
	authGroup.Post("/login", handler.Login)
	return handler
}

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req registerRequest
	if err := c.BodyParser(&req); err != nil {
		//return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
		return HandleError(err).Send(c)
	}

	token, err := h.authSvc.Register(c.Context(), req.Email, req.Password)
	if err != nil {
		h.logger.Warn().Err(err).Str("email", req.Email).Msg("Registration failed")
		if strings.Contains(err.Error(), "email already registered") {
			return HandleError(err).Send(c)
		}
		return HandleError(err).Send(c)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		//return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
		return HandleError(err).Send(c)
	}

	token, err := h.authSvc.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		h.logger.Warn().Str("email", req.Email).Msg("Login failed")
		return HandleError(err).Send(c)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"token":   token,
	})
}
