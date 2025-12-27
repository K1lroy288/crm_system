package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id           uint           `gorm:"primarykey;autoincrement"`
	Username     string         `gorm:"size:100;not null;unique"`
	PasswordHash string         `gorm:"size:255;not null"`
	FirstName    string         `gorm:"size:50"`
	LastName     string         `gorm:"size:50"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
