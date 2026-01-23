package repository

import (
	"context"
	"fmt"
	"message-sender-bot/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
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

	if result.RowsAffected() == 0 {
		return fmt.Errorf("failed to create message: no rows affected")
	}
	return nil
}

func (r *MessageRepository) GetUnsentMessages(ctx context.Context, limit, offset int) ([]models.UnsentMessage, error) {
	query := `
			SELECT
				m.id,
				m.planned_date,
				mt.type_name,
				m.text
			FROM messages m
					 JOIN message_types mt ON m.type_id = mt.id
			WHERE m.deleted_at IS NULL
			  AND NOT EXISTS (
				SELECT 1
				FROM message_deliveries md
				WHERE md.message_id = m.id
			)
			ORDER BY m.planned_date ASC
			LIMIT $1 OFFSET $2`
	rows, err := r.dbPool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.UnsentMessage])
	if err != nil {
		return nil, fmt.Errorf("failed to collect rows: %w", err)
	}
	return result, nil
}
