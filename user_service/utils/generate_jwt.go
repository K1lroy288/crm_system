package utils

import (
	"fmt"
	"time"
	"user-service/config"
	"user-service/model"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user model.User) (string, error) {
	cfg := config.GetConfig()
	jwtKey := []byte(cfg.JwtSecret)
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("ошибка в подписи jwt токена: %w", err)
	}

	return tokenString, nil
}
