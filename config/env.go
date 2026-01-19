package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type BotConfig struct {
	BotToken string
	ChatID   int64
}

type LogConfig struct {
	Level  int
	Format string
}

type PGConfig struct {
	PG_DSN string
}

type AuthConfig struct {
	Secret string
	Exp    time.Duration
}

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("successfully loaded .env file")
}

func getInt(key string, defaultString int) int {
	valStr := os.Getenv(key)

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultString
	}

	return val
}

func getString(key, defaultString string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultString
	}
	return val
}

func NewBotConfig() (*BotConfig, error) {
	botToken := getString("BOT_TOKEN", "")
	if botToken == "" {
		return nil, errors.New("bot token required")
	}
	chatID := getInt("CHAT_ID", 0)
	if chatID == 0 {
		return nil, errors.New("bot chat id required")
	}

	return &BotConfig{
		BotToken: botToken,
		ChatID:   int64(chatID),
	}, nil
}

func NewAuthConfig() (*AuthConfig, error) {
	secret := getString("JWT_SECRET", "")
	if secret == "" {
		return nil, errors.New("secret required")
	}
	expiry := getString("JWT_EXPRY", "24h")
	actualExpiry, err := time.ParseDuration(expiry)
	if err != nil {
		return nil, fmt.Errorf("invalid expiry: %w", err)
	}
	
	return &AuthConfig{
		Secret: secret,
		Exp: actualExpiry,
	}, nil
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level:  getInt("LOG_LEVEL", 0),
		Format: getString("LOG_FORMAT", "json"),
	}
}

func NewDBConfig() *PGConfig {
	return &PGConfig{
		PG_DSN: os.Getenv("PG_DSN"),
	}
}
