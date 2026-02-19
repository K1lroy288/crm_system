package utils

import (
	"errors"
	"fmt"
	"strings"
	"user-service/config"
	"user-service/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(ctx *gin.Context) (*model.CustomClaims, error) {
	claims := &model.CustomClaims{}
	tokenReq := ctx.Request.Header.Get("Authorization")

	tokenString := ""

	if tokenReq != "" {
		tokenString = strings.TrimPrefix(tokenReq, "Bearer ")
	}

	if tokenString == "" {
		tokenString = ctx.Query("token")
	}

	if tokenString == "" {
		return nil, fmt.Errorf("missing authorization token")
	}

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
