package repository

import (
	"context"
	"fmt"
	"message-sender-bot/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MessageRepository struct {
	dbPool *pgxpool.Pool
}

func NewMessageRepository(dbPool *pgxpool.Pool) *MessageRepository {
	return &MessageRepository{
		dbPool: dbPool,
	}
}

func (r *MessageRepository) CreateMessage(ctx context.Context, message *models.Message) error {
	query := `INSERT into messages (planned_date, user_id, type_id, text, created_at) VALUES ($1, $2, $3, $4, $5)`
	result, err := r.dbPool.Exec(ctx, query, message.PlannedDate, message.UserID, message.TypeID, message.Text, time.Now())
	if err != nil {
		return fmt.Errorf("failed to create message: %w", err)
	}

	if result.RowsAffected() ==  0 {
		return fmt.Errorf("failed to create message: no rows affected")
	}
	return nil
}
