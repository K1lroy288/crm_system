package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id           uint           `gorm:"primarykey;autoincrement"`
	Username     string         `gorm:"size:100;not null;unique"`
	FirstName    string         `gorm:"size:50"`
	LastName     string         `gorm:"size:50"`
	PasswordHash []byte         `gorm:"not null"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type UserDTO struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
