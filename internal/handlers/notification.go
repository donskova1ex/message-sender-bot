package handlers

import (
	"message-sender-bot/internal/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type Notification struct {
	router             fiber.Router
	telegramBotService services.TelegramBotServiceInterface
	logger             *zerolog.Logger
	StartTime time.Time
}

func NewNotificationHandler(router fiber.Router, telegramBotService services.TelegramBotServiceInterface, jwtMiddlware fiber.Handler, logger *zerolog.Logger) *Notification {
	handler := &Notification{
		router:             router,
		telegramBotService: telegramBotService,
		logger:             logger,
		StartTime: time.Now(),
	}
	api := handler.router.Group("/api/v1")
	api.Get("/health", handler.HealthCheck)
	
	protected := api.Group("", jwtMiddlware)
	protected.Get("/test", handler.Test)
	return handler

}
func (n *Notification) HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "ok",
			"uptime": time.Since(n.StartTime).String(),
		})
}

func (n *Notification) Test(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(int64)
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"status": "ok",
			"user_id": userId,
		})
}