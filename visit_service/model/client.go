package model

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	FirstName string `gorm:"size:50;not null"`
	LastName  string `gorm:"size:50;not null"`
	Surname   string `gorm:"size:50"`
	Phone     string
}
