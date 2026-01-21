package middleware

import (
	"message-sender-bot/internal/services"
	"strings"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"
)

func isHealthCheck(c *fiber.Ctx) bool {
	return c.Path() == "/api/v1/health"
}

func LoggerMiddleware(app *fiber.App, logger *zerolog.Logger) {
	app.Use(
		fiberzerolog.New(
			fiberzerolog.Config{
				Logger: logger,
				Next: isHealthCheck,
			},
		),
	)
}

func RecoverMiddleware(app *fiber.App) {
	app.Use(
		recover.New(),
	)
}

func CORSMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "POST",
	}))

}

func SetupRateLimiter(app *fiber.App) {
	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 60 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"message": "too many requests",
			})
		},
		Next: isHealthCheck,
	}))
}

func ApplyMiddleware(app *fiber.App, logger *zerolog.Logger) {
	LoggerMiddleware(app, logger)
	RecoverMiddleware(app)
	CORSMiddleware(app)
	SetupRateLimiter(app)
}

func JWTAuth(jwtSvc *services.JWTService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing Authorization header"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid Authorization format"})
		}

		claims, err := jwtSvc.ValidateToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid or expired token"})
		}

		c.Locals("user_id", claims.UserID)
		return c.Next()
	}
}