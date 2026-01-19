package services

import (
	"context"
	"errors"
	"message-sender-bot/internal/repository"
	"regexp"
	"strings"
	"unicode"

	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository *repository.UserRepository
	jwtService *JWTService
	logger *zerolog.Logger
}

func NewAuthService(userRepository *repository.UserRepository, jwtService *JWTService, logger *zerolog.Logger) *AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwtService: jwtService,
		logger: logger,
	}
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	var hasUpper, hasLower, hasDigit bool
	for _, symbol := range password {
		if unicode.IsUpper(symbol) {
			hasUpper = true
		}
		if unicode.IsLower(symbol) {
			hasLower = true
		}
		if unicode.IsDigit(symbol) {
			hasDigit = true
		}
		if hasUpper && hasLower && hasDigit {
			return true
		}
	}
	return false
}

func (a *AuthService) Register(ctx context.Context, email, password string) (string, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)
	if !isValidEmail(email) {
		return "", errors.New("invalid email")
	}
	if !isValidPassword(password) {
		return "", errors.New("Invalid password. Password must be at least 8 characters and contain uppercase, lowercase, and digit")
	}

	user, err := a.userRepository.CreateUser(ctx, email, password)
	if err != nil {
		return "", err
	}
	token, err := a.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := a.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}
	return a.jwtService.GenerateToken(user.ID)
}