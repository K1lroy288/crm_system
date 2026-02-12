package utils

import (
	"api-gateway/config"
	"api-gateway/model"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenString string) (*model.CustomClaims, error) {
	claims := &model.CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
		}

		return []byte(config.GetConfig().JwtSecret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token has expired")
		}

		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
