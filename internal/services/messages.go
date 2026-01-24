package services

import (
	"context"
	"fmt"
	"message-sender-bot/internal/dto"
	custom_errors "message-sender-bot/internal/errors"
	"message-sender-bot/internal/models"
	"message-sender-bot/internal/repository"
	"time"

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

func (s *MessageService) ScheduleMessage(ctx context.Context, msgDto *dto.ScheduleMessageRequest, userId int64) error {
	err := s.validate(msgDto)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	msgModel := s.messageRequestDtoToMessageModel(msgDto, userId)
	if err = s.msgRepo.CreateMessage(ctx, msgModel); err != nil {
		return err
	}
	return nil
}

func (s *MessageService) GetUnsentMessages(ctx context.Context, limit, offset int) ([]*dto.MessagesResponse, error) {
	modelsMessages, err := s.msgRepo.GetUnsentMessages(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	modelSlice := make([]*dto.MessagesResponse, len(modelsMessages))
	for i, model := range modelsMessages {
		modelSlice[i] = &dto.MessagesResponse{
			PlannedDate: model.PlannedDate,
			Type:        model.TypeName,
			Text:        model.Text,
			IsSent:      false,
		}
	}
	return modelSlice, nil
}

func (s *MessageService) messageRequestDtoToMessageModel(msgDto *dto.ScheduleMessageRequest, userId int64) *models.Message {
	msgModel := &models.Message{
		UserID:      userId,
		PlannedDate: msgDto.PlannedDate,
		TypeID:      msgDto.TypeID,
		Text:        msgDto.Text,
	}
	return msgModel

}

func (s *MessageService) validate(messageDto *dto.ScheduleMessageRequest) error {
	if messageDto.PlannedDate.IsZero() || messageDto.PlannedDate.Before(time.Now()) {
		return custom_errors.EmptyPlannedDateError
	}
	if messageDto.Text == "" {
		return custom_errors.EmptyTextMessageError
	}

	if messageDto.TypeID <= 0 {
		return custom_errors.EmptyOrNegativeTypeId
	}

	return nil
}
