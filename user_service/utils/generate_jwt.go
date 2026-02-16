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
	var rolesNames []string
	for _, role := range user.Roles {
		rolesNames = append(rolesNames, role.RoleName)
	}

	claims := model.CustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		Roles:    rolesNames,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "user-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("JWT signature error: %w", err)
	}

	return tokenString, nil
}
