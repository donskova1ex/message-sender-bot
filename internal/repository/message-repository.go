package repository

import "github.com/jackc/pgx/v5/pgxpool"

type MessageRepository struct {
	dbPool *pgxpool.Pool
}

func NewMessageRepository(dbPool *pgxpool.Pool) *MessageRepository {
	return &MessageRepository{
		dbPool: dbPool,
	}
}