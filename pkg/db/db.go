package db

import (
	"context"
	"fmt"
	"message-sender-bot/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDBPool(ctx context.Context, cfg *config.PGConfig) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(ctx, cfg.PG_DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to pgSQL: %w", err)
	}
	if err := dbPool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping pgSQL: %w", err)
	}
	return dbPool, nil
}