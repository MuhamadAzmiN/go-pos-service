package utils

import (
	"my-echo-chat_service/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID string, config config.Jwt) (string, error) {
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Duration(config.Exp) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Key))
}


