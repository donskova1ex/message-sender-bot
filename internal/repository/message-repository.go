package repository

import (
	"context"
	"fmt"
	"message-sender-bot/internal/models"
	"time"

	custom_errors "message-sender-bot/internal/errors"

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
		return custom_errors.FailedToCreateMessage
	}
	return nil
}

func (r *MessageRepository) GetUnsentMessages(ctx context.Context, limit, offset int) ([]models.UnsentMessage, error) {
	query := `
			SELECT
				m.id,
				m.planned_date,
				mt.type_name,
				m.text,
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

func (r *MessageRepository) DeleteMessage(ctx context.Context, id int64) error {
	query := `UPDATE messages SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	result, err := r.dbPool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete message: %w", err)
	}
	if result.RowsAffected() == 0 {
		return custom_errors.MessageNotFoundError
	}
	return nil
}
func (r *MessageRepository) UpdateMessage(ctx context.Context, plannedDate *time.Time, typeId *int64, text *string, id int64) error {
	query := `UPDATE messages 
				SET 
				  planned_date = COALESCE($1, planned_date),
				  type_id = COALESCE($2, type_id),
				  text = COALESCE($3, text),
				  updated_at = NOW()
				WHERE 
				  id = $4 
				  AND deleted_at IS NULL`
	result, err := r.dbPool.Exec(ctx, query, plannedDate, typeId, text, id)
	if err != nil {
		return fmt.Errorf("failed to update message: %w", err)
	}
	if result.RowsAffected() == 0 {
		return custom_errors.MessageNotFoundError
	}
	return nil
}

func (r *MessageRepository) GetDeletedMessages(ctx context.Context, limit, offset int) ([]models.DeletedMessage, error) {
	query := `SELECT
					m.id,
					m.planned_date,
					mt.type_name,
					m.text,
					m.deleted_at
				FROM messages m
				JOIN message_types mt ON m.type_id = mt.id
				WHERE m.deleted_at IS NOT NULL
				ORDER BY m.deleted_at DESC
				LIMIT $1 OFFSET $2`

	rows, err := r.dbPool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.DeletedMessage])
	if err != nil {
		return nil, fmt.Errorf("failed to collect rows: %w", err)
	}
	return result, nil
}
