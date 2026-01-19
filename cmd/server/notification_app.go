// cmd/app/main.go
package main

import (
	"context"
	"message-sender-bot/config"
	"message-sender-bot/internal/app"
	"message-sender-bot/internal/repository"
	"message-sender-bot/internal/services"
	"message-sender-bot/pkg/db"
	"message-sender-bot/pkg/logger"
	"os/signal"
	"syscall"
)

func main() {
	config.Init()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logCfg := config.NewLogConfig()
	customLogger := logger.NewLogger(logCfg)

	botCfg, err := config.NewBotConfig()
	if err != nil {
		customLogger.Fatal().Err(err).Msg("Failed to load bot config")
	}
	tgBotSvc, err := services.NewTelegramBotService(botCfg, customLogger)
	if err != nil {
		customLogger.Fatal().Err(err).Msg("Failed to create telegram bot service")
	}

	dbCfg := config.NewDBConfig()
	if dbCfg.PG_DSN == "" {
		customLogger.Fatal().Msg("Postgres connection string is empty. Please set PG_DSN environment variable")
	}
	dbPool, err := db.CreateDBPool(ctx, dbCfg)
	if err != nil {
		customLogger.Fatal().Err(err).Msg("Failed to create database pool")
	}
	defer dbPool.Close()

	authCfg, err := config.NewAuthConfig()
	if err != nil {
		customLogger.Fatal().Err(err).Msg("Failed to load auth config")
	}
	jwtSvc := services.NewJWTService([]byte(authCfg.Secret), authCfg.Exp)
	userRepo := repository.NewUserRepository(dbPool)
	authSvc := services.NewAuthService(userRepo, jwtSvc, customLogger)

	go app.Run(ctx, tgBotSvc, authSvc, jwtSvc, customLogger)
	<-ctx.Done()
}