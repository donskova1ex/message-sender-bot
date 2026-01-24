package handlers

import (
	"message-sender-bot/internal/dto"
	"message-sender-bot/internal/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

const (
	PAGE_ITEMS = 10
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
	api.Get("/health", handler.healthCheck)

	protected := api.Group("", jwtMiddleware)
	protected.Get("/test", handler.Test)
	protected.Post("/scheduled-messages", handler.ScheduleMessage)
	protected.Get("/scheduled-messages", handler.UnsentMessages)
	return handler

}
func (handler *MessageHandler) healthCheck(c *fiber.Ctx) error {
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

func (handler *MessageHandler) ScheduleMessage(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(int64)
	messageDto := &dto.ScheduleMessageRequest{}
	if err := c.BodyParser(messageDto); err != nil {
		handler.logger.Error().Err(err).Msg("failed to parse schedule message request body")
		return HandleError(err).Send(c)
	}
	if err := handler.msgSvc.ScheduleMessage(c.Context(), messageDto, userId); err != nil {
		handler.logger.Error().Err(err).Msg("failed to schedule message")
		return HandleError(err).Send(c)
	}

	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"status":  "ok",
			"message": "message scheduled",
		},
	)
}

func (handler *MessageHandler) UnsentMessages(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	unsentMessages, err := handler.msgSvc.GetUnsentMessages(c.Context(), PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		handler.logger.Error().Err(err).Msg("failed to get unsent messages")
		return HandleError(err).Send(c)
	}
	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"status":  "ok",
			"message": unsentMessages,
		},
	)
}
