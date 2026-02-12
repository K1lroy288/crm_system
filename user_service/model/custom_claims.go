package model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID   uint     `json:"user_id"`
	Roles    []string `json:"roles"`
	Username string   `json:"username"`
	jwt.RegisteredClaims
}
