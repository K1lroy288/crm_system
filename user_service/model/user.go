package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"size:100;not null;unique"`
	FirstName    string `gorm:"size:50"`
	LastName     string `gorm:"size:50"`
	PasswordHash []byte `gorm:"not null"`
}

type UserDTO struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
