package services

import (
	"time"

	custom_errors "message-sender-bot/internal/errors"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct{
	secret []byte
	expiry time.Duration
}

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJWTService(secret []byte, expiry time.Duration) *JWTService {
	return &JWTService{
		secret: secret,
		expiry: expiry,
	}
}

func (j *JWTService) GenerateToken(userID int64) (string, error) {
	expTime := time.Now().Add(j.expiry)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *JWTService) ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, custom_errors.InvalidSingingMethod
		}
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, custom_errors.InvalidTokenError
}