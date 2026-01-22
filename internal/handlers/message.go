package handlers

import (
	"message-sender-bot/internal/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type MessageHandler struct {
	router    fiber.Router
	msgSvc    *services.MessageService
	logger    *zerolog.Logger
	StartTime time.Time
}

func NewMessageHandler(router fiber.Router, msgSvc *services.MessageService, jwtMiddleware fiber.Handler, logger *zerolog.Logger) *MessageHandler {
	handler := &MessageHandler{
		router:    router,
		msgSvc:    msgSvc,
		logger:    logger,
		StartTime: time.Now(),
	}
	api := handler.router.Group("/api/v1")
	api.Get("/health", handler.HealthCheck)

	protected := api.Group("", jwtMiddleware)
	protected.Get("/test", handler.Test)
	return handler

}
func (handler *MessageHandler) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "ok",
			"uptime": time.Since(handler.StartTime).String(),
		})
}

func (handler *MessageHandler) Test(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(int64)
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status":  "ok",
			"user_id": userId,
		})
}
