package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleName string `gorm:"uniqueIndex;size:50;not null"`
	Users    []User `gorm:"many2many:user_roles;"`
}
