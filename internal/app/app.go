// internal/app/run.go
package app

import (
	"context"
	"time"

	"message-sender-bot/internal/handlers"
	"message-sender-bot/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func Run(ctx context.Context, tgBotSvc *services.TelegramBotService, authSvc *services.AuthService, jwtSvc *services.JWTService, logger *zerolog.Logger) {
	app := newApp(fiber.New(), tgBotSvc, authSvc, jwtSvc, logger)
	serverStarted := make(chan bool, 1)

	app.Hooks().OnListen(
		func(data fiber.ListenData) error {
			logger.Info().Str("host", data.Host).Str("port", data.Port).Msg("Server started...")
			serverStarted <- true
			return nil
		},
	)
	go func() {
		logger.Info().Msg("Starting server...")
		if err := app.Listen(":3000"); err != nil {
			logger.Fatal().Err(err).Msg("Server failed to start")
		}
	}()

	select {
	case <-serverStarted:
	case <-time.After(time.Second * 2):
		logger.Fatal().Msg("Server start timed out - failed to start server")
	}

	<-ctx.Done()

	logger.Info().Msg("Context done.Shutting down...")
	if err := app.Shutdown(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to shutdown server")
	}
}

func newApp(
	app *fiber.App,
	tgBotSvc *services.TelegramBotService,
	authSvc *services.AuthService,
	jwtSvc *services.JWTService,
	logger *zerolog.Logger,
) *fiber.App {
	ApplyMiddleware(app, logger)

	handlers.NewAuthHandler(app, authSvc, logger)          
	handlers.NewNotificationHandler(app, tgBotSvc, logger) 


	logger.Info().Msg("Handlers registered")
	return app
}
