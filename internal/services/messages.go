package services

import (
	"message-sender-bot/internal/repository"

	"github.com/rs/zerolog"
)

type MessageService struct {
	msgRepo  *repository.MessageRepository
	tgBotSvc TelegramBotServiceInterface
	logger   *zerolog.Logger
}

func NewMessageService(msgRepo *repository.MessageRepository, tgBotSvc TelegramBotServiceInterface, logger *zerolog.Logger) *MessageService {
	return &MessageService{
		msgRepo:  msgRepo,
		tgBotSvc: tgBotSvc,
		logger:   logger,
	}
}
