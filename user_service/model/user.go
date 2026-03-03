package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint   `gorm:"primary key; not null"`
	Username     string `gorm:"size:100;not null;unique"`
	FirstName    string `gorm:"size:50"`
	LastName     string `gorm:"size:50"`
	PasswordHash []byte `gorm:"not null"`

	Roles []Role `gorm:"many2many:user_roles;"`
}
